/*
Copyright 2020 The Kubernetes Authors.

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

package log

import (
	"fmt"
	"time"

	"github.com/onsi/ginkgo"

	"sigs.k8s.io/cluster-api-provider-nested/virtualcluster/test/e2e/framework/ginkgowrapper"
)

func nowStamp() string {
	return time.Now().Format(time.StampMilli)
}

func logf(level string, format string, args ...interface{}) {
	fmt.Fprintf(ginkgo.GinkgoWriter, nowStamp()+": "+level+": "+format+"\n", args...)
}

// Logf logs the info.
func Logf(format string, args ...interface{}) {
	logf("INFO", format, args...)
}

// Failf logs the fail info.
func Failf(format string, args ...interface{}) {
	FailfWithOffsetf(1, format, args...)
}

// FailfWithOffsetf calls "Fail" and logs the error at "offset" levels above its caller
// (for example, for call chain f -> g -> FailfWithOffsetf(1, ...) error would be logged for "f").
func FailfWithOffsetf(offset int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logf("FAIL", msg)
	ginkgowrapper.Fail(nowStamp()+": "+msg, 1+offset)
}

// Fail is a replacement for ginkgo.Fail which logs the problem as it occurs
// and then calls ginkgowrapper.Fail.
func Fail(msg string, callerSkip ...int) {
	skip := 1
	if len(callerSkip) > 0 {
		skip += callerSkip[0]
	}
	logf("FAIL", msg)
	ginkgowrapper.Fail(nowStamp()+": "+msg, skip)
}
