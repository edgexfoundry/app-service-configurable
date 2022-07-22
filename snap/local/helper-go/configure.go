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
	hooks "github.com/canonical/edgex-snap-hooks/v2"
	"github.com/canonical/edgex-snap-hooks/v2/log"
	"github.com/canonical/edgex-snap-hooks/v2/options"
	"os"
)

// configure is called by the main function
func configure() {

	const service = "app-service-configurable"

	log.SetComponentName("configure")

	prof, err := hooks.NewSnapCtl().Config(hooks.ProfileConfig)
	if err != nil {
		log.Fatalf("Error reading config 'profile': %v", err)
	}

	validateProfile(prof)
	if err != nil {
		log.Fatalf("Error validating profile: %v", err)
	}

	log.Info("Processing legacy env options")
	envJSON, err := hooks.NewSnapCtl().Config(hooks.EnvConfig)
	if err != nil {
		log.Fatalf("Reading config 'env' failed: %v", err)
	}
	if envJSON != "" {
		log.Debugf("envJSON: %s", envJSON)
		err = hooks.HandleEdgeXConfig(service, envJSON, nil)
		if err != nil {
			log.Fatalf("HandleEdgeXConfig failed: %v", err)
		}
	}

	log.Info("Processing config options")
	err = options.ProcessConfig(service)
	if err != nil {
		log.Fatalf("could not process config options: %v", err)
	}

	log.Info("Processing autostart options")
	err = options.ProcessAutostart(service)
	if err != nil {
		log.Fatalf("could not process autostart options: %v", err)
	}

	//// service is stopped/disabled by default in the install hook
	//// only enable if profile exists and autostart is set
	//s, err := snapctl.Services(service).Run()
	//if err != nil {
	//	log.Fatalf("Can't get service - %v", err)
	//}
	//value, _ := s[service]
	//
	//if value.Enabled {
	//	err = snapctl.Start(service).Run()
	//	if err != nil {
	//		log.Fatalf("Can't start service - %v", err)
	//	}
	//}

	// TODO: debugger
	err = hooks.NewSnapCtl().Start("app-service-configurable", true)
	if err != nil {
		hooks.Error(fmt.Sprintf("Can't start service - %v", err))
		os.Exit(1)
	}
}

// validateProfile processes the snap 'profile' configure option, ensuring that the directory
// and associated configuration.toml file in $SNAP_DATA both exist.
//
func validateProfile(prof string) error {
	log.Debugf("validateProfile: profile is %s", prof)

	if prof == "" {
		return nil
	}

	path := fmt.Sprintf("%s/config/res/%s/configuration.toml", hooks.SnapData, prof)
	log.Debugf("validateProfile: checking if %s exists", path)

	_, err := os.Stat(path)
	if err != nil {
		return errors.New(fmt.Sprintf("profile %s has no configuration.toml", prof))
	}

	return nil
}
