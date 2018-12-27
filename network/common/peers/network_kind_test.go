package peers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/network/crossrefs"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestGetKindWithoutPeers(t *testing.T) {
	peers := Peers{}
	thing := crossrefs.NetworkKind{
		Kind:     kind.THING_KIND,
		PeerName: "WeaviateB",
		ID:       "doesnt-matter",
	}

	_, err := peers.RemoteKind(thing)
	assert.NotEqual(t, nil, err, "should have an error")
	assert.Equal(t,
		"kind 'thing' with id 'doesnt-matter' does not exist: no peer 'WeaviateB' in the network",
		err.Error(), "should fail with a good error message")
}

func TestGetKindHappyPath(t *testing.T) {
	server := newFakeServer(t)
	peers := Peers{
		Peer{
			Name:   "WeaviateB",
			Schema: schema.Schema{},
			URI:    strfmt.URI(server.server.URL),
		},
	}
	thing := crossrefs.NetworkKind{
		Kind:     kind.THING_KIND,
		PeerName: "WeaviateB",
		ID:       "best-uuid",
	}

	happyPathHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body := models.ThingGetResponse{
			ThingID: "best-uuid",
			Thing: models.Thing{
				ThingCreate: models.ThingCreate{
					AtClass: "Instrument",
				},
			},
		}
		json.NewEncoder(w).Encode(body)
	}

	t.Run("returns no error", func(t *testing.T) {
		server.matchers = []http.HandlerFunc{happyPathHandler}
		_, err := peers.RemoteKind(thing)
		assert.Equal(t, nil, err, "should not error")
	})

	t.Run("matches the specified schema", func(t *testing.T) {
		server.matchers = []http.HandlerFunc{happyPathHandler}
		result, _ := peers.RemoteKind(thing)
		assert.Equal(t, "Instrument", result.AtClass, "found thing's schema should match")
	})

	t.Run("queries the correct path", func(t *testing.T) {
		server.matchers = []http.HandlerFunc{
			func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "GET", r.Method, "should be a GET request")
				assert.Equal(t, "/weaviate/v1/things/best-uuid", r.URL.String(),
					"should match the right path")
			}, happyPathHandler}
		peers.RemoteKind(thing)
	})
}

func newFakeServer(t *testing.T) *fakeServer {
	server := &fakeServer{t: t}
	server.server = httptest.NewServer(http.HandlerFunc(server.handle))
	return server
}

type fakeServer struct {
	t        *testing.T
	server   *httptest.Server
	matchers []http.HandlerFunc
}

func (f *fakeServer) handle(w http.ResponseWriter, r *http.Request) {
	for _, matcher := range f.matchers {
		matcher(w, r)
	}
}
