// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	hooks "github.com/canonical/edgex-snap-hooks/v2"
)

var cli *hooks.CtlCli = hooks.NewSnapCtl()

// installProfiles copies the profile configuration.toml files from $SNAP to $SNAP_DATA.
func installProfiles() error {
	dataConfP := fmt.Sprintf("%s/config/res", hooks.SnapData)
	snapConfP := fmt.Sprintf("%s/config/res", hooks.Snap)

	configFiles, err := filepath.Glob(filepath.Join(snapConfP, "*", "configuration.toml"))
	if err != nil {
		panic(fmt.Sprintf("internal error: bad glob pattern: %v", err))
	}

	for _, snapConfFile := range configFiles {
		// build the destination SNAP_DATA file by getting the directory name that the glob matched
		dirMatch := filepath.Base(filepath.Dir(snapConfFile))
		if dirMatch == "sample" {
			// TODO: what about sample config dirs ?
			continue
		}

		dataDestFile := filepath.Join(dataConfP, dirMatch, "configuration.toml")
		b, err := ioutil.ReadFile(snapConfFile)
		if err != nil {
			return err
		}

		err = os.MkdirAll(filepath.Dir(dataDestFile), 0755)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(dataDestFile, b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var err error

	if err = hooks.Init(false, "edgex-app-service-configurable"); err != nil {
		fmt.Println(fmt.Sprintf("edgex-asc:install: initialization failure: %v", err))
		os.Exit(1)

	}

	err = installProfiles()
	if err != nil {
		hooks.Error(fmt.Sprintf("edgex-asc:install: %v", err))
		os.Exit(1)
	}
}
