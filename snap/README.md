# EdgeX App Service Configurable Snap
[![edgex-app-service-configurable](https://snapcraft.io/edgex-app-service-configurable/badge.svg)][edgex-app-service-configurable]

This directory contains the snap packaging of the EdgeX App Service Configurable application service.

The snap is built automatically and published on the Snap Store as [edgex-app-service-configurable].

For usage instructions, please refer to [docs](https://docs.edgexfoundry.org/2.2/getting-started/Ch-GettingStartedSnapUsers/#app-service-configurable).

## Build from source
Execute the following command from the top-level directory of this repo:
```
snapcraft
```

This will create binary snap package with a `.snap` extension which can be installed locally by setting the `--dangerous` flag:
```bash
sudo snap install --dangerous <snap>.snap
```

Please refer to [snapcraft documentation](https://snapcraft.io/docs) for more details.

### Obtain a secret store token
When this and the [edgexfoundry] snap are installed from the store, the Secret Store token is provided automatically using the auto-connected `edgex-secretstore-token` [content interface](https://snapcraft.io/docs/content-interface).

If the snap is build locally, the auto connection will not happen. This can be verified by running the `snap connections edgex-app-service-configurable` command.

To manually connect and obtain the token:
```bash
sudo snap connect edgexfoundry:edgex-secretstore-token edgex-app-service-configurable:edgex-secretstore-token
```

To better understand the snap connections, read the [interface management](https://snapcraft.io/docs/interface-management) documentation.

[edgexfounndry]: https://snapcraft.io/edgexfoundry
[edgex-app-service-configurable]: https://snapcraft.io/edgex-app-service-configurable