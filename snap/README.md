# EdgeX App Service Configurable Snap
[![edgex-app-service-configurable](https://snapcraft.io/edgex-app-service-configurable/badge.svg)](https://snapcraft.io/edgex-app-service-configurable)

This folder contains the snap packaging of the EdgeX App Service Configurable application service.

For usage instructions, please refer to [docs](https://docs.edgexfoundry.org/2.2/getting-started/Ch-GettingStartedSnapUsers/#app-service-configurable).

## Build
To build the snap from source, execute the following command from the top-level directory of this repo:
```
snapcraft
```

Please refer to [snapcraft documentation](https://snapcraft.io/docs) for details.

### Obtain a secret store token
When this and the edgexfoundry snap are installed from the store, the Secret Store token is provided automatically using the auto-connected `edgex-secretstore-token` [content interface](https://snapcraft.io/docs/content-interface).

If the snap is build locally, the auto connection will not happen. This can be verified by running the `snap connections <snap>` command. To manually connect and obtain the token:

```bash
sudo snap connect edgexfoundry:edgex-secretstore-token edgex-app-service-configurable:edgex-secretstore-token 
```
