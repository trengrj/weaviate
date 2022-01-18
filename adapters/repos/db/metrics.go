//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package db

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/semi-technologies/weaviate/usecases/monitoring"
	"github.com/sirupsen/logrus"
)

type Metrics struct {
	logger     logrus.FieldLogger
	monitoring bool
	batchTime  prometheus.ObserverVec
}

func NewMetrics(logger logrus.FieldLogger, prom *monitoring.PrometheusMetrics,
	className, shardName string) *Metrics {
	m := &Metrics{
		logger: logger,
	}

	if prom != nil {
		m.monitoring = true
		batchTime, err := prom.BatchTime.CurryWith(prometheus.Labels{
			"class_name": className,
			"shard_name": shardName,
		})
		if err != nil {
			panic(err)
		}
		m.batchTime = batchTime
	}

	return m
}

func (m *Metrics) BatchObject(start time.Time, size int) {
	took := time.Since(start)
	m.logger.WithField("action", "batch_objects").
		WithField("batch_size", size).
		WithField("took", took).
		Tracef("object batch took %s", took)
}

func (m *Metrics) ObjectStore(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "store_object_store").
		WithField("took", took).
		Tracef("storing objects in KV/inverted store took %s", took)

	if m.monitoring {
		m.batchTime.With(prometheus.Labels{"operation": "object_storage"}).
			Observe(float64(took / time.Millisecond))
	}
}

func (m *Metrics) VectorIndex(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "store_vector_index").
		WithField("took", took).
		Tracef("storing objects vector index took %s", took)

	if m.monitoring {
		m.batchTime.With(prometheus.Labels{"operation": "vector_storage"}).
			Observe(float64(took / time.Millisecond))
	}
}

func (m *Metrics) PutObject(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "store_object_store_single_object_in_tx").
		WithField("took", took).
		Tracef("storing single object (complete) in KV/inverted took %s", took)
}

func (m *Metrics) PutObjectUpsertObject(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "store_object_store_upsert_object_data").
		WithField("took", took).
		Tracef("storing object data in KV took %s", took)
}

func (m *Metrics) PutObjectUpdateDocID(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "store_object_store_update_index_id").
		WithField("took", took).
		Tracef("updating doc id in KV/inverted took %s", took)
}

func (m *Metrics) PutObjectUpdateInverted(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "store_object_store_update_inverted").
		WithField("took", took).
		Tracef("updating inverted index for single object took %s", took)
}

func (m *Metrics) InvertedDeleteOld(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "inverted_delete_old").
		WithField("took", took).
		Tracef("deleting old entries from inverted index %s", took)
}

func (m *Metrics) InvertedDeleteDelta(start time.Time) {
	took := time.Since(start)
	m.logger.WithField("action", "inverted_delete_delta").
		WithField("took", took).
		Tracef("deleting delta entries from inverted index %s", took)
}

func (m *Metrics) InvertedExtend(start time.Time, propCount int) {
	took := time.Since(start)
	m.logger.WithField("action", "inverted_extend").
		WithField("took", took).
		WithField("prop_count", propCount).
		Tracef("extending inverted index took %s", took)
}
