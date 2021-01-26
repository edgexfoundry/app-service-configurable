# EdgeX Foundry App Service Configurable Snap
[![snap store badge](https://raw.githubusercontent.com/snapcore/snap-store-badges/master/EN/%5BEN%5D-snap-store-black-uneditable.png)](https://snapcraft.io/edgex-app-service-configurable)

This folder contains snap packaging for the EdgeX Foundry's App Service Configurable application service.

The project maintains a rolling release of the snap on the `edge` channel that is rebuilt and published at least once daily through the jenkins jobs setup for the EdgeX project. You can see the jobs run [here](https://jenkins.edgexfoundry.org/view/Snap/) specifically looking at the `edgex-app-service-configurable-snap-{branch}-stage-snap`.

The snap currently supports both `amd64` and `arm64` platforms.

## Installation

### Installing snapd
The snap can be installed on any system that supports snaps. You can see how to install 
snaps on your system [here](https://snapcraft.io/docs/installing-snapd/6735).

However for full security confinement, the snap should be installed on an 
Ubuntu 16.04 LTS or later Desktop or Server, or a system running Ubuntu Core 16 or later.

### Installing EdgeX App Service Configurable as a snap
The snap is published in the snap store at https://snapcraft.io/edgex-app-service-configurable.
You can see the current revisions available for your machine's architecture by running the command:

```bash
$ snap info edgex-app-service-configurable
```

The latest stable version of the snap can be installed using:

```bash
$ sudo snap install edgex-app-service-configurable
```

**Note** - this snap has only been tested on Ubuntu Desktop/Server 16.04 LTS (or greater) and Ubuntu Core 16 (or greater).

## Using the EdgeX App Service Configurable snap

The App Service Configurable application service allows a variety of use cases to be met by simply providing configuration (vs. writing code). For more information about this service, please refer to the README. As with device-mqtt, this service is disabled when first installed, as it requires configuration changes before it can be run. As with the device-mqtt snap, the configuration.toml file is found in the snap’s writable area:


/var/snap/edgex-app-service-configurable/current/config/res/

### Startup environment variables

EdgeX services by default wait 60s for dependencies (e.g. Core Data) to become available, and will exit after this time if the dependencies aren't met. The following options can be used to override this startup behavior on systems where it takes longer than expected for the dependent services provided by the edgexfoundry snap to start. Note, both options below are specified as a number of seconds.
    
To change the default startup duration (60 seconds), for a service to complete the startup, aka bootstrap, phase of execution by using the following command:

```bash
$ sudo snap set edgex-app-service-configurable startup-duration=60
```

The following environment variable overrides the retry startup interval or sleep time before a failure is retried during the start-up, aka bootstrap, phase of execution by using the following command:

```bash
$ sudo snap set edgex-app-service-configurable startup-interval=1
```

**Note** - Should the environment variables be modified after the service has started, the service must be restarted.

### Profiles
In additional to base configuration.toml in this directory, there are a number of sub-directories that also contain configuration.toml files. These sub-directories are referred to as profiles. The service’s default behavior is to use the configuration.toml file from the /res directory. If you want to use one of the profiles, use the snap set command to instruct the service to read its configuration from one of these sub-directories. For example, to use the push-to-core profile you would run:
```
$ sudo snap set edgex-app-service-configurable profile=push-to-core
```
In addition to instructing the service to read a different configuration file, the profile will also be used to name the service when it registers itself to the system.

### Multiple Instances
Multiple instances of edgex-app-service-configurable can be installed by using snap [Parallel Installs](https://snapcraft.io/docs/parallel-installs). This is an experimental snap feature and must be first be enabled by running this command:
```
sudo snap set system experimental.parallel-instances=true
```
Now you can install multiple instances of the edgex-app-service-configurable snap by specifying a unique instance name when you install the snap. The instance name is a unique suffix which is appended to the snap name following the “_” character used as a delimeter. This name only needs to be specified for the second and susbequent instances of the snap.
```
sudo snap install edgex-app-service-configurable edgex-app-service-configurable_http
```
or
```
sudo snap install edgex-app-service-configurable edgex-app-service-configurable_mqtt
```
**Note** – you must ensure that any configuration values that might cause conflict between the multiple instances (e.g. port, log file path, …) must be modified before enabling the snap’s service.

### Secret Store Usage
Some profile configuration.toml files specify configuration which requires Secret Store (aka Vault) support, however this snap doesn't fully support secure secrets without manual intervention (see below). The snap can also be configured to use insecure secrets as can be done via docker-compose by setting the option ```security-secret-store=false```. Ex.

```
sudo snap set edgex-app-service-configurable security-secret-store=false
```

### Manually configure the Secret Store (aka Vault) token
Here's an example of how to use a configuration/profile which includes secrets:

```
$ sudo snap install edgex-app-service-configurable
$ sudo snap set edgex-app-service-configurable profile=mqtt-export
$ cd /var/snap/edgex-app-service-configurable/current/config/res/mqtt-export
$ sudo cp /var/snap/edgexfoundry/current/secrets/edgex-application-service/secrets-token.json .
```

Next the profile's ```configuration.toml``` file (see ```/var/snap/edgex-app-service-configurable/current/config/res/mqtt-export```) needs to be updated to reference the new token file location in ```$SNAP_DATA```. The config field that need updated is ```SecretStore.TokenFile```. The following example shows how this can be done via the sed command-line tool, however the file can also be easily updated via your favorite editor.


```
$ sudo sed -i -e 's@/vault/config/assets/resp-init.json@/var/snap/edgex-app-service-configurable/current/config/res/mqtt-export/secrets-token.json@' ./configuration.toml
```

**Note** -- these configuration changes need to be made *before* the service is started for the first time. Otherwise, the recommended approach is to stop the service, delete the existing app-service-configurable configuration in Consul's kv store, and then proceed.
