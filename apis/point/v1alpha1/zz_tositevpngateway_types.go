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

type ConnectionConfigurationObservation struct {
}

type ConnectionConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// +kubebuilder:validation:Required
	VpnClientAddressPool []VpnClientAddressPoolParameters `json:"vpnClientAddressPool" tf:"vpn_client_address_pool,omitempty"`
}

type PropagatedRouteTableObservation struct {
}

type PropagatedRouteTableParameters struct {

	// +kubebuilder:validation:Required
	Ids []*string `json:"ids" tf:"ids,omitempty"`

	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type RouteObservation struct {
}

type RouteParameters struct {

	// +kubebuilder:validation:Required
	AssociatedRouteTableID *string `json:"associatedRouteTableId" tf:"associated_route_table_id,omitempty"`

	// +kubebuilder:validation:Optional
	PropagatedRouteTable []PropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type ToSiteVpnGatewayObservation struct {
}

type ToSiteVpnGatewayParameters struct {

	// +kubebuilder:validation:Required
	ConnectionConfiguration []ConnectionConfigurationParameters `json:"connectionConfiguration" tf:"connection_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	ScaleUnit *int64 `json:"scaleUnit" tf:"scale_unit,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualHubID *string `json:"virtualHubId" tf:"virtual_hub_id,omitempty"`

	// +kubebuilder:validation:Required
	VpnServerConfigurationID *string `json:"vpnServerConfigurationId" tf:"vpn_server_configuration_id,omitempty"`
}

type VpnClientAddressPoolObservation struct {
}

type VpnClientAddressPoolParameters struct {

	// +kubebuilder:validation:Required
	AddressPrefixes []*string `json:"addressPrefixes" tf:"address_prefixes,omitempty"`
}

// ToSiteVpnGatewaySpec defines the desired state of ToSiteVpnGateway
type ToSiteVpnGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ToSiteVpnGatewayParameters `json:"forProvider"`
}

// ToSiteVpnGatewayStatus defines the observed state of ToSiteVpnGateway.
type ToSiteVpnGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ToSiteVpnGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ToSiteVpnGateway is the Schema for the ToSiteVpnGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ToSiteVpnGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ToSiteVpnGatewaySpec   `json:"spec"`
	Status            ToSiteVpnGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ToSiteVpnGatewayList contains a list of ToSiteVpnGateways
type ToSiteVpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToSiteVpnGateway `json:"items"`
}

// Repository type metadata.
var (
	ToSiteVpnGateway_Kind             = "ToSiteVpnGateway"
	ToSiteVpnGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ToSiteVpnGateway_Kind}.String()
	ToSiteVpnGateway_KindAPIVersion   = ToSiteVpnGateway_Kind + "." + CRDGroupVersion.String()
	ToSiteVpnGateway_GroupVersionKind = CRDGroupVersion.WithKind(ToSiteVpnGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&ToSiteVpnGateway{}, &ToSiteVpnGatewayList{})
}
