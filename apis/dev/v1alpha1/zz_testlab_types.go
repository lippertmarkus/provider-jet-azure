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

type TestLabObservation struct {
	ArtifactsStorageAccountID *string `json:"artifactsStorageAccountId,omitempty" tf:"artifacts_storage_account_id,omitempty"`

	DefaultPremiumStorageAccountID *string `json:"defaultPremiumStorageAccountId,omitempty" tf:"default_premium_storage_account_id,omitempty"`

	DefaultStorageAccountID *string `json:"defaultStorageAccountId,omitempty" tf:"default_storage_account_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	PremiumDataDiskStorageAccountID *string `json:"premiumDataDiskStorageAccountId,omitempty" tf:"premium_data_disk_storage_account_id,omitempty"`

	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type TestLabParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TestLabSpec defines the desired state of TestLab
type TestLabSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TestLabParameters `json:"forProvider"`
}

// TestLabStatus defines the observed state of TestLab.
type TestLabStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TestLabObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TestLab is the Schema for the TestLabs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TestLab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TestLabSpec   `json:"spec"`
	Status            TestLabStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TestLabList contains a list of TestLabs
type TestLabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestLab `json:"items"`
}

// Repository type metadata.
var (
	TestLab_Kind             = "TestLab"
	TestLab_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TestLab_Kind}.String()
	TestLab_KindAPIVersion   = TestLab_Kind + "." + CRDGroupVersion.String()
	TestLab_GroupVersionKind = CRDGroupVersion.WithKind(TestLab_Kind)
)

func init() {
	SchemeBuilder.Register(&TestLab{}, &TestLabList{})
}
