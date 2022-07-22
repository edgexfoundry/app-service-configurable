
build:
	go build -ldflags="-s -w" -o helper-go

tidy:
	go mod tidy

clean:
	rm -f helper-go

# TODO: fixup unit testing
test:
	go test -coverprofile=coverage.out ./...
	go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]