//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
)

const (
	serviceKey = "app-" + interfaces.ProfileSuffixPlaceholder
)

func main() {
	service, ok := pkg.NewAppService(serviceKey)
	if !ok {
		os.Exit(-1)
	}

	lc := service.LoggingClient()

	lc.Info("Loading Configurable Pipeline...")

	configuredPipelines, err := service.LoadConfigurableFunctionPipelines()
	if err != nil {
		lc.Errorf("failed to load configured function pipelines: %s", err.Error())
		os.Exit(-1)
	}

	lc.Debugf("Found %d configured pipeline(s)", len(configuredPipelines))

	for _, pipeline := range configuredPipelines {
		switch pipeline.Id {
		case interfaces.DefaultPipelineId:
			if err = service.SetDefaultFunctionsPipeline(pipeline.Transforms...); err != nil {
				lc.Errorf("Unable to Set Default Functions Pipeline: %s", pipeline.Id, err.Error())
				os.Exit(-1)
			}
		default:
			if err = service.AddFunctionsPipelineForTopics(pipeline.Id, pipeline.Topics, pipeline.Transforms...); err != nil {
				lc.Errorf("Unable to Add Functions Pipeline for pipeline id '%s': %s", pipeline.Id, err.Error())
				os.Exit(-1)
			}
		}
	}

	if err = service.Run(); err != nil {
		lc.Errorf("Run returned error: %s", err.Error())
		os.Exit(-1)
	}

	os.Exit(0)
}
