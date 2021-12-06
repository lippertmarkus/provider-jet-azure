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

type ConfigurationKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConfigurationKeyParameters struct {

	// +kubebuilder:validation:Required
	ConfigurationStoreID *string `json:"configurationStoreId" tf:"configuration_store_id,omitempty"`

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// +kubebuilder:validation:Optional
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// +kubebuilder:validation:Optional
	VaultKeyReference *string `json:"vaultKeyReference,omitempty" tf:"vault_key_reference,omitempty"`
}

// ConfigurationKeySpec defines the desired state of ConfigurationKey
type ConfigurationKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationKeyParameters `json:"forProvider"`
}

// ConfigurationKeyStatus defines the observed state of ConfigurationKey.
type ConfigurationKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationKey is the Schema for the ConfigurationKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ConfigurationKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationKeySpec   `json:"spec"`
	Status            ConfigurationKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationKeyList contains a list of ConfigurationKeys
type ConfigurationKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationKey `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationKey_Kind             = "ConfigurationKey"
	ConfigurationKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationKey_Kind}.String()
	ConfigurationKey_KindAPIVersion   = ConfigurationKey_Kind + "." + CRDGroupVersion.String()
	ConfigurationKey_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationKey{}, &ConfigurationKeyList{})
}
