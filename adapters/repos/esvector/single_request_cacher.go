package esvector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v5/esapi"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/sirupsen/logrus"
)

func newCacher(repo *Repo) *cacher {
	return &cacher{
		logger: repo.logger,
		repo:   repo,
		store:  map[storageIdentifier]search.Result{},
	}
}

type cacherJob struct {
	si       storageIdentifier
	props    traverser.SelectProperties
	complete bool
}

type cacher struct {
	sync.Mutex
	jobs   []cacherJob
	logger logrus.FieldLogger
	repo   *Repo
	store  map[storageIdentifier]search.Result
	meta   *bool // meta is immutable for the lifetime of the request cacher, so we can safely store it
}

func (c *cacher) get(si storageIdentifier) (search.Result, bool) {
	sr, ok := c.store[si]
	return sr, ok
}

func (c *cacher) build(ctx context.Context, sr searchResponse,
	properties traverser.SelectProperties, meta bool) error {
	if c.meta == nil {
		// store meta prop if we haven't yet
		c.meta = &meta
	}

	err := c.findJobsFromResponse(sr, properties)
	if err != nil {
		return fmt.Errorf("build request cache: %v", err)
	}

	c.dedupJobList()
	err = c.fetchJobs(ctx)
	if err != nil {
		return fmt.Errorf("build request cache: %v", err)
	}

	return nil
}

func (c *cacher) findJobsFromResponse(sr searchResponse, properties traverser.SelectProperties) error {
	for _, hit := range sr.Hits.Hits {

		var err error
		properties, err = c.replaceInitialPropertiesWithSpecific(hit, properties)
		if err != nil {
			return err
		}

		for key, value := range hit.Source {
			if isInternal(key) {
				continue
			}

			asSlice, ok := value.([]interface{})
			if !ok {
				// not a slice, can't be ref, not interested
				continue
			}

			refKey := uppercaseFirstLetter(key)
			selectProp := properties.FindProperty(refKey)
			if selectProp == nil {
				// user is not interested in this prop
				continue
			}

			for _, selectPropRef := range selectProp.Refs {
				innerProperties := selectPropRef.RefProperties

				for _, item := range asSlice {
					refMap, ok := item.(map[string]interface{})
					if !ok {
						return fmt.Errorf("expected ref item to be a map, but got %T", item)
					}

					beacon, ok := refMap["beacon"]
					if !ok {
						return fmt.Errorf("expected ref object to have field beacon, but got %#v", refMap)
					}

					ref, err := crossref.Parse(beacon.(string))
					if err != nil {
						return err
					}
					c.addJob(storageIdentifier{ref.TargetID.String(), ref.Kind, selectPropRef.ClassName}, innerProperties)
				}
			}
		}
	}

	return nil
}

func (c *cacher) replaceInitialPropertiesWithSpecific(hit hit,
	properties traverser.SelectProperties) (traverser.SelectProperties, error) {

	if properties != nil {
		// don't overwrite the properties if the caller has explicitly set them,
		// this can only mean they're at the root level
		return properties, nil
	}

	// this is a nested level, we cannot rely on global initialSelectProperties
	// anymore, instead we need to find the selectProperties for exactly this
	// ID
	id := hit.ID

	k, err := kind.Parse(hit.Source[keyKind.String()].(string))
	if err != nil {
		return nil, err
	}

	className := hit.Source[keyClassName.String()].(string)
	job, ok := c.findJob(storageIdentifier{id, k, className})
	if ok {
		return job.props, nil
	}

	return properties, nil
}

func (c *cacher) addJob(si storageIdentifier, props traverser.SelectProperties) {

	c.jobs = append(c.jobs, cacherJob{si, props, false})
}

func (c *cacher) findJob(si storageIdentifier) (cacherJob, bool) {
	for _, job := range c.jobs {
		if job.si == si {
			return job, true

		}
	}

	return cacherJob{}, false
}

func (c *cacher) incompleteJobs() []cacherJob {

	out := make([]cacherJob, len(c.jobs))
	n := 0
	for _, job := range c.jobs {
		if job.complete == false {
			out[n] = job
			n++
		}
	}

	return out[:n]
}

