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

func TestBuildProject(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)
	RunGoBuild(t, "go", "build", "-o", "default", "cmd/foo.go")
	RunGoBuild(t, "go", "build", "-o", "./cmd", "./cmd")
	RunGoBuild(t, "go", "build", "cmd/foo.go", "cmd/bar.go")

}

func TestBuildProject2(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)

	RunGoBuild(t, "go", "build", ".")
	RunGoBuild(t, "go", "build", "./...")
}

func TestBuildProject3(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)

	RunGoBuildFallible(t, "go", "build", "m2") // not used in go.work
}

func TestBuildProject4(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)

	RunSet(t, "-disabledefault=false", "-rule=../../tool/data/default.json")
	RunGoBuildFallible(t, "go", "build", "m1") // duplicated default rules
	RunSet(t, "-rule=../../tool/data/default")
	RunGoBuildFallible(t, "go", "build", "m1")
	RunSet(t, "-disabledefault=true", "-rule=../../tool/data/default.json,../../tool/data/test_fmt.json")
	RunGoBuild(t, "go", "build", "m1")
}

func TestBuildProject5(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)

	RunSet(t, "-disabledefault=false", "-verbose", "-rule=../../tool/data/test_fmt.json")
	RunGoBuild(t, "go", "build", "m1")
	// both test_fmt.json and default.json rules should be available
	// because we always append new -rule to the default.json by default
	ExpectDebugLogContains(t, "fmt")
	ExpectDebugLogContains(t, "github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/rules/http")
}

func TestBuildProject6(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)

	RunSet(t, "-disabledefault=true", "-rule=../../tool/data/test_fmt.json,../../tool/data/test_runtime.json", "-verbose")
	RunGoBuild(t, "go", "build", "m1")
	// only test_fmt.json should be available because -disabledefault is set
	ExpectDebugLogContains(t, "fmt")
	ExpectDebugLogNotContains(t, "github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/rules/http")
}

func TestGoInstall(t *testing.T) {
	const AppName = "build"
	UseApp(AppName)
	RunGoBuild(t, "go", "install", "./cmd/...")
}
