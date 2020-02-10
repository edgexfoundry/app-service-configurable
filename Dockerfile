#
# Copyright (c) 2019 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# This file will work as is for local development. No need to use Dockerfile.build

#build stage
ARG BASE=golang:1.13-alpine
FROM ${BASE} AS builder

ARG ALPINE_PKG_BASE="make git gcc libc-dev libsodium-dev zeromq-dev"
ARG ALPINE_PKG_EXTRA=""

LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2019: Intel'
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

ARG MAKE="make build"
RUN $MAKE

#final stage
FROM alpine:latest
LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2019: Intel'
LABEL Name=app-service-configurable Version=${VERSION}

RUN apk --no-cache add ca-certificates zeromq
COPY --from=builder /app/res/ /res/
COPY --from=builder /app/app-service-configurable /app-service-configurable
EXPOSE 48095

# Default configuation has been moved to the new "default" profile.
# Must always specify the profile using
# environment:
#   - edgex_profile: <profile>
# or use
# command: "-profile=<profile>"
# If not you will recive error:
# SDK initialization failed: Could not load configuration file (./res/configuration.toml)...

ENTRYPOINT ["/app-service-configurable","--registry","--confdir=/res"]

