package http_client

import (
	"fmt"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
	"github.com/sirupsen/logrus"
	"net/http"

	"io/ioutil"
)

type Client struct {
	endpoint string
	client   http.Client
	logger   *logrus.Logger
}

func NewClient(endpoint string) *Client {
	logger := logrus.New()
	logger.Out = ioutil.Discard

	c := Client{
		endpoint: endpoint,
		client:   http.Client{},
		logger:   logger,
	}

	return &c
}

func (c *Client) SetLogger(logger *logrus.Logger) {
	c.logger = logger
}

func (c *Client) Ping() error {
	q := gremlin.RawQuery("1+41")
	response, err := c.Execute(q)
	if err != nil {
		return err
	}

	i, err := response.OneInt()
	if err != nil {
		return err
	}

	if i != 42 {
		return fmt.Errorf("Could not connnected to Gremlin server. Expected the answer to a test query to be 42', but it was %v", i)
	}

	return nil
}
