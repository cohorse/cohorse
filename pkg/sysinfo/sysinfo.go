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

package sysinfo

var (
	// Version is injected on every production build with a SemVer-compliant
	// version string.
	Version = "unknown"
	// Build is injected on every production build with the build number of the
	// current build.
	Build = "dev"
	// Commit is injected on every production build with the latest commit hash.
	Commit = "unknown"
)
