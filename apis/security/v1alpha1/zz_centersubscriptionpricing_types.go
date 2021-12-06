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

type CenterSubscriptionPricingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CenterSubscriptionPricingParameters struct {

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

// CenterSubscriptionPricingSpec defines the desired state of CenterSubscriptionPricing
type CenterSubscriptionPricingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CenterSubscriptionPricingParameters `json:"forProvider"`
}

// CenterSubscriptionPricingStatus defines the observed state of CenterSubscriptionPricing.
type CenterSubscriptionPricingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CenterSubscriptionPricingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CenterSubscriptionPricing is the Schema for the CenterSubscriptionPricings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type CenterSubscriptionPricing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CenterSubscriptionPricingSpec   `json:"spec"`
	Status            CenterSubscriptionPricingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CenterSubscriptionPricingList contains a list of CenterSubscriptionPricings
type CenterSubscriptionPricingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CenterSubscriptionPricing `json:"items"`
}

// Repository type metadata.
var (
	CenterSubscriptionPricing_Kind             = "CenterSubscriptionPricing"
	CenterSubscriptionPricing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CenterSubscriptionPricing_Kind}.String()
	CenterSubscriptionPricing_KindAPIVersion   = CenterSubscriptionPricing_Kind + "." + CRDGroupVersion.String()
	CenterSubscriptionPricing_GroupVersionKind = CRDGroupVersion.WithKind(CenterSubscriptionPricing_Kind)
)

func init() {
	SchemeBuilder.Register(&CenterSubscriptionPricing{}, &CenterSubscriptionPricingList{})
}
