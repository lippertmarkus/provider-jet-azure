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

type EndpointCustomDomainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointCustomDomainParameters struct {

	// +kubebuilder:validation:Required
	CdnEndpointID *string `json:"cdnEndpointId" tf:"cdn_endpoint_id,omitempty"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// EndpointCustomDomainSpec defines the desired state of EndpointCustomDomain
type EndpointCustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointCustomDomainParameters `json:"forProvider"`
}

// EndpointCustomDomainStatus defines the observed state of EndpointCustomDomain.
type EndpointCustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointCustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointCustomDomain is the Schema for the EndpointCustomDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EndpointCustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointCustomDomainSpec   `json:"spec"`
	Status            EndpointCustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointCustomDomainList contains a list of EndpointCustomDomains
type EndpointCustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointCustomDomain `json:"items"`
}

// Repository type metadata.
var (
	EndpointCustomDomain_Kind             = "EndpointCustomDomain"
	EndpointCustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointCustomDomain_Kind}.String()
	EndpointCustomDomain_KindAPIVersion   = EndpointCustomDomain_Kind + "." + CRDGroupVersion.String()
	EndpointCustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(EndpointCustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointCustomDomain{}, &EndpointCustomDomainList{})
}
