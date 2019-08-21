# App Service Configurable

App-Service-Configurable is provided as an easy way to get started with processing data flowing through EdgeX. This service leverages the [App Functions SDK](https://github.com/edgexfoundry/app-functions-sdk-go) and provides a way for developers to use configuration instead of having to compile standalone services to utilize built in functions in the SDK. For a full list of supported/built-in functions view the [README](https://github.com/edgexfoundry/app-functions-sdk-go) located in the App Functions SDK repository. 

# Getting Started 

To get started with the configurable app service, you'll want to start by determining which functions are required in your pipeline. Using a simple example.
let's assume you wish to use the following functions from the SDK:

1) [FilterByDeviceName](https://github.com/edgexfoundry/app-functions-sdk-go#filtering) - to filter events for a specific device.
2) [TransformToXML](https://github.com/edgexfoundry/app-functions-sdk-go#conversion) - to transform the data to XML
3) [HTTPPost](https://github.com/edgexfoundry/app-functions-sdk-go#export-functions) - to send the data to an HTTP endpoint that takes our XML data
4) [MarkAsPushed](https://github.com/edgexfoundry/app-functions-sdk-go/blob/master/README.md#CoreData-Functions) - to call Core Data API to mark the event as having been pushed

Once the functions have been identified, we'll go ahead and build out the configuration in the `configuration.toml` file under the `[Writable.Pipeline]` section:

```toml
[Writable]
  LogLevel = 'DEBUG'
  [Writable.Pipeline]
    ExecutionOrder = "FilterByDeviceName,TransformToXML,HTTPPost,MarkAsPushed"
    [Writable.Pipeline.Functions.FilterByDeviceName]
      [Writable.Pipeline.Functions.FilterByDeviceName.Parameters]
        FilterValues = "Random-Float-Device,Random-Integer-Device"
    [Writable.Pipeline.Functions.TransformToXML]
    [Writable.Pipeline.Functions.MarkAsPushed]
    [Writable.Pipeline.Functions.HTTPPost]
      [Writable.Pipeline.Functions.HTTPPost.Parameters]
        url = "http://my.api.net/edgexdata"
        mimeType = "" #OPTIONAL - default application/json
```

The first line of note is `ExecutionOrder = "FilterByDeviceName,TransformToXML,HTTPPost,MarkAsPushed"`. This specifies the order in which to execute your functions. Each function specified here must also be placed in the `[Writeable.Pipeline.Functions]` section. 

Next, each function and its required information is listed. Each function typically has associated Parameters that must be configured to properly execute the function as designated by `[Writable.Pipeline.Functions.{FunctionName}.Parameters]`. Knowing which parameters are required for each function, can be referenced by taking a look at the documentation located [here](https://github.com/edgexfoundry/app-functions-sdk-go#built-in-transformsfunctions).
In a few cases, such as `TransformToXML`, `TransformToJSON`, or `SetOutputData`, there are no parameters required.


> Note: By default, the configuration provided is set to use MessageBus as a trigger from CoreData. This means you must have EdgeX Running with devices sending data in order to trigger the pipeline. You can also change the trigger to be HTTP. For more on triggers, view the documentation [here](https://github.com/edgexfoundry/app-functions-sdk-go#triggers).

That's it! Now we can run/deploy this service and the functions pipeline will process the data with functions we've defined.

## What if I want to deploy multiple pipelines using this service?

It not uncommon to have different sets of function pipelines that need to be deployed. This is where `--profiles` comes in handy. You can create different "profile" folders inside the `/res` directory with different configurations. This is typically used in other EdgeX services for having a separate configuration for Docker based deployments and native deployments. Since the pipeline is specified in the `configuration.toml` file, we can use this as a way to run the `app-service-configurable` application with different profiles by specifying the `"--profile=http-export"` and `"--confdir=/res"` as command line options when deploying. If running with Registry enabled, the service key used will be `AppService-[profile name]`, e.g `AppService-http-export` when using --profile option or just `AppService` if not using the --profile option.

## What if my input data isn't an EdgeX Event ?

The default `TargetType` for data flowing into the functions pipeline is an EdgeX event. There are cases when this incoming data might not be an EdgeX event. In these cases the `Pipeline` can be configured using `UseTargetTypeOfByteArray=true` to set the `TargetType` to be a byte array, i.e. `byte[]`. The first function in the pipeline must then be one that can handle the `byte[]`data. The **compression**,  **encryption** and **export** functions are examples of pipeline functions that will take input data that is `byte[]`. Here is an example of how to configure the functions pipeline to **compress**, **encrypt** and then **export** the  `byte[]` data via HTTP.

```
[Writable]
  LogLevel = 'DEBUG'
  [Writable.Pipeline]
    UseTargetTypeOfByteArray = true
    ExecutionOrder = "CompressWithGZIP, EncryptWithAES, HTTPPost"
    [Writable.Pipeline.Functions.CompressWithGZIP]
    [Writable.Pipeline.Functions.EncryptWithAES]
      [Writable.Pipeline.Functions.EncryptWithAES.Parameters]
        Key = "aquqweoruqwpeoruqwpoeruqwpoierupqoweiurpoqwiuerpqowieurqpowieurpoqiweuroipwqure"
        InitVector = "123456789012345678901234567890"
    [Writable.Pipeline.Functions.HTTPPost]
      [Writable.Pipeline.Functions.HTTPPost.Parameters]
        url = "http://my.api.net/edgexdata"
```

If along with this pipeline configuration, you also configured the `Binding` to be `http` trigger,  you could then send any data to the app-service-configurable's `/trigger` endpoint and have it compressed, encrypted and sent to your configured URL above.

```
[Binding]
Type="http"
```

