/*
Copyright 2021 The logr Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/go-logr/logr"

	"github.com/golang/glog"

	"github.com/go-logr/glogr"
)

func init() {
	_ = flag.Set("v", "1")
	_ = flag.Set("logtostderr", "true")
	os.Stderr, _ = os.Open("/dev/null")
}

//go:noinline
func doInfoOneArg(b *testing.B, log logr.Logger) {
	for i := 0; i < b.N; i++ {
		log.Info("this is", "a", "string")
	}
}

//go:noinline
func doInfoOneArg_glog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.Info("this is", "a", "string")
	}
}

//go:noinline
func doInfoSeveralArgs(b *testing.B, log logr.Logger) {
	for i := 0; i < b.N; i++ {
		log.Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doInfoSeveralArgs_glog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doV0Info(b *testing.B, log logr.Logger) {
	for i := 0; i < b.N; i++ {
		log.V(0).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doV0Info_glog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.V(0).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}


//go:noinline
func doV9Info(b *testing.B, log logr.Logger) {
	for i := 0; i < b.N; i++ {
		log.V(9).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doV9Info_glog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		glog.V(glog.Level(9)).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doError(b *testing.B, log logr.Logger) {
	err := fmt.Errorf("error message")
	for i := 0; i < b.N; i++ {
		log.Error(err, "multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doError_glog(b *testing.B) {
	err := fmt.Errorf("error message")
	for i := 0; i < b.N; i++ {
		glog.Error(err, "multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doWithValues(b *testing.B, log logr.Logger) {
	for i := 0; i < b.N; i++ {
		l := log.WithValues("k1", "v1", "k2", "v2")
		_ = l
	}
}

//go:noinline
func doWithName(b *testing.B, log logr.Logger) {
	for i := 0; i < b.N; i++ {
		l := log.WithName("name")
		_ = l
	}
}

func BenchmarkDiscardInfoOneArg(b *testing.B) {
	var log logr.Logger = glogr.New()
	doInfoOneArg(b, log)
}

func BenchmarkDiscardInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = glogr.New()
	doInfoSeveralArgs(b, log)
}

func BenchmarkDiscardV0Info(b *testing.B) {
	var log logr.Logger = glogr.New()
	doV0Info(b, log)
}

func BenchmarkDiscardV9Info(b *testing.B) {
	var log logr.Logger = glogr.New()
	doV9Info(b, log)
}

func BenchmarkDiscardError(b *testing.B) {
	var log logr.Logger = glogr.New()
	doError(b, log)
}

func BenchmarkDiscardWithValues(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithValues(b, log)
}

func BenchmarkDiscardWithName(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithName(b, log)
}

func BenchmarkFuncrInfoOneArg(b *testing.B) {
	var log logr.Logger = glogr.New()
	doInfoOneArg(b, log)
}

func BenchmarkFuncrInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = glogr.New()
	doInfoSeveralArgs(b, log)
}

func BenchmarkFuncrV0Info(b *testing.B) {
	var log logr.Logger = glogr.New()
	doV0Info(b, log)
}

func BenchmarkFuncrV9Info(b *testing.B) {
	var log logr.Logger = glogr.New()
	doV9Info(b, log)
}

func BenchmarkFuncrError(b *testing.B) {
	var log logr.Logger = glogr.New()
	doError(b, log)
}

func BenchmarkFuncrWithValues(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithValues(b, log)
}

func BenchmarkFuncrWithName(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithName(b, log)
}


func BenchmarkDiscardInfoOneArg_glog(b *testing.B) {
	doInfoOneArg_glog(b)
}

func BenchmarkDiscardInfoSeveralArgs_glog(b *testing.B) {
	doInfoSeveralArgs_glog(b)
}

func BenchmarkDiscardV0Info_glog(b *testing.B) {
	doV0Info_glog(b)
}

func BenchmarkDiscardV9Info_glog(b *testing.B) {
	doV9Info_glog(b)
}

func BenchmarkFuncrInfoOneArg_glog(b *testing.B) {
	doInfoOneArg_glog(b)
}

func BenchmarkFuncrInfoSeveralArgs_glog(b *testing.B) {
	doInfoSeveralArgs_glog(b)
}

func BenchmarkFuncrV0Info_glog(b *testing.B) {
	doV0Info_glog(b)
}

func BenchmarkFuncrV9Info_glog(b *testing.B) {
	doV9Info_glog(b)
}

func BenchmarkFuncrError_glog(b *testing.B) {
	doError_glog(b)
}
