//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskByStatement) DeepCopyInto(out *TaskByStatement) {
	*out = *in
	in.TaskRunValueRef.DeepCopyInto(&out.TaskRunValueRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskByStatement.
func (in *TaskByStatement) DeepCopy() *TaskByStatement {
	if in == nil {
		return nil
	}
	out := new(TaskByStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMetric) DeepCopyInto(out *TaskMetric) {
	*out = *in
	if in.By != nil {
		in, out := &in.By, &out.By
		*out = make([]TaskByStatement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(TaskMetricHistogramDuration)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(TaskMetricGaugeMatch)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMetric.
func (in *TaskMetric) DeepCopy() *TaskMetric {
	if in == nil {
		return nil
	}
	out := new(TaskMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMetricGaugeMatch) DeepCopyInto(out *TaskMetricGaugeMatch) {
	*out = *in
	in.Key.DeepCopyInto(&out.Key)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMetricGaugeMatch.
func (in *TaskMetricGaugeMatch) DeepCopy() *TaskMetricGaugeMatch {
	if in == nil {
		return nil
	}
	out := new(TaskMetricGaugeMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMetricHistogramDuration) DeepCopyInto(out *TaskMetricHistogramDuration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMetricHistogramDuration.
func (in *TaskMetricHistogramDuration) DeepCopy() *TaskMetricHistogramDuration {
	if in == nil {
		return nil
	}
	out := new(TaskMetricHistogramDuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMonitor) DeepCopyInto(out *TaskMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMonitor.
func (in *TaskMonitor) DeepCopy() *TaskMonitor {
	if in == nil {
		return nil
	}
	out := new(TaskMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMonitorList) DeepCopyInto(out *TaskMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TaskMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMonitorList.
func (in *TaskMonitorList) DeepCopy() *TaskMonitorList {
	if in == nil {
		return nil
	}
	out := new(TaskMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMonitorSpec) DeepCopyInto(out *TaskMonitorSpec) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]TaskMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMonitorSpec.
func (in *TaskMonitorSpec) DeepCopy() *TaskMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(TaskMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskMonitorStatus) DeepCopyInto(out *TaskMonitorStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskMonitorStatus.
func (in *TaskMonitorStatus) DeepCopy() *TaskMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(TaskMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRunValueRef) DeepCopyInto(out *TaskRunValueRef) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.Param != nil {
		in, out := &in.Param, &out.Param
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRunValueRef.
func (in *TaskRunValueRef) DeepCopy() *TaskRunValueRef {
	if in == nil {
		return nil
	}
	out := new(TaskRunValueRef)
	in.DeepCopyInto(out)
	return out
}
