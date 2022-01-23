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

type GatewayAPIObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayAPIParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Required
	GatewayID *string `json:"gatewayId" tf:"gateway_id,omitempty"`
}

// GatewayAPISpec defines the desired state of GatewayAPI
type GatewayAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAPIParameters `json:"forProvider"`
}

// GatewayAPIStatus defines the observed state of GatewayAPI.
type GatewayAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAPI is the Schema for the GatewayAPIs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type GatewayAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAPISpec   `json:"spec"`
	Status            GatewayAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAPIList contains a list of GatewayAPIs
type GatewayAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAPI `json:"items"`
}

// Repository type metadata.
var (
	GatewayAPI_Kind             = "GatewayAPI"
	GatewayAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayAPI_Kind}.String()
	GatewayAPI_KindAPIVersion   = GatewayAPI_Kind + "." + CRDGroupVersion.String()
	GatewayAPI_GroupVersionKind = CRDGroupVersion.WithKind(GatewayAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayAPI{}, &GatewayAPIList{})
}
