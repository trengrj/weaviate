package clients

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func (c *vectorizer) WaitForStartup(initCtx context.Context,
	interval time.Duration) error {
	t := time.Tick(interval)
	expired := initCtx.Done()
	var lastErr error
	for {
		select {
		case <-t:
			lastErr = c.checkReady(initCtx)
			if lastErr == nil {
				return nil
			}
		case <-expired:
			return errors.Wrapf(lastErr, "init context expired before remote was ready")
		}
	}
}

func (c *vectorizer) checkReady(initCtx context.Context) error {
	// spawn a new context (derived on the overall context) which is used to
	// consider an individual request timed out
	requestCtx, cancel := context.WithTimeout(initCtx, 500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(requestCtx, http.MethodGet,
		c.url("/.well-known/ready"), nil)
	if err != nil {
		return errors.Wrap(err, "create check ready request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "send check ready request")
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return errors.Errorf("not ready: status %d", res.StatusCode)
	}

	return nil
}
