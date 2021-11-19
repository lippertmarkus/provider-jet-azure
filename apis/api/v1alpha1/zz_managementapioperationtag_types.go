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

type ManagementApiOperationTagObservation struct {
}

type ManagementApiOperationTagParameters struct {

	// +kubebuilder:validation:Required
	APIOperationID *string `json:"apiOperationId" tf:"api_operation_id,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ManagementApiOperationTagSpec defines the desired state of ManagementApiOperationTag
type ManagementApiOperationTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementApiOperationTagParameters `json:"forProvider"`
}

// ManagementApiOperationTagStatus defines the observed state of ManagementApiOperationTag.
type ManagementApiOperationTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementApiOperationTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementApiOperationTag is the Schema for the ManagementApiOperationTags API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagementApiOperationTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementApiOperationTagSpec   `json:"spec"`
	Status            ManagementApiOperationTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementApiOperationTagList contains a list of ManagementApiOperationTags
type ManagementApiOperationTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementApiOperationTag `json:"items"`
}

// Repository type metadata.
var (
	ManagementApiOperationTag_Kind             = "ManagementApiOperationTag"
	ManagementApiOperationTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementApiOperationTag_Kind}.String()
	ManagementApiOperationTag_KindAPIVersion   = ManagementApiOperationTag_Kind + "." + CRDGroupVersion.String()
	ManagementApiOperationTag_GroupVersionKind = CRDGroupVersion.WithKind(ManagementApiOperationTag_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementApiOperationTag{}, &ManagementApiOperationTagList{})
}
