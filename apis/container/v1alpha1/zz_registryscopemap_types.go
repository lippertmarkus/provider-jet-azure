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

type RegistryScopeMapObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegistryScopeMapParameters struct {

	// +kubebuilder:validation:Required
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// +kubebuilder:validation:Required
	ContainerRegistryName *string `json:"containerRegistryName" tf:"container_registry_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// RegistryScopeMapSpec defines the desired state of RegistryScopeMap
type RegistryScopeMapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryScopeMapParameters `json:"forProvider"`
}

// RegistryScopeMapStatus defines the observed state of RegistryScopeMap.
type RegistryScopeMapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryScopeMapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryScopeMap is the Schema for the RegistryScopeMaps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RegistryScopeMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistryScopeMapSpec   `json:"spec"`
	Status            RegistryScopeMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryScopeMapList contains a list of RegistryScopeMaps
type RegistryScopeMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryScopeMap `json:"items"`
}

// Repository type metadata.
var (
	RegistryScopeMap_Kind             = "RegistryScopeMap"
	RegistryScopeMap_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryScopeMap_Kind}.String()
	RegistryScopeMap_KindAPIVersion   = RegistryScopeMap_Kind + "." + CRDGroupVersion.String()
	RegistryScopeMap_GroupVersionKind = CRDGroupVersion.WithKind(RegistryScopeMap_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryScopeMap{}, &RegistryScopeMapList{})
}
