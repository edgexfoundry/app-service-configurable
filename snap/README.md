# EdgeX Foundry App Service Configurable Snap
[![snap store badge](https://raw.githubusercontent.com/snapcore/snap-store-badges/master/EN/%5BEN%5D-snap-store-black-uneditable.png)](https://snapcraft.io/edgex-app-service-configurable)

This folder contains snap packaging for the EdgeX Foundry's App Service Configurable application service.

The project maintains a rolling release of the snap on the `edge` channel that is rebuilt and published at least once daily.

The snap currently supports both `amd64` and `arm64` platforms.

## Installation

### Installing snapd
The snap can be installed on any system that supports snaps. You can see how to install 
snaps on your system [here](https://snapcraft.io/docs/installing-snapd/6735).

However for full security confinement, the snap should be installed on an 
Ubuntu 18.04 LTS or later (Desktop or Server), or a system running Ubuntu Core 18 or later.

### Installing EdgeX App Service Configurable as a snap
The snap is published in the snap store as [edgex-app-service-configurable](https://snapcraft.io/edgex-app-service-configurable).
See the current revisions available for your machine's architecture:

```bash
snap info edgex-app-service-configurable
```

The latest stable version of the snap can be installed using:

```bash
sudo snap install edgex-app-service-configurable
```

A specific version of the snap can be installed by setting the channel, 
for instance for 2.1 (Jakarta):

```bash
sudo snap install edgex-app-service-configurable --channel=2.1
```

The latest development version of the snap can be installed using:

```bash
sudo snap install edgex-app-service-configurable --edge
```

**Note** - the snap has only been tested on Ubuntu Core, Desktop, and Server.

## Using the EdgeX App Service Configurable snap

The App Service Configurable application service allows a variety of use cases to be met by simply providing configuration. 

This service is disabled when first installed, 
as a profile must first be selected before the service is started. 
See [profile](#profiles) section below.

For more information about this service, 
please refer to the [EdgeX App Service Configurable README](../README.md).

### Profiles
There are a number of sub-directories in the snap's configuration directory, 
these sub-directories are referred to as profiles. 
Before you can start the service, you must select one of them, 
using the `snap set` command (or via [snapd's REST API](https://snapcraft.io/docs/snapd-api). 
For example, to use the mqtt-export profile you would run:

```bash
sudo snap set edgex-app-service-configurable profile=mqtt-export
```

In addition to instructing the service to read a different configuration file, 
the profile will also be used to name the service when it registers itself to the system.
For details on currently available profiles, please see [here](../res).

The `configuration.toml` files are found in the snap’s writable area: 
```
/var/snap/edgex-app-service-configurable/current/config/res/
```

The service can then be started as follows. The `--enable` option
ensures that as well as starting the service now, it will be automatically started on boot:
```bash
sudo snap start --enable edgex-app-service-configurable
```
**Note** - Should the environment variables be modified after the service has started, 
the service must be restarted.

### Configuration Overrides
While it's possible to manually edit the profile-specific `configuration.toml` files (found in `$SNAP_DATA/config/res/<profile>`)
prior to enabling the service, quite often only minor changes to existing profiles are required. These changes can be accomplished via
support for [EdgeX environment variable configuration overrides](#service-environment-configuration-overrides) via the snap's configure hook. 
If the service has already been started,
setting one of these overrides currently requires the service to be restarted via the command-line or snapd's REST API. 
If the overrides are provided via the snap configuration defaults capability of a [gadget snap](https://snapcraft.io/docs/gadget-snap), 
the overrides will be picked up when the services are first started.

The following syntax is used to specify service-specific configuration overrides:

```
env.<stanza>.<config option>
```

For instance, to setup an override of the service's port:

```bash
sudo snap set env.service.port=2112
```

And restart the service:

```bash
sudo snap restart edgex-app-service-configurable
```

**Note** - at this time changes to configuration values in the [Writable] section are not supported.

For details on the mapping of configuration options to Config options, 
please refer to [Service Environment Configuration Overrides](#service-environment-configuration-overrides) section below.

### Using a content interface to set app configuration

The `app-config` content interface allows another snap to seed this application snap 
with configuration files under the `$SNAP_DATA/config` directory.

Note that the `app-config` content interface does NOT support seeding of the Secret Store Token
because that file is expected at a different path.

Please refer to [edgex-config-provider](https://github.com/canonical/edgex-config-provider), for an example and further instructions.

### Autostart
By default, the edgex-app-service-configurable disables its service on install, 
as the expectation is that the default profile configuration files will be customized, 
and thus this behavior allows the profile `configuration.toml` files in [`$SNAP_DATA`](https://snapcraft.io/docs/environment-variables) to be modified before the service is first started.

This behavior can be overridden by setting the `autostart` configuration setting to `true`. 
This is useful when configuration and/or device profiles are being provided via configuration or gadget snap content interface.

**Note** - this option is typically set from a gadget snap.


## Service Environment Configuration Overrides
Application services implement a service dependency check on startup which ensures that all of the runtime dependencies of a particular service are met before the service transitions to an active state.

Snapd doesn't support orchestration between services in different snaps. 
It is therefore possible on a reboot for an application service to come up faster than all of the required services running in the main edgexfoundry snap. 
If this happens, the application service may repeatedly fail startup, 
and if it exceeds the systemd default limits, 
then it might be left in a failed state. 
This situation might be more likely on constrained hardware (e.g. RPi).

This snap, therefore, implements a basic retry loop with a maximum `startup-duration` and `startup-interval`. 
If the dependent services are not available, 
the service sleeps for the defined interval and then tries again up to a maximum duration. 
EdgeX services wait for dependencies (e.g. core-data) to become available and will exit after reaching the maximum duration if the dependencies aren't met. 
    
This environment variable sets the total duration in seconds allowed for the services to complete the bootstrap start-up. 
Default is 60 seconds. Change the maximum duration:

```bash
sudo snap set edgex-app-service-configurable startup-duration=120
```

This environment variable sets the retry interval in seconds for the services retrying a failed action during the bootstrap start-up.
Default is 1 second. Change the interval between retries:

```bash
sudo snap set edgex-app-service-configurable startup-interval=3
```

**Note** - all of the configuration options below must be specified with the prefix: `env.`

```
[Service]
service.health-check-interval   // Service.HealthCheckInterval
service.host                    // Service.Host
service.server-bind-addr        // Service.ServerBindAddr
service.port                    // Service.Port
service.max-result-count        // Service.MaxResultCount
service.max-request-size        // Service.MaxRequestSize
service.startup-msg             // Service.StartupMsg
service.request-timeout         // Service.RequestTimeout

[SecretStore]
secret-store.secrets-file               // SecretStore.SecretsFile
secret-store.disable-scrub-secrets-file // SecretStore.DisableScrubSecretsFile

[Clients.core-command]
clients.core-command.port       // Clients.core-command.Port

[Clients.core-data]
clients.core-data.port          // Clients.core-data.Port

[Clients.core-metadata]
clients.core-metadata.port      // Clients.core-metadata.Port

[Clients.support-notifications]
clients.support-notifications.port  // Clients.support-notifications.Port

[Triger]
[Trigger.EdgexMessageBus]
trigger.edgex-message-bus.type                             // Trigger.EdgexMessageBus.Type

[Trigger.EdgexMessageBus.SubscribeHost]
trigger.edgex-message-bus.subscribe-host.port              // Trigger.EdgexMessageBus.SubscribeHost.Port
trigger.edgex-message-bus.subscribe-host.protocol          // Trigger.EdgexMessageBus.SubscribeHost.Protocol
trigger.edgex-message-bus.subscribe-host.subscribe-topics  // Trigger.EdgexMessageBus.SubscribeHost.SubscribeTopics

[Trigger.EdgexMessageBus.PublishHost]
trigger.edgex-message-bus.publish-host.port                // Trigger.EdgexMessageBus.PublishHost.Port
trigger.edgex-message-bus.publish-host.protocol            // Trigger.EdgexMessageBus.PublishHost.Protocol
trigger.edgex-message-bus.publish-host.publish-topic       // Trigger.EdgexMessageBus.PublishHost.PublishTopic
```

