//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogFilesSpec) DeepCopyInto(out *MailhogFilesSpec) {
	*out = *in
	if in.SmtpUpstreams != nil {
		in, out := &in.SmtpUpstreams, &out.SmtpUpstreams
		*out = make([]MailhogUpstreamSpec, len(*in))
		copy(*out, *in)
	}
	if in.WebUsers != nil {
		in, out := &in.WebUsers, &out.WebUsers
		*out = make([]MailhogWebUserSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogFilesSpec.
func (in *MailhogFilesSpec) DeepCopy() *MailhogFilesSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogFilesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogInstance) DeepCopyInto(out *MailhogInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogInstance.
func (in *MailhogInstance) DeepCopy() *MailhogInstance {
	if in == nil {
		return nil
	}
	out := new(MailhogInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MailhogInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogInstanceList) DeepCopyInto(out *MailhogInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MailhogInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogInstanceList.
func (in *MailhogInstanceList) DeepCopy() *MailhogInstanceList {
	if in == nil {
		return nil
	}
	out := new(MailhogInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MailhogInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogInstanceSettingsSpec) DeepCopyInto(out *MailhogInstanceSettingsSpec) {
	*out = *in
	out.StorageMongoDb = in.StorageMongoDb
	out.StorageMaildir = in.StorageMaildir
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(MailhogFilesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	out.Jim = in.Jim
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogInstanceSettingsSpec.
func (in *MailhogInstanceSettingsSpec) DeepCopy() *MailhogInstanceSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogInstanceSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogInstanceSpec) DeepCopyInto(out *MailhogInstanceSpec) {
	*out = *in
	in.Settings.DeepCopyInto(&out.Settings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogInstanceSpec.
func (in *MailhogInstanceSpec) DeepCopy() *MailhogInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogInstanceStatus) DeepCopyInto(out *MailhogInstanceStatus) {
	*out = *in
	in.Pods.DeepCopyInto(&out.Pods)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogInstanceStatus.
func (in *MailhogInstanceStatus) DeepCopy() *MailhogInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(MailhogInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogJimSpec) DeepCopyInto(out *MailhogJimSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogJimSpec.
func (in *MailhogJimSpec) DeepCopy() *MailhogJimSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogJimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogStorageMaildirSpec) DeepCopyInto(out *MailhogStorageMaildirSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogStorageMaildirSpec.
func (in *MailhogStorageMaildirSpec) DeepCopy() *MailhogStorageMaildirSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogStorageMaildirSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogStorageMongoDbSpec) DeepCopyInto(out *MailhogStorageMongoDbSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogStorageMongoDbSpec.
func (in *MailhogStorageMongoDbSpec) DeepCopy() *MailhogStorageMongoDbSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogStorageMongoDbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogUpstreamSpec) DeepCopyInto(out *MailhogUpstreamSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogUpstreamSpec.
func (in *MailhogUpstreamSpec) DeepCopy() *MailhogUpstreamSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogUpstreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailhogWebUserSpec) DeepCopyInto(out *MailhogWebUserSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailhogWebUserSpec.
func (in *MailhogWebUserSpec) DeepCopy() *MailhogWebUserSpec {
	if in == nil {
		return nil
	}
	out := new(MailhogWebUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfig) DeepCopyInto(out *OperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfig.
func (in *OperatorConfig) DeepCopy() *OperatorConfig {
	if in == nil {
		return nil
	}
	out := new(OperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Restarting != nil {
		in, out := &in.Restarting, &out.Restarting
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Other != nil {
		in, out := &in.Other, &out.Other
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}
