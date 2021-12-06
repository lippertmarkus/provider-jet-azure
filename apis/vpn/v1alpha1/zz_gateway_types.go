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

type BgpSettingsObservation struct {
	BgpPeeringAddress *string `json:"bgpPeeringAddress,omitempty" tf:"bgp_peering_address,omitempty"`
}

type BgpSettingsParameters struct {

	// +kubebuilder:validation:Required
	Asn *int64 `json:"asn" tf:"asn,omitempty"`

	// +kubebuilder:validation:Optional
	Instance0BgpPeeringAddress []Instance0BgpPeeringAddressParameters `json:"instance0BgpPeeringAddress,omitempty" tf:"instance_0_bgp_peering_address,omitempty"`

	// +kubebuilder:validation:Optional
	Instance1BgpPeeringAddress []Instance1BgpPeeringAddressParameters `json:"instance1BgpPeeringAddress,omitempty" tf:"instance_1_bgp_peering_address,omitempty"`

	// +kubebuilder:validation:Required
	PeerWeight *int64 `json:"peerWeight" tf:"peer_weight,omitempty"`
}

type GatewayObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayParameters struct {

	// +kubebuilder:validation:Optional
	BgpSettings []BgpSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleUnit *int64 `json:"scaleUnit,omitempty" tf:"scale_unit,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualHubID *string `json:"virtualHubId" tf:"virtual_hub_id,omitempty"`
}

type Instance0BgpPeeringAddressObservation struct {
	DefaultIps []*string `json:"defaultIps,omitempty" tf:"default_ips,omitempty"`

	IPConfigurationID *string `json:"ipConfigurationId,omitempty" tf:"ip_configuration_id,omitempty"`

	TunnelIps []*string `json:"tunnelIps,omitempty" tf:"tunnel_ips,omitempty"`
}

type Instance0BgpPeeringAddressParameters struct {

	// +kubebuilder:validation:Required
	CustomIps []*string `json:"customIps" tf:"custom_ips,omitempty"`
}

type Instance1BgpPeeringAddressObservation struct {
	DefaultIps []*string `json:"defaultIps,omitempty" tf:"default_ips,omitempty"`

	IPConfigurationID *string `json:"ipConfigurationId,omitempty" tf:"ip_configuration_id,omitempty"`

	TunnelIps []*string `json:"tunnelIps,omitempty" tf:"tunnel_ips,omitempty"`
}

type Instance1BgpPeeringAddressParameters struct {

	// +kubebuilder:validation:Required
	CustomIps []*string `json:"customIps" tf:"custom_ips,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
