# Dockerfile for development purposes.
# Read docs/development.md for more information

# vi: ft=dockerfile

###############################################################################
# Base build image
FROM golang:1.11-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/creativesoftwarefdn/weaviate
ENV GO111MODULE=on
# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

###############################################################################
# This image builds the old acceptance testss
FROM build_base AS acceptance_test
COPY . .
ENTRYPOINT ["go", "test", "./test/full_test.go"]

###############################################################################
# This image builds the new acceptance testss
FROM build_base AS new_acceptance_test
COPY . .
ENTRYPOINT ["go", "test", "./test/acceptance"]


###############################################################################
# This image builds the weavaite server
FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/weaviate-server

###############################################################################
# This image builds the contextionary fixtures.
FROM alpine:3.8 AS contextionary_fixture_builder
RUN apk add --no-cache curl
RUN apk add --no-cache jq
RUN apk add --no-cache bash
COPY ./tools/download_latest_contextionary.sh ./tools/
RUN ./tools/download_latest_contextionary.sh

###############################################################################
# This creates an image that can be run to import the demo dataset for development
FROM build_base AS data_importer
COPY . .
ENTRYPOINT ["./tools/dev/import_demo_data.sh"]

###############################################################################
# This image builds the contextionary fixtures FOR DEV OR TEST.
FROM build_base AS contextionary_fixture_builder_dev
COPY . .
RUN ./test/contextionary/gen_simple_contextionary.sh

###############################################################################
# This is the base image for running waviates configurations; contains the executable & contextionary
FROM alpine AS weaviate_base
COPY --from=server_builder /go/bin/weaviate-server /bin/weaviate
COPY --from=build_base /etc/ssl/certs /etc/ssl/certs
COPY --from=contextionary_fixture_builder ./contextionary/contextionary.idx /contextionary/contextionary.idx
COPY --from=contextionary_fixture_builder ./contextionary/contextionary.knn /contextionary/contextionary.knn
ENTRYPOINT ["/bin/weaviate"]

###############################################################################
# This is the base image for running waviates configurations IN DEV OR TEST; contains the executable & contextionary
FROM alpine AS weaviate_base_dev
COPY --from=server_builder /go/bin/weaviate-server /bin/weaviate
COPY --from=build_base /etc/ssl/certs /etc/ssl/certs
COPY --from=contextionary_fixture_builder_dev /go/src/github.com/creativesoftwarefdn/weaviate/test/contextionary/example.idx /contextionary/contextionary.idx
COPY --from=contextionary_fixture_builder_dev /go/src/github.com/creativesoftwarefdn/weaviate/test/contextionary/example.knn /contextionary/contextionary.knn
ENTRYPOINT ["/bin/weaviate"]

###############################################################################
# Development configuration with demo dataset
FROM weaviate_base_dev AS development
COPY ./tools/dev/schema /schema
COPY ./tools/dev/config.json /weaviate.conf.json
CMD [ "--host", "0.0.0.0", "--port", "8080", "--scheme", "http", "--config", "janusgraph_docker"]

###############################################################################
# Configuration used for the acceptance tests.
FROM weaviate_base_dev AS test
COPY ./test/schema/test-action-schema.json /schema/actions_schema.json
COPY ./test/schema/test-thing-schema.json /schema/things_schema.json
COPY ./tools/dev/config.json /weaviate.conf.json
CMD [ "--host", "0.0.0.0", "--port", "8080", "--scheme", "http", "--config", "janusgraph_docker"]
