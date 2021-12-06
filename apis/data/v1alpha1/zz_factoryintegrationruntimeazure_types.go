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

type FactoryIntegrationRuntimeAzureObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FactoryIntegrationRuntimeAzureParameters struct {

	// +kubebuilder:validation:Optional
	CleanupEnabled *bool `json:"cleanupEnabled,omitempty" tf:"cleanup_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ComputeType *string `json:"computeType,omitempty" tf:"compute_type,omitempty"`

	// +kubebuilder:validation:Optional
	CoreCount *int64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	TimeToLiveMin *int64 `json:"timeToLiveMin,omitempty" tf:"time_to_live_min,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkEnabled *bool `json:"virtualNetworkEnabled,omitempty" tf:"virtual_network_enabled,omitempty"`
}

// FactoryIntegrationRuntimeAzureSpec defines the desired state of FactoryIntegrationRuntimeAzure
type FactoryIntegrationRuntimeAzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryIntegrationRuntimeAzureParameters `json:"forProvider"`
}

// FactoryIntegrationRuntimeAzureStatus defines the observed state of FactoryIntegrationRuntimeAzure.
type FactoryIntegrationRuntimeAzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryIntegrationRuntimeAzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryIntegrationRuntimeAzure is the Schema for the FactoryIntegrationRuntimeAzures API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryIntegrationRuntimeAzure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryIntegrationRuntimeAzureSpec   `json:"spec"`
	Status            FactoryIntegrationRuntimeAzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryIntegrationRuntimeAzureList contains a list of FactoryIntegrationRuntimeAzures
type FactoryIntegrationRuntimeAzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryIntegrationRuntimeAzure `json:"items"`
}

// Repository type metadata.
var (
	FactoryIntegrationRuntimeAzure_Kind             = "FactoryIntegrationRuntimeAzure"
	FactoryIntegrationRuntimeAzure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryIntegrationRuntimeAzure_Kind}.String()
	FactoryIntegrationRuntimeAzure_KindAPIVersion   = FactoryIntegrationRuntimeAzure_Kind + "." + CRDGroupVersion.String()
	FactoryIntegrationRuntimeAzure_GroupVersionKind = CRDGroupVersion.WithKind(FactoryIntegrationRuntimeAzure_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryIntegrationRuntimeAzure{}, &FactoryIntegrationRuntimeAzureList{})
}
