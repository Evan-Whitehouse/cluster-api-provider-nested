/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	"k8s.io/client-go/kubernetes/scheme"

	vcapis "sigs.k8s.io/cluster-api-provider-nested/virtualcluster/pkg/apis"
	"sigs.k8s.io/cluster-api-provider-nested/virtualcluster/test/e2e/framework"

	// test sources
	_ "sigs.k8s.io/cluster-api-provider-nested/virtualcluster/test/e2e/multi-tenancy"
)

func TestMain(m *testing.M) {
	// Register test flags, then parse flags.
	framework.HandleFlags()

	framework.AfterReadingAllFlags(&framework.TestContext)

	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	flag.Parse()
	err := vcapis.AddToScheme(scheme.Scheme)
	if err != nil {
		t.Fatal(err)
	}

	RunE2ETests(t)
}
