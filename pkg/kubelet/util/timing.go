/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package util

import (
	"time"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
)

// HackStartTimeByAnnotation returns the current time if the proper annotation
// is on the pod. Otherwise it reurns a zero time
func HackStartTimeByAnnotation(pod *api.Pod) time.Time {
	if _, ok := pod.Annotations["kuberenetes.io/hack.13052"]; ok {
		return time.Now()
	}
	return time.Time{}
}

// HackEndTimeByAnnotation logs the time difference if startTime is non-zero
func HackEndTimeByAnnotation(startTime time.Time, pod *api.Pod, name string) {
	if !startTime.IsZero() {
		glog.Warningf("13052 HACK: %q,%q,%q", pod.UID, name, time.Now().Sub(startTime))
	}
}
