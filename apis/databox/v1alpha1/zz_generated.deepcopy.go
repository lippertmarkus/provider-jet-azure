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
func (in *ContactObservation) DeepCopyInto(out *ContactObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactObservation.
func (in *ContactObservation) DeepCopy() *ContactObservation {
	if in == nil {
		return nil
	}
	out := new(ContactObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactParameters) DeepCopyInto(out *ContactParameters) {
	*out = *in
	if in.CompanyName != nil {
		in, out := &in.CompanyName, &out.CompanyName
		*out = new(string)
		**out = **in
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactParameters.
func (in *ContactParameters) DeepCopy() *ContactParameters {
	if in == nil {
		return nil
	}
	out := new(ContactParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePropertiesObservation) DeepCopyInto(out *DevicePropertiesObservation) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
		**out = **in
	}
	if in.ConfiguredRoleTypes != nil {
		in, out := &in.ConfiguredRoleTypes, &out.ConfiguredRoleTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Culture != nil {
		in, out := &in.Culture, &out.Culture
		*out = new(string)
		**out = **in
	}
	if in.HcsVersion != nil {
		in, out := &in.HcsVersion, &out.HcsVersion
		*out = new(string)
		**out = **in
	}
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int64)
		**out = **in
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.SoftwareVersion != nil {
		in, out := &in.SoftwareVersion, &out.SoftwareVersion
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePropertiesObservation.
func (in *DevicePropertiesObservation) DeepCopy() *DevicePropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(DevicePropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePropertiesParameters) DeepCopyInto(out *DevicePropertiesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePropertiesParameters.
func (in *DevicePropertiesParameters) DeepCopy() *DevicePropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(DevicePropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDevice) DeepCopyInto(out *EdgeDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDevice.
func (in *EdgeDevice) DeepCopy() *EdgeDevice {
	if in == nil {
		return nil
	}
	out := new(EdgeDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceList) DeepCopyInto(out *EdgeDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceList.
func (in *EdgeDeviceList) DeepCopy() *EdgeDeviceList {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceObservation) DeepCopyInto(out *EdgeDeviceObservation) {
	*out = *in
	if in.DeviceProperties != nil {
		in, out := &in.DeviceProperties, &out.DeviceProperties
		*out = make([]DevicePropertiesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceObservation.
func (in *EdgeDeviceObservation) DeepCopy() *EdgeDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceParameters) DeepCopyInto(out *EdgeDeviceParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceParameters.
func (in *EdgeDeviceParameters) DeepCopy() *EdgeDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceSpec) DeepCopyInto(out *EdgeDeviceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceSpec.
func (in *EdgeDeviceSpec) DeepCopy() *EdgeDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceStatus) DeepCopyInto(out *EdgeDeviceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceStatus.
func (in *EdgeDeviceStatus) DeepCopy() *EdgeDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeOrder) DeepCopyInto(out *EdgeOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeOrder.
func (in *EdgeOrder) DeepCopy() *EdgeOrder {
	if in == nil {
		return nil
	}
	out := new(EdgeOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeOrderList) DeepCopyInto(out *EdgeOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeOrderList.
func (in *EdgeOrderList) DeepCopy() *EdgeOrderList {
	if in == nil {
		return nil
	}
	out := new(EdgeOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeOrderObservation) DeepCopyInto(out *EdgeOrderObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReturnTracking != nil {
		in, out := &in.ReturnTracking, &out.ReturnTracking
		*out = make([]ReturnTrackingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.ShipmentHistory != nil {
		in, out := &in.ShipmentHistory, &out.ShipmentHistory
		*out = make([]ShipmentHistoryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShipmentTracking != nil {
		in, out := &in.ShipmentTracking, &out.ShipmentTracking
		*out = make([]ShipmentTrackingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]StatusObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeOrderObservation.
func (in *EdgeOrderObservation) DeepCopy() *EdgeOrderObservation {
	if in == nil {
		return nil
	}
	out := new(EdgeOrderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeOrderParameters) DeepCopyInto(out *EdgeOrderParameters) {
	*out = *in
	if in.Contact != nil {
		in, out := &in.Contact, &out.Contact
		*out = make([]ContactParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ShipmentAddress != nil {
		in, out := &in.ShipmentAddress, &out.ShipmentAddress
		*out = make([]ShipmentAddressParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeOrderParameters.
func (in *EdgeOrderParameters) DeepCopy() *EdgeOrderParameters {
	if in == nil {
		return nil
	}
	out := new(EdgeOrderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeOrderSpec) DeepCopyInto(out *EdgeOrderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeOrderSpec.
func (in *EdgeOrderSpec) DeepCopy() *EdgeOrderSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeOrderStatus) DeepCopyInto(out *EdgeOrderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeOrderStatus.
func (in *EdgeOrderStatus) DeepCopy() *EdgeOrderStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReturnTrackingObservation) DeepCopyInto(out *ReturnTrackingObservation) {
	*out = *in
	if in.CarrierName != nil {
		in, out := &in.CarrierName, &out.CarrierName
		*out = new(string)
		**out = **in
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.TrackingID != nil {
		in, out := &in.TrackingID, &out.TrackingID
		*out = new(string)
		**out = **in
	}
	if in.TrackingURL != nil {
		in, out := &in.TrackingURL, &out.TrackingURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReturnTrackingObservation.
func (in *ReturnTrackingObservation) DeepCopy() *ReturnTrackingObservation {
	if in == nil {
		return nil
	}
	out := new(ReturnTrackingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReturnTrackingParameters) DeepCopyInto(out *ReturnTrackingParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReturnTrackingParameters.
func (in *ReturnTrackingParameters) DeepCopy() *ReturnTrackingParameters {
	if in == nil {
		return nil
	}
	out := new(ReturnTrackingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentAddressObservation) DeepCopyInto(out *ShipmentAddressObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentAddressObservation.
func (in *ShipmentAddressObservation) DeepCopy() *ShipmentAddressObservation {
	if in == nil {
		return nil
	}
	out := new(ShipmentAddressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentAddressParameters) DeepCopyInto(out *ShipmentAddressParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.City != nil {
		in, out := &in.City, &out.City
		*out = new(string)
		**out = **in
	}
	if in.Country != nil {
		in, out := &in.Country, &out.Country
		*out = new(string)
		**out = **in
	}
	if in.PostalCode != nil {
		in, out := &in.PostalCode, &out.PostalCode
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentAddressParameters.
func (in *ShipmentAddressParameters) DeepCopy() *ShipmentAddressParameters {
	if in == nil {
		return nil
	}
	out := new(ShipmentAddressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentHistoryObservation) DeepCopyInto(out *ShipmentHistoryObservation) {
	*out = *in
	if in.AdditionalDetails != nil {
		in, out := &in.AdditionalDetails, &out.AdditionalDetails
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
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = new(string)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentHistoryObservation.
func (in *ShipmentHistoryObservation) DeepCopy() *ShipmentHistoryObservation {
	if in == nil {
		return nil
	}
	out := new(ShipmentHistoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentHistoryParameters) DeepCopyInto(out *ShipmentHistoryParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentHistoryParameters.
func (in *ShipmentHistoryParameters) DeepCopy() *ShipmentHistoryParameters {
	if in == nil {
		return nil
	}
	out := new(ShipmentHistoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentTrackingObservation) DeepCopyInto(out *ShipmentTrackingObservation) {
	*out = *in
	if in.CarrierName != nil {
		in, out := &in.CarrierName, &out.CarrierName
		*out = new(string)
		**out = **in
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.TrackingID != nil {
		in, out := &in.TrackingID, &out.TrackingID
		*out = new(string)
		**out = **in
	}
	if in.TrackingURL != nil {
		in, out := &in.TrackingURL, &out.TrackingURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentTrackingObservation.
func (in *ShipmentTrackingObservation) DeepCopy() *ShipmentTrackingObservation {
	if in == nil {
		return nil
	}
	out := new(ShipmentTrackingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentTrackingParameters) DeepCopyInto(out *ShipmentTrackingParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentTrackingParameters.
func (in *ShipmentTrackingParameters) DeepCopy() *ShipmentTrackingParameters {
	if in == nil {
		return nil
	}
	out := new(ShipmentTrackingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusObservation) DeepCopyInto(out *StatusObservation) {
	*out = *in
	if in.AdditionalDetails != nil {
		in, out := &in.AdditionalDetails, &out.AdditionalDetails
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
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = new(string)
		**out = **in
	}
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = new(string)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusObservation.
func (in *StatusObservation) DeepCopy() *StatusObservation {
	if in == nil {
		return nil
	}
	out := new(StatusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusParameters) DeepCopyInto(out *StatusParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusParameters.
func (in *StatusParameters) DeepCopy() *StatusParameters {
	if in == nil {
		return nil
	}
	out := new(StatusParameters)
	in.DeepCopyInto(out)
	return out
}
