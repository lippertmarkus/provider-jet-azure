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

type ServiceCustomHostnameBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	VirtualIP *string `json:"virtualIp,omitempty" tf:"virtual_ip,omitempty"`
}

type ServiceCustomHostnameBindingParameters struct {

	// +kubebuilder:validation:Required
	AppServiceName *string `json:"appServiceName" tf:"app_service_name,omitempty"`

	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SslState *string `json:"sslState,omitempty" tf:"ssl_state,omitempty"`

	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

// ServiceCustomHostnameBindingSpec defines the desired state of ServiceCustomHostnameBinding
type ServiceCustomHostnameBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceCustomHostnameBindingParameters `json:"forProvider"`
}

// ServiceCustomHostnameBindingStatus defines the observed state of ServiceCustomHostnameBinding.
type ServiceCustomHostnameBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceCustomHostnameBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceCustomHostnameBinding is the Schema for the ServiceCustomHostnameBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ServiceCustomHostnameBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceCustomHostnameBindingSpec   `json:"spec"`
	Status            ServiceCustomHostnameBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceCustomHostnameBindingList contains a list of ServiceCustomHostnameBindings
type ServiceCustomHostnameBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceCustomHostnameBinding `json:"items"`
}

// Repository type metadata.
var (
	ServiceCustomHostnameBinding_Kind             = "ServiceCustomHostnameBinding"
	ServiceCustomHostnameBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceCustomHostnameBinding_Kind}.String()
	ServiceCustomHostnameBinding_KindAPIVersion   = ServiceCustomHostnameBinding_Kind + "." + CRDGroupVersion.String()
	ServiceCustomHostnameBinding_GroupVersionKind = CRDGroupVersion.WithKind(ServiceCustomHostnameBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceCustomHostnameBinding{}, &ServiceCustomHostnameBindingList{})
}
