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

type ChannelEmailObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ChannelEmailParameters struct {

	// +kubebuilder:validation:Required
	BotName *string `json:"botName" tf:"bot_name,omitempty"`

	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// +kubebuilder:validation:Required
	EmailPasswordSecretRef v1.SecretKeySelector `json:"emailPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// ChannelEmailSpec defines the desired state of ChannelEmail
type ChannelEmailSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChannelEmailParameters `json:"forProvider"`
}

// ChannelEmailStatus defines the observed state of ChannelEmail.
type ChannelEmailStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChannelEmailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelEmail is the Schema for the ChannelEmails API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ChannelEmail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelEmailSpec   `json:"spec"`
	Status            ChannelEmailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelEmailList contains a list of ChannelEmails
type ChannelEmailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChannelEmail `json:"items"`
}

// Repository type metadata.
var (
	ChannelEmail_Kind             = "ChannelEmail"
	ChannelEmail_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ChannelEmail_Kind}.String()
	ChannelEmail_KindAPIVersion   = ChannelEmail_Kind + "." + CRDGroupVersion.String()
	ChannelEmail_GroupVersionKind = CRDGroupVersion.WithKind(ChannelEmail_Kind)
)

func init() {
	SchemeBuilder.Register(&ChannelEmail{}, &ChannelEmailList{})
}
