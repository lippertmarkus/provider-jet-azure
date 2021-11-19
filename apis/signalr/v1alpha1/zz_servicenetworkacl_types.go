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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivateEndpointObservation struct {
}

type PrivateEndpointParameters struct {

	// +kubebuilder:validation:Optional
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// +kubebuilder:validation:Optional
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type PublicNetworkObservation struct {
}

type PublicNetworkParameters struct {

	// +kubebuilder:validation:Optional
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// +kubebuilder:validation:Optional
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`
}

type ServiceNetworkAclObservation struct {
}

type ServiceNetworkAclParameters struct {

	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateEndpoint []PrivateEndpointParameters `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	PublicNetwork []PublicNetworkParameters `json:"publicNetwork" tf:"public_network,omitempty"`

	// +kubebuilder:validation:Required
	SignalrServiceID *string `json:"signalrServiceId" tf:"signalr_service_id,omitempty"`
}

// ServiceNetworkAclSpec defines the desired state of ServiceNetworkAcl
type ServiceNetworkAclSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceNetworkAclParameters `json:"forProvider"`
}

// ServiceNetworkAclStatus defines the observed state of ServiceNetworkAcl.
type ServiceNetworkAclStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceNetworkAclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkAcl is the Schema for the ServiceNetworkAcls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ServiceNetworkAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceNetworkAclSpec   `json:"spec"`
	Status            ServiceNetworkAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkAclList contains a list of ServiceNetworkAcls
type ServiceNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkAcl `json:"items"`
}

// Repository type metadata.
var (
	ServiceNetworkAcl_Kind             = "ServiceNetworkAcl"
	ServiceNetworkAcl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceNetworkAcl_Kind}.String()
	ServiceNetworkAcl_KindAPIVersion   = ServiceNetworkAcl_Kind + "." + CRDGroupVersion.String()
	ServiceNetworkAcl_GroupVersionKind = CRDGroupVersion.WithKind(ServiceNetworkAcl_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceNetworkAcl{}, &ServiceNetworkAclList{})
}
