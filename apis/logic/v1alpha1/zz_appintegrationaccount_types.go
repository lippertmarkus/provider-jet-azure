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

type AppIntegrationAccountObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppIntegrationAccountParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AppIntegrationAccountSpec defines the desired state of AppIntegrationAccount
type AppIntegrationAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountParameters `json:"forProvider"`
}

// AppIntegrationAccountStatus defines the observed state of AppIntegrationAccount.
type AppIntegrationAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccount is the Schema for the AppIntegrationAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppIntegrationAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppIntegrationAccountSpec   `json:"spec"`
	Status            AppIntegrationAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountList contains a list of AppIntegrationAccounts
type AppIntegrationAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccount `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccount_Kind             = "AppIntegrationAccount"
	AppIntegrationAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccount_Kind}.String()
	AppIntegrationAccount_KindAPIVersion   = AppIntegrationAccount_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccount_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccount{}, &AppIntegrationAccountList{})
}
