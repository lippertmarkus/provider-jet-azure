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

type ProtectedVmObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProtectedVmParameters struct {

	// +kubebuilder:validation:Required
	BackupPolicyID *string `json:"backupPolicyId" tf:"backup_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceVMID *string `json:"sourceVmId" tf:"source_vm_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProtectedVmSpec defines the desired state of ProtectedVm
type ProtectedVmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtectedVmParameters `json:"forProvider"`
}

// ProtectedVmStatus defines the observed state of ProtectedVm.
type ProtectedVmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtectedVmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectedVm is the Schema for the ProtectedVms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ProtectedVm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtectedVmSpec   `json:"spec"`
	Status            ProtectedVmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectedVmList contains a list of ProtectedVms
type ProtectedVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtectedVm `json:"items"`
}

// Repository type metadata.
var (
	ProtectedVm_Kind             = "ProtectedVm"
	ProtectedVm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtectedVm_Kind}.String()
	ProtectedVm_KindAPIVersion   = ProtectedVm_Kind + "." + CRDGroupVersion.String()
	ProtectedVm_GroupVersionKind = CRDGroupVersion.WithKind(ProtectedVm_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtectedVm{}, &ProtectedVmList{})
}
