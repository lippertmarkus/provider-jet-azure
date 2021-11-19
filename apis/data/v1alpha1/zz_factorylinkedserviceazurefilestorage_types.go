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

type FactoryLinkedServiceAzureFileStorageKeyVaultPasswordObservation struct {
}

type FactoryLinkedServiceAzureFileStorageKeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type FactoryLinkedServiceAzureFileStorageObservation struct {
}

type FactoryLinkedServiceAzureFileStorageParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FileShare *string `json:"fileShare,omitempty" tf:"file_share,omitempty"`

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultPassword []FactoryLinkedServiceAzureFileStorageKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// FactoryLinkedServiceAzureFileStorageSpec defines the desired state of FactoryLinkedServiceAzureFileStorage
type FactoryLinkedServiceAzureFileStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceAzureFileStorageParameters `json:"forProvider"`
}

// FactoryLinkedServiceAzureFileStorageStatus defines the observed state of FactoryLinkedServiceAzureFileStorage.
type FactoryLinkedServiceAzureFileStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceAzureFileStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureFileStorage is the Schema for the FactoryLinkedServiceAzureFileStorages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryLinkedServiceAzureFileStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceAzureFileStorageSpec   `json:"spec"`
	Status            FactoryLinkedServiceAzureFileStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureFileStorageList contains a list of FactoryLinkedServiceAzureFileStorages
type FactoryLinkedServiceAzureFileStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceAzureFileStorage `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceAzureFileStorage_Kind             = "FactoryLinkedServiceAzureFileStorage"
	FactoryLinkedServiceAzureFileStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryLinkedServiceAzureFileStorage_Kind}.String()
	FactoryLinkedServiceAzureFileStorage_KindAPIVersion   = FactoryLinkedServiceAzureFileStorage_Kind + "." + CRDGroupVersion.String()
	FactoryLinkedServiceAzureFileStorage_GroupVersionKind = CRDGroupVersion.WithKind(FactoryLinkedServiceAzureFileStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceAzureFileStorage{}, &FactoryLinkedServiceAzureFileStorageList{})
}
