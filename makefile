.PHONY: build test clean docker

GO=CGO_ENABLED=1 go

APPVERSION=$(shell cat ./VERSION)

# This pulls the version of the SDK from the go.mod file. If the SDK is the only required module,
# it must first remove the word 'required' so the offset of $2 is the same if there are multiple required modules
SDKVERSION=$(shell cat ./go.mod | grep 'github.com/edgexfoundry/app-functions-sdk-go v' | sed 's/require//g' | awk '{print $$2}')

MICROSERVICE=app-service-configurable
GOFLAGS=-ldflags "-X github.com/edgexfoundry/app-functions-sdk-go/internal.SDKVersion=$(SDKVERSION) -X github.com/edgexfoundry/app-functions-sdk-go/internal.ApplicationVersion=$(APPVERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build:
	$(GO) build $(GOFLAGS) -o $(MICROSERVICE)

docker:
	docker build \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-app-service-configurable:$(GIT_SHA) \
		-t edgexfoundry/docker-app-service-configurable:$(VERSION)-dev \
		.
version:
	echo $(SDKVERSION)

test:
	$(GO) test ./... -cover
	$(GO) vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]

clean:
	rm -f $(MICROSERVICE)

