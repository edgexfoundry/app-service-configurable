.PHONY: build tidy docker test clean vendor

# change the following boolean flag to enable or disable the Full RELRO (RELocation Read Only) for linux ELF (Executable and Linkable Format) binaries
ENABLE_FULL_RELRO=true
# change the following boolean flag to enable or disable PIE for linux binaries which is needed for ASLR (Address Space Layout Randomization) on Linux, the ASLR support on Windows is enabled by default
ENABLE_PIE=true

# VERSION file is not needed for local development, In the CI/CD pipeline, a temporary VERSION file is written
# if you need a specific version, just override below
APPVERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)

# This pulls the version of the SDK from the go.mod file. If the SDK is the only required module,
# it must first remove the word 'required' so the offset of $2 is the same if there are multiple required modules
# the last awk command removes the leading 'v' from the version
SDKVERSION=$(shell cat ./go.mod | grep 'github.com/edgexfoundry/app-functions-sdk-go/v4 v' | sed 's/require//g' | awk '{print $$2}' | awk '!x{x=sub("v","")}1')

ifeq ($(ENABLE_FULL_RELRO), true)
	ENABLE_FULL_RELRO_GOFLAGS = -bindnow
endif

MICROSERVICE=app-service-configurable
GOFLAGS=-ldflags "-s -w -X github.com/edgexfoundry/app-functions-sdk-go/v4/internal.SDKVersion=$(SDKVERSION) \
                   -X github.com/edgexfoundry/app-functions-sdk-go/v4/internal.ApplicationVersion=$(APPVERSION) \
                   $(ENABLE_FULL_RELRO_GOFLAGS)" \
                   -trimpath -mod=readonly
GOTESTFLAGS?=-race

GIT_SHA=$(shell git rev-parse HEAD)

ifeq ($(ENABLE_PIE), true)
	GOFLAGS += -buildmode=pie
endif

# CGO is enabled by default and causes docker builds to fail due to no gcc,
# but is required for test with -race, so must disable it for the builds only
build:
	CGO_ENABLED=0 go build -tags "$(ADD_BUILD_TAGS)" $(GOFLAGS) -o $(MICROSERVICE)

build-nats:
	make -e ADD_BUILD_TAGS=include_nats_messaging build

build-noziti:
	make -e ADD_BUILD_TAGS=no_openziti build

tidy:
	go mod tidy

# NOTE: This is only used for local development. Jenkins CI does not use this make target
docker:
	docker build \
		--build-arg ADD_BUILD_TAGS=$(ADD_BUILD_TAGS) \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/app-service-configurable:$(GIT_SHA) \
		-t edgexfoundry/app-service-configurable:${APPVERSION}-dev \
		.

docker-nats:
	make -C . -e ADD_BUILD_TAGS=include_nats_messaging docker

docker-noziti:
	make -e ADD_BUILD_TAGS=no_openziti docker

lint:
	@which golangci-lint >/dev/null || echo "WARNING: go linter not installed. To install, run make install-lint"
	@if [ "z${ARCH}" = "zx86_64" ] && which golangci-lint >/dev/null ; then golangci-lint run --config .golangci.yml ; else echo "WARNING: Linting skipped (not on x86_64 or linter not installed)"; fi

install-lint:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.61.0

unittest:
	go test $(GOTESTFLAGS) -coverprofile=coverage.out ./...

test: unittest lint
	go vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]
	./bin/test-attribution-txt.sh

clean:
	rm -f $(MICROSERVICE)

vendor:
	go mod vendor
