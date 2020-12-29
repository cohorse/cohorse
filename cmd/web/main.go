// Copyright 2020 The Cohorse Author
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/cohorse/cohorse/pkg/sysinfo"
	_ "github.com/cohorse/cohorse/pkg/web/vuex"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := make(chan struct{}, 0)

	log.WithFields(log.Fields{
		"version": sysinfo.Version,
		"build":   sysinfo.Build,
		"commit":  sysinfo.Commit,
	}).Info("Starting Cohorse WASM binary.")

	<-c
}
