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

type VaultManagedHardwareSecurityModuleObservation struct {
	HsmURI *string `json:"hsmUri,omitempty" tf:"hsm_uri,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VaultManagedHardwareSecurityModuleParameters struct {

	// +kubebuilder:validation:Required
	AdminObjectIds []*string `json:"adminObjectIds" tf:"admin_object_ids,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PurgeProtectionEnabled *bool `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	SoftDeleteRetentionDays *int64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// VaultManagedHardwareSecurityModuleSpec defines the desired state of VaultManagedHardwareSecurityModule
type VaultManagedHardwareSecurityModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultManagedHardwareSecurityModuleParameters `json:"forProvider"`
}

// VaultManagedHardwareSecurityModuleStatus defines the observed state of VaultManagedHardwareSecurityModule.
type VaultManagedHardwareSecurityModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultManagedHardwareSecurityModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultManagedHardwareSecurityModule is the Schema for the VaultManagedHardwareSecurityModules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VaultManagedHardwareSecurityModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultManagedHardwareSecurityModuleSpec   `json:"spec"`
	Status            VaultManagedHardwareSecurityModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultManagedHardwareSecurityModuleList contains a list of VaultManagedHardwareSecurityModules
type VaultManagedHardwareSecurityModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultManagedHardwareSecurityModule `json:"items"`
}

// Repository type metadata.
var (
	VaultManagedHardwareSecurityModule_Kind             = "VaultManagedHardwareSecurityModule"
	VaultManagedHardwareSecurityModule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultManagedHardwareSecurityModule_Kind}.String()
	VaultManagedHardwareSecurityModule_KindAPIVersion   = VaultManagedHardwareSecurityModule_Kind + "." + CRDGroupVersion.String()
	VaultManagedHardwareSecurityModule_GroupVersionKind = CRDGroupVersion.WithKind(VaultManagedHardwareSecurityModule_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultManagedHardwareSecurityModule{}, &VaultManagedHardwareSecurityModuleList{})
}
