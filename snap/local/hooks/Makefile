HOOKS=cmd/configure/configure cmd/install/install

tidy:
	go mod tidy

build: tidy $(HOOKS)

cmd/configure/configure:
	go build -o $@ ./cmd/configure

cmd/install/install:
	go build -o $@ ./cmd/install

# TODO: fixup unit testing
test:
	go test -coverprofile=coverage.out ./...
	go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]

clean:
	rm -f $(HOOKS)

