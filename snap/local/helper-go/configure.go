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
	"os"

	"github.com/canonical/edgex-snap-hooks/v3/env"
	"github.com/canonical/edgex-snap-hooks/v3/log"
	"github.com/canonical/edgex-snap-hooks/v3/options"
	"github.com/canonical/edgex-snap-hooks/v3/snapctl"
)

// validateProfile processes the snap 'profile' configure option, ensuring that the directory
// and associated configuration.yaml file in $SNAP_DATA both exist.
func validateProfile() error {
	prof, err := snapctl.Get("profile").Run()
	if err != nil {
		return fmt.Errorf("error getting snap option: %v", err)
	}

	log.Debugf("validateProfile: profile is %s", prof)

	if prof == "" {
		return nil
	}

	path := fmt.Sprintf("%s/config/res/%s/configuration.yaml", env.SnapData, prof)
	log.Debugf("validateProfile: checking if %s exists", path)

	_, err = os.Stat(path)
	if err != nil {
		return fmt.Errorf("profile %s has no configuration.yaml", prof)
	}

	return nil
}

func configure() {
	log.SetComponentName("configure")

	log.Info("Validating profile configuration")
	if err := validateProfile(); err != nil {
		log.Fatalf("Error validating profile: %v", err)
	}

	if err := options.ProcessConfig(app); err != nil {
		log.Fatalf("Error processing config options: %v", err)
	}

	if err := options.ProcessAutostart(app); err != nil {
		log.Fatalf("Error processing autostart options: %v", err)
	}
}
