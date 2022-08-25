# App Service Configurable
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/app-service-configurable/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/app-service-configurable/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/app-service-configurable/branch/main/graph/badge.svg?token=wMrc2PZbwj)](https://codecov.io/gh/edgexfoundry/app-service-configurable) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/app-service-configurable)](https://goreportcard.com/report/github.com/edgexfoundry/app-service-configurable) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/app-service-configurable?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/app-service-configurable/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/app-service-configurable?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/app-service-configurable)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/app-service-configurable) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/app-service-configurable)](https://github.com/edgexfoundry/app-service-configurable/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/app-service-configurable)](https://github.com/edgexfoundry/app-service-configurable/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/app-service-configurable-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/app-service-configurable)](https://github.com/edgexfoundry/app-service-configurable/commits)


App-Service-Configurable is provided as an easy way to get started with processing data flowing through EdgeX.

For latest documentation please visit https://docs.edgexfoundry.org and navigate to **Microservices** (On top menu) then **Application Service -> App Service Configurable** (On left side menu)

For a list of examples please visit https://docs.edgexfoundry.org and navigate to **Example and Tutorials** (On top menu) then **Application Services** tab and follow the link in the description.

## Build Prerequisites

Please see the [edgex-go README](https://github.com/edgexfoundry/edgex-go/blob/main/README.md).

## Building
```
make build - builds local binary 
make -e ADD_BUILD_TAGS=include_nats_messaging build - Builds local binary with NATS MessageBus support

make docker - build local docker image
make -e ADD_BUILD_TAGS=include_nats_messaging docker - Builds local docker image with NATS MessageBus support
```
## Packaging

This component is packaged as docker image and snap.

For docker, please refer to the [Dockerfile](Dockerfile) and [Docker Compose Builder](https://github.com/edgexfoundry/edgex-compose/tree/main/compose-builder) scripts.

For the snap, refer to the [snap](snap) directory.

## Versioning

Please refer to the EdgeX Foundry [versioning policy](https://wiki.edgexfoundry.org/pages/viewpage.action?pageId=21823969) for information on how EdgeX services are released and how EdgeX services are compatible with one another.  Specifically, device services (and the associated SDK), application services (and the associated app functions SDK), and client tools (like the EdgeX CLI and UI) can have independent minor releases, but these services must be compatible with the latest major release of EdgeX.

## Long Term Support

Please refer to the EdgeX Foundry [LTS policy](https://wiki.edgexfoundry.org/pages/viewpage.action?pageId=69173332#LongTermSupport(v2)-LTS) for information on support of EdgeX releases. The EdgeX community does not offer support on any non-LTS release outside of the latest release.

