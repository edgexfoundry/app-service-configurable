module github.com/edgexfoundry/app-service-configurable

go 1.16

require github.com/edgexfoundry/app-functions-sdk-go/v2 v2.0.0-dev.73

replace (
	github.com/edgexfoundry/app-functions-sdk-go/v2 => ../app-functions-sdk-go
)
