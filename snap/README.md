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
The snap is published in the snap store at https://snapcraft.io/edgex-app-service-configurable.
You can see the current revisions available for your machine's architecture by running the command:

```bash
$ snap info edgex-app-service-configurable
```

The latest stable version of the snap can be installed using:

```bash
$ sudo snap install edgex-app-service-configurable
```

The 2.1 (Jakarta) release of the snap can be instaing using:

```bash
$ sudo snap install edgex-app-service-configurable --channel=2.1
```

The latest development version of the snap can be installed using:

```bash
$ sudo snap install edgex-app-service-configurable --edge
```

**Note** - the snap has only been tested on Ubuntu Core, Desktop, and Server.

## Using the EdgeX App Service Configurable snap

The App Service Configurable application service allows a variety of use cases to be met by simply providing configuration (vs. writing code). For more information about this service, please refer to the README. As with device-mqtt, this service is disabled when first installed, as a profile must first be selected (see below) before the service is started. As with other EdgeX snaps, the `configuration.toml` files are found in the snap’s writable area:

/var/snap/edgex-app-service-configurable/current/config/res/

### Configuration Overrides
While it's possible to manually edit the profile-specific ```configuration.toml``` files (found in ```$SNAP_DATA/config/res/<profile>```)
prior to enabling the service, quite often only minor changes to existing profiles are required. These changes can be accomplished via
support for EdgeX environment variable configuration overrides via the snap's configure hook. If the service has already been started,
setting one of these overrides currently requires the service to be restarted via the command-line or snapd's REST API. If the overrides
are provided via the snap configuration defaults capability of a gadget snap, the overrides will be picked up when the services are first
started.

The following syntax is used to specify service-specific configuration overrides:

```env.<stanza>.<config option>```

For instance, to setup an override of the service's Port use:

```$ sudo snap set env.service.port=2112```

And restart the service:

```$ sudo snap restart edgex-app-service-configurable```

**Note** - at this time changes to configuration values in the [Writable] section are not supported.

For details on the mapping of configuration options to Config options, please refer to "Service Environment Configuration Overrides".

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
There are a number of sub-directories in the snap's configuration directory, these sub-directories are referred to as profiles. Before you can start the service, you must select one of them, using the `snap set` command (or via snapd's REST API). For example, to use the mqtt-export profile you would run:

```
$ sudo snap set edgex-app-service-configurable profile=mqtt-export
```

In addition to instructing the service to read a different configuration file, the profile will also be used to name the service when it registers itself to the system.

**Note** - currently the only two profiles available for Jakarta (2.1) are http-export and mqtt-export.

## Service Environment Configuration Overrides
**Note** - all of the configuration options below must be specified with the prefix: 'env.'

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
