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

#build stage
FROM golang:1.12-alpine AS builder
LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2019: Intel'
WORKDIR /app
RUN apk update && apk add --no-cache make git gcc libc-dev libsodium-dev zeromq-dev
COPY . .
RUN make build

#final stage
FROM alpine:latest
LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2019: Intel'
LABEL Name=app-service-configurable Version=0.0.1

RUN apk --no-cache add ca-certificates zeromq
COPY --from=builder /app/res/ /res/
COPY --from=builder /app/app-service-configurable /app-service-configurable
EXPOSE 48095

ENTRYPOINT ["/app-service-configurable","--registry","--profile=OVERWRITEME","--confdir=/res"]

