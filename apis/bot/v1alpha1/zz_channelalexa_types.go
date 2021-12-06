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

type ChannelAlexaObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ChannelAlexaParameters struct {

	// +kubebuilder:validation:Required
	BotName *string `json:"botName" tf:"bot_name,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkillID *string `json:"skillId" tf:"skill_id,omitempty"`
}

// ChannelAlexaSpec defines the desired state of ChannelAlexa
type ChannelAlexaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChannelAlexaParameters `json:"forProvider"`
}

// ChannelAlexaStatus defines the observed state of ChannelAlexa.
type ChannelAlexaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChannelAlexaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelAlexa is the Schema for the ChannelAlexas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ChannelAlexa struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelAlexaSpec   `json:"spec"`
	Status            ChannelAlexaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelAlexaList contains a list of ChannelAlexas
type ChannelAlexaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChannelAlexa `json:"items"`
}

// Repository type metadata.
var (
	ChannelAlexa_Kind             = "ChannelAlexa"
	ChannelAlexa_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ChannelAlexa_Kind}.String()
	ChannelAlexa_KindAPIVersion   = ChannelAlexa_Kind + "." + CRDGroupVersion.String()
	ChannelAlexa_GroupVersionKind = CRDGroupVersion.WithKind(ChannelAlexa_Kind)
)

func init() {
	SchemeBuilder.Register(&ChannelAlexa{}, &ChannelAlexaList{})
}
