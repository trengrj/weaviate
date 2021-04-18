#!/usr/bin/env bash

CONFIG=${1:-local-development}

# Jump to root directory
cd "$( dirname "${BASH_SOURCE[0]}" )"/../..

export GO111MODULE=on
export DEVELOPMENT_UI=on
export LOG_LEVEL=trace
export LOG_FORMAT=text

case $CONFIG in
  local-development)
      CONTEXTIONARY_URL=localhost:9999 \
      QUERY_DEFAULTS_LIMIT=20 \
      ORIGIN=http://localhost:8080 \
      AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true \
      DEFAULT_VECTORIZER_MODULE=text2vec-contextionary \
      PERSISTENCE_DATA_PATH="./data" \
      ENABLE_MODULES="text2vec-contextionary" \
      go run ./cmd/weaviate-server \
        --scheme http \
        --host "127.0.0.1" \
        --port 8080 \
        --read-timeout=600s \
        --write-timeout=600s
    ;;
  local-transformers)
      CONTEXTIONARY_URL=localhost:9999 \
      QUERY_DEFAULTS_LIMIT=20 \
      ORIGIN=http://localhost:8080 \
      AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true \
      DEFAULT_VECTORIZER_MODULE=text2vec-transformers \
      PERSISTENCE_DATA_PATH="./data" \
      TRANSFORMERS_INFERENCE_API="http://localhost:8000" \
      ENABLE_MODULES="text2vec-transformers" \
      go run ./cmd/weaviate-server \
        --scheme http \
        --host "127.0.0.1" \
        --port 8080 \
        --read-timeout=600s \
        --write-timeout=600s
    ;;
  local-qna)
      CONTEXTIONARY_URL=localhost:9999 \
      QUERY_DEFAULTS_LIMIT=20 \
      ORIGIN=http://localhost:8080 \
      AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true \
      DEFAULT_VECTORIZER_MODULE=text2vec-contextionary \
      PERSISTENCE_DATA_PATH="./data" \
      QNA_INFERENCE_API="http://localhost:8001" \
      ENABLE_MODULES="text2vec-contextionary,qna-transformers" \
      go run ./cmd/weaviate-server \
        --scheme http \
        --host "127.0.0.1" \
        --port 8080 \
        --read-timeout=600s \
        --write-timeout=600s
    ;;
  local-oidc)
      CONTEXTIONARY_URL=localhost:9999 \
      QUERY_DEFAULTS_LIMIT=20 \
      ORIGIN=http://localhost:8080 \
      PERSISTENCE_DATA_PATH="./data" \
      AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=false \
      AUTHENTICATION_OIDC_ENABLED=true \
      AUTHENTICATION_OIDC_ISSUER=http://localhost:9090/auth/realms/weaviate \
      AUTHENTICATION_OIDC_USERNAME_CLAIM=email \
      AUTHENTICATION_OIDC_GROUPS_CLAIM=groups \
      AUTHENTICATION_OIDC_CLIENT_ID=demo \
      AUTHORIZATION_ADMINLIST_ENABLED=true \
      AUTHORIZATION_ADMINLIST_USERS=john@doe.com \
      DEFAULT_VECTORIZER_MODULE=text2vec-contextionary \
      go run ./cmd/weaviate-server \
        --scheme http \
        --host "127.0.0.1" \
        --port 8080
    ;;

  *) 
    echo "Invalid config" 2>&1
    exit 1
    ;;
esac

