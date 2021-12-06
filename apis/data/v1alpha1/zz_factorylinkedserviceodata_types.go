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

type BasicAuthenticationObservation struct {
}

type BasicAuthenticationParameters struct {

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type FactoryLinkedServiceOdataObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FactoryLinkedServiceOdataParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	BasicAuthentication []BasicAuthenticationParameters `json:"basicAuthentication,omitempty" tf:"basic_authentication,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// FactoryLinkedServiceOdataSpec defines the desired state of FactoryLinkedServiceOdata
type FactoryLinkedServiceOdataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceOdataParameters `json:"forProvider"`
}

// FactoryLinkedServiceOdataStatus defines the observed state of FactoryLinkedServiceOdata.
type FactoryLinkedServiceOdataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceOdataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceOdata is the Schema for the FactoryLinkedServiceOdatas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryLinkedServiceOdata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceOdataSpec   `json:"spec"`
	Status            FactoryLinkedServiceOdataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceOdataList contains a list of FactoryLinkedServiceOdatas
type FactoryLinkedServiceOdataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceOdata `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceOdata_Kind             = "FactoryLinkedServiceOdata"
	FactoryLinkedServiceOdata_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryLinkedServiceOdata_Kind}.String()
	FactoryLinkedServiceOdata_KindAPIVersion   = FactoryLinkedServiceOdata_Kind + "." + CRDGroupVersion.String()
	FactoryLinkedServiceOdata_GroupVersionKind = CRDGroupVersion.WithKind(FactoryLinkedServiceOdata_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceOdata{}, &FactoryLinkedServiceOdataList{})
}
