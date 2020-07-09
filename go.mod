module github.com/semi-technologies/weaviate

require (
	github.com/TylerBrock/colorjson v0.0.0-20180527164720-95ec53f28296
	github.com/bmatcuk/doublestar v1.1.3
	github.com/boltdb/bolt v1.3.1
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-oidc v2.0.0+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elastic/go-elasticsearch/v5 v5.6.0
	github.com/etiennedi/go-tsne v0.0.0-20200514072915-0785d987ebe8
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/color v1.7.0 // indirect
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/runtime v0.19.19
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.0
	github.com/graphql-go/graphql v0.7.7
	github.com/hokaccha/go-prettyjson v0.0.0-20190818114111-108c894c2c0e // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nyaruka/phonenumbers v1.0.54
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.5.0
	github.com/satori/go.uuid v0.0.0-20180103174451-36e9d2ebbde5
	github.com/semi-technologies/contextionary v0.0.0-20200622144407-0a253c7254f2
	github.com/sirupsen/logrus v1.4.2
	github.com/square/go-jose v2.3.0+incompatible
	github.com/stretchr/testify v1.6.1
	github.com/ugorji/go/codec v0.0.0-20190309163734-c4a1c341dc93
	github.com/vektah/gqlparser v1.1.2 // indirect
	go.mongodb.org/mongo-driver v1.3.5 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/tools v0.0.0-20200708183856-df98bc6d456c // indirect
	gonum.org/v1/gonum v0.7.0
	google.golang.org/grpc v1.24.0
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v2 v2.3.0
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

go 1.14
