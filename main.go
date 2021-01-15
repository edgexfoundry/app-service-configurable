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
	"fmt"
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/v2/appsdk"
)

const (
	serviceKey = "AppService-" + appsdk.ProfileSuffixPlaceholder
)

func main() {
	fmt.Println("Starting Configurable Application Service...")
	edgexSdk := &appsdk.AppFunctionsSDK{ServiceKey: serviceKey}
	if err := edgexSdk.Initialize(); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK initialization failed: %v\n", err))
		os.Exit(-1)
	}

	edgexSdk.LoggingClient.Info("Loading Configurable Pipeline...")

	transforms, err := edgexSdk.LoadConfigurablePipeline()
	if err != nil {
		edgexSdk.LoggingClient.Error("failed to create function pipeline from configuration: " + err.Error())
		os.Exit(-1)
	}

	edgexSdk.SetFunctionsPipeline(transforms...)

	if err := edgexSdk.MakeItRun(); err != nil {
		edgexSdk.LoggingClient.Error("MakeItRun returned error: ", err.Error())
		os.Exit(-1)
	}

	os.Exit(0)
}
