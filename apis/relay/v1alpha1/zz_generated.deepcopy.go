// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnection) DeepCopyInto(out *HybridConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnection.
func (in *HybridConnection) DeepCopy() *HybridConnection {
	if in == nil {
		return nil
	}
	out := new(HybridConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRule) DeepCopyInto(out *HybridConnectionAuthorizationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRule.
func (in *HybridConnectionAuthorizationRule) DeepCopy() *HybridConnectionAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnectionAuthorizationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleList) DeepCopyInto(out *HybridConnectionAuthorizationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HybridConnectionAuthorizationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleList.
func (in *HybridConnectionAuthorizationRuleList) DeepCopy() *HybridConnectionAuthorizationRuleList {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnectionAuthorizationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleObservation) DeepCopyInto(out *HybridConnectionAuthorizationRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleObservation.
func (in *HybridConnectionAuthorizationRuleObservation) DeepCopy() *HybridConnectionAuthorizationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleParameters) DeepCopyInto(out *HybridConnectionAuthorizationRuleParameters) {
	*out = *in
	if in.HybridConnectionName != nil {
		in, out := &in.HybridConnectionName, &out.HybridConnectionName
		*out = new(string)
		**out = **in
	}
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleParameters.
func (in *HybridConnectionAuthorizationRuleParameters) DeepCopy() *HybridConnectionAuthorizationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleSpec) DeepCopyInto(out *HybridConnectionAuthorizationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleSpec.
func (in *HybridConnectionAuthorizationRuleSpec) DeepCopy() *HybridConnectionAuthorizationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleStatus) DeepCopyInto(out *HybridConnectionAuthorizationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleStatus.
func (in *HybridConnectionAuthorizationRuleStatus) DeepCopy() *HybridConnectionAuthorizationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionList) DeepCopyInto(out *HybridConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HybridConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionList.
func (in *HybridConnectionList) DeepCopy() *HybridConnectionList {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionObservation) DeepCopyInto(out *HybridConnectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionObservation.
func (in *HybridConnectionObservation) DeepCopy() *HybridConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionParameters) DeepCopyInto(out *HybridConnectionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceName != nil {
		in, out := &in.RelayNamespaceName, &out.RelayNamespaceName
		*out = new(string)
		**out = **in
	}
	if in.RequiresClientAuthorization != nil {
		in, out := &in.RequiresClientAuthorization, &out.RequiresClientAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.UserMetadata != nil {
		in, out := &in.UserMetadata, &out.UserMetadata
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionParameters.
func (in *HybridConnectionParameters) DeepCopy() *HybridConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionSpec) DeepCopyInto(out *HybridConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionSpec.
func (in *HybridConnectionSpec) DeepCopy() *HybridConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionStatus) DeepCopyInto(out *HybridConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionStatus.
func (in *HybridConnectionStatus) DeepCopy() *HybridConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Namespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRule) DeepCopyInto(out *NamespaceAuthorizationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRule.
func (in *NamespaceAuthorizationRule) DeepCopy() *NamespaceAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceAuthorizationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleList) DeepCopyInto(out *NamespaceAuthorizationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceAuthorizationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleList.
func (in *NamespaceAuthorizationRuleList) DeepCopy() *NamespaceAuthorizationRuleList {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceAuthorizationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleObservation) DeepCopyInto(out *NamespaceAuthorizationRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleObservation.
func (in *NamespaceAuthorizationRuleObservation) DeepCopy() *NamespaceAuthorizationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleParameters) DeepCopyInto(out *NamespaceAuthorizationRuleParameters) {
	*out = *in
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleParameters.
func (in *NamespaceAuthorizationRuleParameters) DeepCopy() *NamespaceAuthorizationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleSpec) DeepCopyInto(out *NamespaceAuthorizationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleSpec.
func (in *NamespaceAuthorizationRuleSpec) DeepCopy() *NamespaceAuthorizationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleStatus) DeepCopyInto(out *NamespaceAuthorizationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleStatus.
func (in *NamespaceAuthorizationRuleStatus) DeepCopy() *NamespaceAuthorizationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceList) DeepCopyInto(out *NamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Namespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceList.
func (in *NamespaceList) DeepCopy() *NamespaceList {
	if in == nil {
		return nil
	}
	out := new(NamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceObservation) DeepCopyInto(out *NamespaceObservation) {
	*out = *in
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceObservation.
func (in *NamespaceObservation) DeepCopy() *NamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(NamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceParameters) DeepCopyInto(out *NamespaceParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceParameters.
func (in *NamespaceParameters) DeepCopy() *NamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSpec) DeepCopyInto(out *NamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSpec.
func (in *NamespaceSpec) DeepCopy() *NamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceStatus) DeepCopyInto(out *NamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceStatus.
func (in *NamespaceStatus) DeepCopy() *NamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceStatus)
	in.DeepCopyInto(out)
	return out
}
