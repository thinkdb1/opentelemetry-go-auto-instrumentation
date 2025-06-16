// Copyright (c) 2024 Alibaba Group Holding Ltd.
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

package test

import (
	"testing"
)

const AppName = "flags"

func TestFlags(t *testing.T) {
	UseApp(AppName)

	RunGoBuildFallible(t, "go", "build", "-thisisnotvalid")
	ExpectDebugLogContains(t, "Fatal Error")

	RunVersion(t)
	ExpectStdoutContains(t, "version")

	RunGoBuildFallible(t, "go", "build", "notevenaflag")
	ExpectDebugLogContains(t, "Fatal Error")

	RunSet(t, "-verbose")
	RunGoBuild(t, "go", "build", `-ldflags=-X main.Placeholder=replaced`)
	_, stderr := RunApp(t, AppName)
	ExpectContains(t, stderr, "placeholder:replaced")

	RunGoBuild(t, "go")
	RunGoBuild(t)
	RunGoBuild(t, "")
}

func TestFlagEnvOverwrite(t *testing.T) {
	UseApp(AppName)

	RunSet(t, "-verbose=false")
	RunGoBuildWithEnv(t, []string{"OTELTOOL_VERBOSE=true"},
		"go", "build")
	ExpectDebugLogContains(t, "Available")

	RunSet(t, "-verbose=true")
	RunGoBuildWithEnv(t, []string{"OTELTOOL_VERBOSE=false"},
		"go", "build")
	ExpectDebugLogNotContains(t, "Available")
}
