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

type ManagedDiskObservation struct {
}

type ManagedDiskParameters struct {

	// +kubebuilder:validation:Required
	DiskID *string `json:"diskId" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Required
	StagingStorageAccountID *string `json:"stagingStorageAccountId" tf:"staging_storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetDiskEncryptionSetID *string `json:"targetDiskEncryptionSetId,omitempty" tf:"target_disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetDiskType *string `json:"targetDiskType" tf:"target_disk_type,omitempty"`

	// +kubebuilder:validation:Required
	TargetReplicaDiskType *string `json:"targetReplicaDiskType" tf:"target_replica_disk_type,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceGroupID *string `json:"targetResourceGroupId" tf:"target_resource_group_id,omitempty"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	RecoveryPublicIPAddressID *string `json:"recoveryPublicIpAddressId,omitempty" tf:"recovery_public_ip_address_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceNetworkInterfaceID *string `json:"sourceNetworkInterfaceId,omitempty" tf:"source_network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetStaticIP *string `json:"targetStaticIp,omitempty" tf:"target_static_ip,omitempty"`

	// +kubebuilder:validation:Optional
	TargetSubnetName *string `json:"targetSubnetName,omitempty" tf:"target_subnet_name,omitempty"`
}

type RecoveryReplicatedVmObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RecoveryReplicatedVmParameters struct {

	// +kubebuilder:validation:Optional
	ManagedDisk []ManagedDiskParameters `json:"managedDisk,omitempty" tf:"managed_disk,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryReplicationPolicyID *string `json:"recoveryReplicationPolicyId" tf:"recovery_replication_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceRecoveryFabricName *string `json:"sourceRecoveryFabricName" tf:"source_recovery_fabric_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceRecoveryProtectionContainerName *string `json:"sourceRecoveryProtectionContainerName" tf:"source_recovery_protection_container_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceVMID *string `json:"sourceVmId" tf:"source_vm_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetAvailabilitySetID *string `json:"targetAvailabilitySetId,omitempty" tf:"target_availability_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetNetworkID *string `json:"targetNetworkId,omitempty" tf:"target_network_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetRecoveryFabricID *string `json:"targetRecoveryFabricId" tf:"target_recovery_fabric_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetRecoveryProtectionContainerID *string `json:"targetRecoveryProtectionContainerId" tf:"target_recovery_protection_container_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceGroupID *string `json:"targetResourceGroupId" tf:"target_resource_group_id,omitempty"`
}

// RecoveryReplicatedVmSpec defines the desired state of RecoveryReplicatedVm
type RecoveryReplicatedVmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecoveryReplicatedVmParameters `json:"forProvider"`
}

// RecoveryReplicatedVmStatus defines the observed state of RecoveryReplicatedVm.
type RecoveryReplicatedVmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecoveryReplicatedVmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryReplicatedVm is the Schema for the RecoveryReplicatedVms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RecoveryReplicatedVm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryReplicatedVmSpec   `json:"spec"`
	Status            RecoveryReplicatedVmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryReplicatedVmList contains a list of RecoveryReplicatedVms
type RecoveryReplicatedVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecoveryReplicatedVm `json:"items"`
}

// Repository type metadata.
var (
	RecoveryReplicatedVm_Kind             = "RecoveryReplicatedVm"
	RecoveryReplicatedVm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecoveryReplicatedVm_Kind}.String()
	RecoveryReplicatedVm_KindAPIVersion   = RecoveryReplicatedVm_Kind + "." + CRDGroupVersion.String()
	RecoveryReplicatedVm_GroupVersionKind = CRDGroupVersion.WithKind(RecoveryReplicatedVm_Kind)
)

func init() {
	SchemeBuilder.Register(&RecoveryReplicatedVm{}, &RecoveryReplicatedVmList{})
}
