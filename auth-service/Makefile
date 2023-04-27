export RELEASE_VERSION	?= $(shell git show -q --format=%h)
export DOCKER_REGISTRY	?= auth-service
export DEPLOY			?= image

all: build start
db-migrate:
	rel migrate
db-rollback:
	rel rollback
gen:
	go generate ./...
build: gen
	go build -o bin/authApp ./cmd/api
test: gen
	go test -race ./...
start:
	./bin/authApp
env:
	cp .env.sample .env
docker:
	docker build -t $(DOCKER_REGISTRY)/$(DEPLOY):$(RELEASE_VERSION) -f ./deploy/$(DEPLOY)/Dockerfile .
push:
	docker push $(DOCKER_REGISTRY)/$(DEPLOY):$(RELEASE_VERSION)