.PHONY: build test clean docker

GO=CGO_ENABLED=1 go

VERSION=$(shell cat ./VERSION)

MICROSERVICE=app-service-configurable
GOFLAGS=-ldflags "-X $(MICROSERVICE).Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: 
	$(GO) build -o $(MICROSERVICE)

docker:
	docker build \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-app-service-configurable:$(GIT_SHA) \
		-t edgexfoundry/docker-app-service-configurable:$(VERSION)-dev \
		.

test:
	$(GO) test ./... -cover
	$(GO) vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]

clean:
	rm -f $(MICROSERVICE)

