# App Service Configurable

App-Service-Configurable is provided as an easy way to get started with processing data flowing through EdgeX. This service leverages the [App Functions SDK](https://github.com/edgexfoundry/app-functions-sdk-go) and provides a way for developers to use configuration instead of having to compile standalone services to utilize built in functions in the SDK. For a full list of supported/built-in functions view the [README](https://github.com/edgexfoundry/app-functions-sdk-go) located in the App Functions SDK repository. 

<!--ts-->

- [Getting Started](#getting-started)
- [Environment Variable Overrides For Docker](#environment-variable-overrides-for-docker)
- [Deploying Multiple Instances](#deploying-multiple-instances)
- [Input Data Not An EdgeX Event](#input-data-not-an-edgex-event)

 <!--te-->

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

## Environment Variable Overrides For Docker

App Service Configurable no longer has docker specific profiles. It now relies on environment variable overrides in the docker compose files for the docker specific differences. The following environment settings are required in the compose files when using App Service Configurable.

```
edgex_registry: consul://edgex-core-consul:8500
edgex_profile : [target profile]
edgex_service : http://[service name]:[port]
Service_Host : [service name]
Clients_CoreData_Host: edgex-core-data
Clients_Logging_Host : edgex-support-logging
Logging_EnableRemote: "true"
Database_Host : edgex-mongo
Database_Username : appservice
Database_Password : password
```

The following is an example docker compose entry for **App Service Configurable**:

```yaml
  app-service-configurable-rules:
    image: edgexfoundry/docker-app-service-configurable:1.1.0
    environment:
      edgex_registry: consul://edgex-core-consul:8500
      edgex_service: http://edgex-app-service-configurable-rules:48096
      edgex_profile: rules-engine
      Service_Host: edgex-app-service-configurable-rules
      Clients_CoreData_Host: edgex-core-data
      Clients_Logging_Host : edgex-support-logging
      Logging_EnableRemote: "true"      
      MessageBus_SubscribeHost_Host: edgex-core-data
    ports:
      - "48096:48096"
    container_name: edgex-app-service-configurable-rules
    hostname: edgex-app-service-configurable-rules
    networks:
      edgex-network:
        aliases:
          - edgex-app-service-configurable-rules
    depends_on:
      - data
      - command
```

> *Note: **App Service Configurable** is designed to be run multiple times each with different profiles. This is why in the above example the name `edgex-app-service-configurable-rules` is used for the instance running the `rules-engine` profile.*

## Deploying Multiple Instances using profiles

App Service Configurable was designed to be deployed as multiple instances with different purposes. Since the function pipeline is specified in the `configuration.toml` file, we can use this as a way to run each instance with a different function pipeline. App Service Configurable does not have the standard default configuration at `/res/configuration.toml`. This default configuration has been moved to the `sample` profile. This forces you to specify the profile for the configuration you would like to run. The profile is specified using the `-p/-profile=[profilename]` command line option or the `edgex_profile=[profilename]` environment variable override. The profile name selected is used in the service key (`AppService-[profile name]`) to make each instance unique, e.g. `AppService-sample` when specifying `sample` as the profile.

The following profiles and their purposes are provided with App Service Configurable. 

- **blackbox-tests** - Profile used for black box testing the SDK  
- **http-export** - Starter profile used for exporting data via HTTP. 
  Requires further configuration which can easily be accomplished using environment variable overrides
  - Required:
    - `Writable_Pipeline_Functions_HTTPPostJSON_Parameters_Url:[Your URL]`
  - Optional: 
    - `Writable_Pipeline_Functions_HTTPPostJSON_Parameters_PersistOnError:["true"/"false"]`
  - Optional: 
    - `Writable_Pipeline_Functions_FilterByDeviceName_Parameters_DeviceNames:"[comma separated list]"`
  - Optional: 
    - `Writable_Pipeline_Functions_FilterByValueDescriptor_Parameters_ValueDescriptors: "[comma separated list]"`
- **mqtt-export** - Starter profile used for exporting data via MQTT.
  Requires further configuration which can easily be accomplished using environment variable overrides
  - Required:
    - `Writable_Pipeline_Functions_MQTTSend_Addressable_Address:[Your Address]`
  - Optional: 
    - `Writable_Pipeline_Functions_MQTTSend_Addressable`
      - `_Port:["your port"]`
      - `_Protocol:[tcp or tcps]`  
      - `_Publisher:[your name]`
      - `_User:[your username]`
      - `_Password:[your passowrd`
      - `_Topic:[your topic]`
  - Optional: 
    - `Writable_Pipeline_Functions_MQTTSend_Parameters`
      - `_Qos:["your quality or service"]`
      - `_Key:[your Key]`  
      - `_Cert:[your Certificate]`
      - `_Autoreconnect:["true" or "false"]`
      - `_Retain:["true" or "false"]`
      - `_PersistOnError:["true" or "false"]`
- **rules-engine** - Profile used to push Event messages to the Rules Engine via ZMQ.
- **sample** - Sample profile with all available functions declared and a sample pipeline. Provided as a sample that can be copy and modified to create new custom profiles.

> *Note: Running multiple instances with the same profile undermines this design and is not recommended.*

> *Note: Functions can be declared in a profile but not used in the pipeline `ExecutionOrder` allowing them to be added to the pipeline `ExecutionOrder` later at runtime if needed.*

## Input Data Not An EdgeX Event

What if my input data isn't an EdgeX Event ?

The default `TargetType` for data flowing into the functions pipeline is an EdgeX event. There are cases when this incoming data might not be an EdgeX event. In these cases the `Pipeline` can be configured using `UseTargetTypeOfByteArray=true` to set the `TargetType` to be a byte array, i.e. `byte[]`. The first function in the pipeline must then be one that can handle the `byte[]`data. The **compression**,  **encryption** and **export** functions are examples of pipeline functions that will take input data that is `byte[]`. Here is an example of how to configure the functions pipeline to **compress**, **encrypt** and then **export** the  `byte[]` data via HTTP.

```toml
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

If along with this pipeline configuration, you also configured the `Binding` to be `http` trigger,  you could then send any data to the app-service-configurable' s `/api/v1/trigger` endpoint and have it compressed, encrypted and sent to your configured URL above.

```
[Binding]
Type="http"
```

