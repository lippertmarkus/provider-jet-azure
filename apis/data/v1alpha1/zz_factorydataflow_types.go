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

type DatasetObservation struct {
}

type DatasetParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FactoryDataFlowObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FactoryDataFlowParameters struct {

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryID *string `json:"dataFactoryId" tf:"data_factory_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Script *string `json:"script" tf:"script,omitempty"`

	// +kubebuilder:validation:Required
	Sink []SinkParameters `json:"sink" tf:"sink,omitempty"`

	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	Transformation []TransformationParameters `json:"transformation,omitempty" tf:"transformation,omitempty"`
}

type SchemaLinkedServiceObservation struct {
}

type SchemaLinkedServiceParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkLinkedServiceObservation struct {
}

type SinkLinkedServiceParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkObservation struct {
}

type SinkParameters struct {

	// +kubebuilder:validation:Optional
	Dataset []DatasetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LinkedService []SinkLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaLinkedService []SchemaLinkedServiceParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceDatasetObservation struct {
}

type SourceDatasetParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceLinkedServiceObservation struct {
}

type SourceLinkedServiceParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Optional
	Dataset []SourceDatasetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LinkedService []SourceLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaLinkedService []SourceSchemaLinkedServiceParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceSchemaLinkedServiceObservation struct {
}

type SourceSchemaLinkedServiceParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationObservation struct {
}

type TransformationParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// FactoryDataFlowSpec defines the desired state of FactoryDataFlow
type FactoryDataFlowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryDataFlowParameters `json:"forProvider"`
}

// FactoryDataFlowStatus defines the observed state of FactoryDataFlow.
type FactoryDataFlowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryDataFlowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryDataFlow is the Schema for the FactoryDataFlows API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryDataFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryDataFlowSpec   `json:"spec"`
	Status            FactoryDataFlowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryDataFlowList contains a list of FactoryDataFlows
type FactoryDataFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryDataFlow `json:"items"`
}

// Repository type metadata.
var (
	FactoryDataFlow_Kind             = "FactoryDataFlow"
	FactoryDataFlow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryDataFlow_Kind}.String()
	FactoryDataFlow_KindAPIVersion   = FactoryDataFlow_Kind + "." + CRDGroupVersion.String()
	FactoryDataFlow_GroupVersionKind = CRDGroupVersion.WithKind(FactoryDataFlow_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryDataFlow{}, &FactoryDataFlowList{})
}
