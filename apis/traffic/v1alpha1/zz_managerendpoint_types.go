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

type CustomHeaderObservation struct {
}

type CustomHeaderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ManagerEndpointObservation struct {
	EndpointMonitorStatus *string `json:"endpointMonitorStatus,omitempty" tf:"endpoint_monitor_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagerEndpointParameters struct {

	// +kubebuilder:validation:Optional
	CustomHeader []CustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointStatus *string `json:"endpointStatus,omitempty" tf:"endpoint_status,omitempty"`

	// +kubebuilder:validation:Optional
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// +kubebuilder:validation:Optional
	MinChildEndpoints *int64 `json:"minChildEndpoints,omitempty" tf:"min_child_endpoints,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumRequiredChildEndpointsIPv4 *int64 `json:"minimumRequiredChildEndpointsIpv4,omitempty" tf:"minimum_required_child_endpoints_ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumRequiredChildEndpointsIPv6 *int64 `json:"minimumRequiredChildEndpointsIpv6,omitempty" tf:"minimum_required_child_endpoints_ipv6,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	ProfileName *string `json:"profileName" tf:"profile_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Subnet []SubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SubnetObservation struct {
}

type SubnetParameters struct {

	// +kubebuilder:validation:Required
	First *string `json:"first" tf:"first,omitempty"`

	// +kubebuilder:validation:Optional
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *int64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

// ManagerEndpointSpec defines the desired state of ManagerEndpoint
type ManagerEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerEndpointParameters `json:"forProvider"`
}

// ManagerEndpointStatus defines the observed state of ManagerEndpoint.
type ManagerEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerEndpoint is the Schema for the ManagerEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagerEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerEndpointSpec   `json:"spec"`
	Status            ManagerEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerEndpointList contains a list of ManagerEndpoints
type ManagerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ManagerEndpoint_Kind             = "ManagerEndpoint"
	ManagerEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerEndpoint_Kind}.String()
	ManagerEndpoint_KindAPIVersion   = ManagerEndpoint_Kind + "." + CRDGroupVersion.String()
	ManagerEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ManagerEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerEndpoint{}, &ManagerEndpointList{})
}
