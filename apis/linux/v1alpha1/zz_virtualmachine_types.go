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

type AdditionalCapabilitiesObservation struct {
}

type AdditionalCapabilitiesParameters struct {

	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type AdminSSHKeyObservation struct {
}

type AdminSSHKeyParameters struct {

	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type BootDiagnosticsObservation struct {
}

type BootDiagnosticsParameters struct {

	// +kubebuilder:validation:Optional
	StorageAccountURI *string `json:"storageAccountUri,omitempty" tf:"storage_account_uri,omitempty"`
}

type CertificateObservation struct {
}

type CertificateParameters struct {

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type DiffDiskSettingsObservation struct {
}

type DiffDiskSettingsParameters struct {

	// +kubebuilder:validation:Required
	Option *string `json:"option" tf:"option,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type OsDiskObservation struct {
}

type OsDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// +kubebuilder:validation:Optional
	DiffDiskSettings []DiffDiskSettingsParameters `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type PlanObservation struct {
}

type PlanParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

type SecretObservation struct {
}

type SecretParameters struct {

	// +kubebuilder:validation:Required
	Certificate []CertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`
}

type SourceImageReferenceObservation struct {
}

type SourceImageReferenceParameters struct {

	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type VirtualMachineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`

	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type VirtualMachineParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalCapabilities []AdditionalCapabilitiesParameters `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`

	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AdminSSHKey []AdminSSHKeyParameters `json:"adminSshKey,omitempty" tf:"admin_ssh_key,omitempty"`

	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// +kubebuilder:validation:Optional
	AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty" tf:"allow_extension_operations,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilitySetID *string `json:"availabilitySetId,omitempty" tf:"availability_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	BootDiagnostics []BootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`

	// +kubebuilder:validation:Optional
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDataSecretRef *v1.SecretKeySelector `json:"customDataSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DedicatedHostID *string `json:"dedicatedHostId,omitempty" tf:"dedicated_host_id,omitempty"`

	// +kubebuilder:validation:Optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterfaceIds []*string `json:"networkInterfaceIds" tf:"network_interface_ids,omitempty"`

	// +kubebuilder:validation:Required
	OsDisk []OsDiskParameters `json:"osDisk" tf:"os_disk,omitempty"`

	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformFaultDomain *int64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisionVMAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`

	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Secret []SecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// +kubebuilder:validation:Required
	Size *string `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImageID *string `json:"sourceImageId,omitempty" tf:"source_image_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImageReference []SourceImageReferenceParameters `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualMachineScaleSetID *string `json:"virtualMachineScaleSetId,omitempty" tf:"virtual_machine_scale_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// VirtualMachineSpec defines the desired state of VirtualMachine
type VirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualMachineParameters `json:"forProvider"`
}

// VirtualMachineStatus defines the observed state of VirtualMachine.
type VirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachine is the Schema for the VirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineSpec   `json:"spec"`
	Status            VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachines
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	VirtualMachine_Kind             = "VirtualMachine"
	VirtualMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualMachine_Kind}.String()
	VirtualMachine_KindAPIVersion   = VirtualMachine_Kind + "." + CRDGroupVersion.String()
	VirtualMachine_GroupVersionKind = CRDGroupVersion.WithKind(VirtualMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