func (c *cacher) completeJobs() []cacherJob {

	out := make([]cacherJob, len(c.jobs))
	n := 0
	for _, job := range c.jobs {
		if job.complete == true {
			out[n] = job
			n++
		}
	}

	return out[:n]
}

func (c *cacher) dedupJobList() {
	incompleteJobs := c.incompleteJobs()
	before := len(incompleteJobs)
	if before == 0 {
		// nothing to do
		return
	}
	c.logger.
		WithFields(logrus.Fields{
			"action": "request_cacher_dedup_joblist_start",
			"jobs":   before,
		}).
		Debug("starting job list deduplication")
	deduped := make([]cacherJob, len(incompleteJobs))
	found := map[storageIdentifier]struct{}{}

	n := 0
	for _, job := range incompleteJobs {
		if _, ok := found[job.si]; ok {
			continue
		}

		found[job.si] = struct{}{}
		deduped[n] = job
		n++
	}

	c.jobs = append(c.completeJobs(), deduped[:n]...)

	c.logger.
		WithFields(logrus.Fields{
			"action":      "request_cacher_dedup_joblist_complete",
			"jobs":        n,
			"removedJobs": before - n,
		}).
		Debug("completed job list deduplication")
}

type mgetBody struct {
	Docs []mgetDoc `json:"docs"`
}

type mgetDoc struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
}

func (c *cacher) fetchJobs(ctx context.Context) error {
	before := time.Now()
	jobs := c.incompleteJobs()
	if len(jobs) == 0 {
		c.logger.
			WithFields(
				logrus.Fields{
					"action": "request_cacher_fetch_jobs_skip",
				}).
			Debug("skip fetch jobs, have no incomplete jobs")
		return nil
	}

	docs := make([]mgetDoc, len(jobs))
	for i, job := range jobs {
		docs[i] = mgetDoc{
			Index: classIndexFromClassName(job.si.kind, job.si.className),
			ID:    job.si.id,
		}

	}
	body := mgetBody{docs}

	c.repo.requestCounter.Inc()

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}

	req := esapi.MgetRequest{
		Body: &buf,
	}

	res, err := req.Do(ctx, c.repo.client)
	if err != nil {
		return err
	}

	if err := errorResToErr(res, c.logger); err != nil {
		return err
	}

	took := time.Since(before)

	c.logger.
		WithFields(
			logrus.Fields{
				"action": "request_cacher_fetch_jobs_complete",
				"took":   took,
				"jobs":   len(jobs),
			}).
		Debug("fetch jobs complete")

	return c.parseAndStore(ctx, res)
}

type mgetResponse struct {
	Docs []hit `json:"docs"`
}

func (c *cacher) parseAndStore(ctx context.Context, res *esapi.Response) error {
	if err := errorResToErr(res, c.logger); err != nil {
		return err
	}

	var mgetRes mgetResponse
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&mgetRes)
	if err != nil {
		return fmt.Errorf("decode json: %v", err)
	}

	sr := mgetResToSearchResponse(mgetRes)

	// this is our exit condition for the recursion
	c.markAllJobsAsDone()

	err = c.build(ctx, sr, nil, *c.meta)
	if err != nil {
		return fmt.Errorf("build nested cache: %v", err)
	}

	asResults, err := sr.toResults(c.repo, nil, false, c)
	if err != nil {
		return err
	}

	err = c.storeResults(asResults)
	if err != nil {
		return err
	}

	return nil
}

// remaps from x.docs to x.hits.hits, so we can use existing infrastructure to
// parse it
func mgetResToSearchResponse(in mgetResponse) searchResponse {
	return searchResponse{
		Hits: struct {
			Hits []hit `json:"hits"`
		}{
			Hits: removeEmptyResults(in.Docs),
		},
	}
}

func removeEmptyResults(in []hit) []hit {

	out := make([]hit, len(in))
	n := 0
	for _, hit := range in {
		if hit.Source != nil {
			out[n] = hit
			n++
		}
	}

	return out[0:n]
}

func (c *cacher) storeResults(res search.Results) error {

	for _, item := range res {
		c.store[storageIdentifier{
			id:        item.ID.String(),
			kind:      item.Kind,
			className: item.ClassName,
		}] = item
	}

	return nil
}

func (c *cacher) markAllJobsAsDone() {

	for i := range c.jobs {
		c.jobs[i].complete = true
	}
}
