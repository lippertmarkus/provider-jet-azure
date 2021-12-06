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

type TwinsInstanceObservation struct {
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TwinsInstanceParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TwinsInstanceSpec defines the desired state of TwinsInstance
type TwinsInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TwinsInstanceParameters `json:"forProvider"`
}

// TwinsInstanceStatus defines the observed state of TwinsInstance.
type TwinsInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TwinsInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TwinsInstance is the Schema for the TwinsInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TwinsInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TwinsInstanceSpec   `json:"spec"`
	Status            TwinsInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TwinsInstanceList contains a list of TwinsInstances
type TwinsInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TwinsInstance `json:"items"`
}

// Repository type metadata.
var (
	TwinsInstance_Kind             = "TwinsInstance"
	TwinsInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TwinsInstance_Kind}.String()
	TwinsInstance_KindAPIVersion   = TwinsInstance_Kind + "." + CRDGroupVersion.String()
	TwinsInstance_GroupVersionKind = CRDGroupVersion.WithKind(TwinsInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&TwinsInstance{}, &TwinsInstanceList{})
}
