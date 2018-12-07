package restapi

import (
	"crypto/tls"
	"net/http"
	"time"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/creativesoftwarefdn/weaviate/genesis/models"
	"github.com/creativesoftwarefdn/weaviate/genesis/restapi/operations"
	libstate "github.com/creativesoftwarefdn/weaviate/genesis/state"

	log "github.com/sirupsen/logrus"
)

//go:generate swagger generate server --target .. --name weaviate-genesis --spec ../openapi-spec.json --default-scheme https

func configureFlags(api *operations.WeaviateGenesisAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

var state libstate.State

func configureAPI(api *operations.WeaviateGenesisAPI) http.Handler {
	log.SetLevel(log.DebugLevel)

	state = libstate.NewInMemoryState()
	log.Info("Created in memory state")

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Infof

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GenesisPeersLeaveHandler = operations.GenesisPeersLeaveHandlerFunc(func(params operations.GenesisPeersLeaveParams) middleware.Responder {
		err := (state).RemovePeer(params.PeerID)

		if err == nil {
			return operations.NewGenesisPeersLeaveNoContent()
		} else {
			return operations.NewGenesisPeersLeaveNotFound()
		}
	})

	api.GenesisPeersPingHandler = operations.GenesisPeersPingHandlerFunc(func(params operations.GenesisPeersPingParams) middleware.Responder {
		err := state.UpdateLastContact(params.PeerID, time.Now())

		if err == nil {
			return operations.NewGenesisPeersPingOK()
		} else {
			return operations.NewGenesisPeersPingNotFound()
		}
	})

	api.GenesisPeersRegisterHandler = operations.GenesisPeersRegisterHandlerFunc(func(params operations.GenesisPeersRegisterParams) middleware.Responder {
		var err error = nil

		if err == nil {
			peer, err := (state).RegisterPeer(params.Body.PeerName, params.Body.PeerURI)
			if err != nil {
				return operations.NewGenesisPeersRegisterForbidden()
			} else {
				response_peer := models.Peer{
					PeerUpdate: models.PeerUpdate{
						PeerURI:  peer.URI(),
						PeerName: peer.Name(),
					},
					ID:            peer.Id,
					LastContactAt: peer.LastContactAt.Unix(),
				}
				payload := models.PeerRegistrationResponse{
					Peer: &response_peer,
				}
				return operations.NewGenesisPeersRegisterOK().WithPayload(&payload)
			}
		} else {
			return operations.NewGenesisPeersRegisterBadRequest()
		}
	})

	api.GenesisPeersListHandler = operations.GenesisPeersListHandlerFunc(func(params operations.GenesisPeersListParams) middleware.Responder {
		peers := make([]*models.Peer, 0)

		listed_peers, err := (state).ListPeers()
		if err != nil {
			log.Infof("Failed to list peers, because %+v", err)
			return operations.NewGenesisPeersListInternalServerError()
		} else {
			for _, peer := range listed_peers {
				p := models.Peer{
					PeerUpdate: models.PeerUpdate{
						PeerURI:  peer.URI(),
						PeerName: peer.Name(),
					},
					ID:            peer.Id,
					LastContactAt: peer.LastContactAt.Unix(),
				}

				peers = append(peers, &p)
			}

			reply := operations.NewGenesisPeersListOK()
			reply.SetPayload(peers)
			return reply
		}
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return addLogging(handler)
}

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Received request: %+v %+v", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
