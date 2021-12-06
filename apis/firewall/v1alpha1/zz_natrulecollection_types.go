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

type NatRuleCollectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NatRuleCollectionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	AzureFirewallName *string `json:"azureFirewallName" tf:"azure_firewall_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Rule []NatRuleCollectionRuleParameters `json:"rule" tf:"rule,omitempty"`
}

type NatRuleCollectionRuleObservation struct {
}

type NatRuleCollectionRuleParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DestinationAddresses []*string `json:"destinationAddresses" tf:"destination_addresses,omitempty"`

	// +kubebuilder:validation:Required
	DestinationPorts []*string `json:"destinationPorts" tf:"destination_ports,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Protocols []*string `json:"protocols" tf:"protocols,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// +kubebuilder:validation:Required
	TranslatedAddress *string `json:"translatedAddress" tf:"translated_address,omitempty"`

	// +kubebuilder:validation:Required
	TranslatedPort *string `json:"translatedPort" tf:"translated_port,omitempty"`
}

// NatRuleCollectionSpec defines the desired state of NatRuleCollection
type NatRuleCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NatRuleCollectionParameters `json:"forProvider"`
}

// NatRuleCollectionStatus defines the observed state of NatRuleCollection.
type NatRuleCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NatRuleCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NatRuleCollection is the Schema for the NatRuleCollections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type NatRuleCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatRuleCollectionSpec   `json:"spec"`
	Status            NatRuleCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatRuleCollectionList contains a list of NatRuleCollections
type NatRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatRuleCollection `json:"items"`
}

// Repository type metadata.
var (
	NatRuleCollection_Kind             = "NatRuleCollection"
	NatRuleCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NatRuleCollection_Kind}.String()
	NatRuleCollection_KindAPIVersion   = NatRuleCollection_Kind + "." + CRDGroupVersion.String()
	NatRuleCollection_GroupVersionKind = CRDGroupVersion.WithKind(NatRuleCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&NatRuleCollection{}, &NatRuleCollectionList{})
}
