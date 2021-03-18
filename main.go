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

	"github.com/edgexfoundry/app-functions-sdk-go/v2/pkg"
	"github.com/edgexfoundry/app-functions-sdk-go/v2/pkg/interfaces"
)

const (
	serviceKey = "AppService-" + interfaces.ProfileSuffixPlaceholder
)

func main() {
	service, ok := pkg.NewAppService(serviceKey)
	if !ok {
		os.Exit(-1)
	}

	lc := service.LoggingClient()

	lc.Info("Loading Configurable Pipeline...")

	transforms, err := service.LoadConfigurablePipeline()
	if err != nil {
		lc.Errorf("failed to create function pipeline from configuration: %s", err.Error())
		os.Exit(-1)
	}

	if err = service.SetFunctionsPipeline(transforms...); err != nil {
		lc.Errorf("Unable to Set the Functions Pipeline: %s", err.Error())
		os.Exit(-1)
	}

	if err = service.MakeItRun(); err != nil {
		lc.Errorf("MakeItRun returned error: %s", err.Error())
		os.Exit(-1)
	}

	os.Exit(0)
}
