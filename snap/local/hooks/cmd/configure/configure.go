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
	"errors"
	"fmt"
	"os"
	"strings"

	hooks "github.com/canonical/edgex-snap-hooks/v2"
)

var cli *hooks.CtlCli = hooks.NewSnapCtl()

// validateProfile processes the snap 'profile' configure option, ensuring that the directory
// and associated configuration.toml file in $SNAP_DATA both exist.
//
func validateProfile(prof string) error {
	hooks.Debug(fmt.Sprintf("edgex-asc:configure:validateProfile: profile is %s", prof))

	if prof == "" {
		return nil
	}

	path := fmt.Sprintf("%s/config/res/%s/configuration.toml", hooks.SnapData, prof)
	hooks.Debug(fmt.Sprintf("edgex-asc:configure:validateProfile: checking if %s exists", path))

	_, err := os.Stat(path)
	if err != nil {
		return errors.New(fmt.Sprintf("profile %s has no configuration.toml", prof))
	}

	return nil
}

func main() {
	var debug = false
	var enable = true
	var err error
	var envJSON, prof string

	status, err := cli.Config("debug")
	if err != nil {
		fmt.Println(fmt.Sprintf("edgex-asc:configure: can't read value of 'debug': %v", err))
		os.Exit(1)
	}
	if status == "true" {
		debug = true
	}

	if err = hooks.Init(debug, "edgex-app-service-configurable"); err != nil {
		fmt.Println(fmt.Sprintf("edgex-asc:configure: initialization failure: %v", err))
		os.Exit(1)

	}

	prof, err = cli.Config(hooks.ProfileConfig)
	if err != nil {
		hooks.Error(fmt.Sprintf("Error reading config 'profile': %v", err))
		os.Exit(1)
	}

	validateProfile(prof)
	if err != nil {
		hooks.Error(fmt.Sprintf("Error validating profile: %v", err))
		os.Exit(1)
	}

	envJSON, err = cli.Config(hooks.EnvConfig)
	if err != nil {
		hooks.Error(fmt.Sprintf("Reading config 'env' failed: %v", err))
		os.Exit(1)
	}

	err = hooks.HandleEdgeXConfig("app-service-configurable", envJSON, nil)
	if err != nil {
		hooks.Error(fmt.Sprintf("HandleEdgeXConfig failed: %v", err))
		os.Exit(1)
	}

	// If autostart is not explicitly set, default to "no"
	// as only example service configuration and profiles
	// are provided by default.
	autostart, err := cli.Config(hooks.AutostartConfig)
	if err != nil {
		hooks.Error(fmt.Sprintf("Reading config 'autostart' failed: %v", err))
		os.Exit(1)
	}
	if autostart == "" {
		hooks.Debug("edgex-asc: autostart is NOT set, initializing to 'no'")
		autostart = "no"
	}

	autostart = strings.ToLower(autostart)
	if autostart == "true" || autostart == "yes" {
		if prof == "" {
			hooks.Warn(fmt.Sprintf("autostart is %s, but no profile set", autostart))
			enable = false
		}
	} else if autostart == "false" || autostart == "no" {
		enable = false
	} else {
		hooks.Error(fmt.Sprintf("Invalid value for 'autostart' : %s", autostart))
		os.Exit(1)
	}

	// service is stopped/disabled by default in the install hook
	// only enable if profile exists and autstart is set
	if enable {
		err = cli.Start("app-service-configurable", true)
		if err != nil {
			hooks.Error(fmt.Sprintf("Can't start service - %v", err))
			os.Exit(1)
		}
	}
}
