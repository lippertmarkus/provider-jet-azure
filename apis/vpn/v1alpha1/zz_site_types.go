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

type BgpObservation struct {
}

type BgpParameters struct {

	// +kubebuilder:validation:Required
	Asn *int64 `json:"asn" tf:"asn,omitempty"`

	// +kubebuilder:validation:Required
	PeeringAddress *string `json:"peeringAddress" tf:"peering_address,omitempty"`
}

type LinkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkParameters struct {

	// +kubebuilder:validation:Optional
	Bgp []BgpParameters `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// +kubebuilder:validation:Optional
	SpeedInMbps *int64 `json:"speedInMbps,omitempty" tf:"speed_in_mbps,omitempty"`
}

type SiteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SiteParameters struct {

	// +kubebuilder:validation:Optional
	AddressCidrs []*string `json:"addressCidrs,omitempty" tf:"address_cidrs,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceModel *string `json:"deviceModel,omitempty" tf:"device_model,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceVendor *string `json:"deviceVendor,omitempty" tf:"device_vendor,omitempty"`

	// +kubebuilder:validation:Optional
	Link []LinkParameters `json:"link,omitempty" tf:"link,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualWanID *string `json:"virtualWanId" tf:"virtual_wan_id,omitempty"`
}

// SiteSpec defines the desired state of Site
type SiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteParameters `json:"forProvider"`
}

// SiteStatus defines the observed state of Site.
type SiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Site is the Schema for the Sites API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Site struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteSpec   `json:"spec"`
	Status            SiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteList contains a list of Sites
type SiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Site `json:"items"`
}

// Repository type metadata.
var (
	Site_Kind             = "Site"
	Site_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Site_Kind}.String()
	Site_KindAPIVersion   = Site_Kind + "." + CRDGroupVersion.String()
	Site_GroupVersionKind = CRDGroupVersion.WithKind(Site_Kind)
)

func init() {
	SchemeBuilder.Register(&Site{}, &SiteList{})
}
