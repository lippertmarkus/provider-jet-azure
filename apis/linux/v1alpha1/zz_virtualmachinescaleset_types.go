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

type AutomaticInstanceRepairObservation struct {
}

type AutomaticInstanceRepairParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	GracePeriod *string `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`
}

type AutomaticOsUpgradePolicyObservation struct {
}

type AutomaticOsUpgradePolicyParameters struct {

	// +kubebuilder:validation:Required
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback" tf:"disable_automatic_rollback,omitempty"`

	// +kubebuilder:validation:Required
	EnableAutomaticOsUpgrade *bool `json:"enableAutomaticOsUpgrade" tf:"enable_automatic_os_upgrade,omitempty"`
}

type DataDiskObservation struct {
}

type DataDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// +kubebuilder:validation:Optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskIopsReadWrite *int64 `json:"diskIopsReadWrite,omitempty" tf:"disk_iops_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	DiskMbpsReadWrite *int64 `json:"diskMbpsReadWrite,omitempty" tf:"disk_mbps_read_write,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *int64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	Lun *int64 `json:"lun" tf:"lun,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type ExtensionObservation struct {
}

type ExtensionParameters struct {

	// +kubebuilder:validation:Optional
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`

	// +kubebuilder:validation:Optional
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProtectedSettingsSecretRef *v1.SecretKeySelector `json:"protectedSettingsSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProvisionAfterExtensions []*string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	TypeHandlerVersion *string `json:"typeHandlerVersion" tf:"type_handler_version,omitempty"`
}

type IPConfigurationObservation struct {
}

type IPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationGatewayBackendAddressPoolIds []*string `json:"applicationGatewayBackendAddressPoolIds,omitempty" tf:"application_gateway_backend_address_pool_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIds []*string `json:"applicationSecurityGroupIds,omitempty" tf:"application_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerBackendAddressPoolIds []*string `json:"loadBalancerBackendAddressPoolIds,omitempty" tf:"load_balancer_backend_address_pool_ids,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerInboundNatRulesIds []*string `json:"loadBalancerInboundNatRulesIds,omitempty" tf:"load_balancer_inbound_nat_rules_ids,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPAddress []PublicIPAddressParameters `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type IPTagObservation struct {
}

type IPTagParameters struct {

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`

	// +kubebuilder:validation:Optional
	EnableIPForwarding *bool `json:"enableIpForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`

	// +kubebuilder:validation:Required
	IPConfiguration []IPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`
}

type OsDiskDiffDiskSettingsObservation struct {
}

type OsDiskDiffDiskSettingsParameters struct {

	// +kubebuilder:validation:Required
	Option *string `json:"option" tf:"option,omitempty"`
}

type PublicIPAddressObservation struct {
}

type PublicIPAddressParameters struct {

	// +kubebuilder:validation:Optional
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// +kubebuilder:validation:Optional
	IPTag []IPTagParameters `json:"ipTag,omitempty" tf:"ip_tag,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`
}

type RollingUpgradePolicyObservation struct {
}

type RollingUpgradePolicyParameters struct {

	// +kubebuilder:validation:Required
	MaxBatchInstancePercent *int64 `json:"maxBatchInstancePercent" tf:"max_batch_instance_percent,omitempty"`

	// +kubebuilder:validation:Required
	MaxUnhealthyInstancePercent *int64 `json:"maxUnhealthyInstancePercent" tf:"max_unhealthy_instance_percent,omitempty"`

	// +kubebuilder:validation:Required
	MaxUnhealthyUpgradedInstancePercent *int64 `json:"maxUnhealthyUpgradedInstancePercent" tf:"max_unhealthy_upgraded_instance_percent,omitempty"`

	// +kubebuilder:validation:Required
	PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches" tf:"pause_time_between_batches,omitempty"`
}

type SecretCertificateObservation struct {
}

type SecretCertificateParameters struct {

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type TerminateNotificationObservation struct {
}

type TerminateNotificationParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type VirtualMachineScaleSetAdditionalCapabilitiesObservation struct {
}

type VirtualMachineScaleSetAdditionalCapabilitiesParameters struct {

	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type VirtualMachineScaleSetAdminSSHKeyObservation struct {
}

type VirtualMachineScaleSetAdminSSHKeyParameters struct {

	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type VirtualMachineScaleSetBootDiagnosticsObservation struct {
}

type VirtualMachineScaleSetBootDiagnosticsParameters struct {

	// +kubebuilder:validation:Optional
	StorageAccountURI *string `json:"storageAccountUri,omitempty" tf:"storage_account_uri,omitempty"`
}

type VirtualMachineScaleSetIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`
}

type VirtualMachineScaleSetIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VirtualMachineScaleSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type VirtualMachineScaleSetOsDiskObservation struct {
}

type VirtualMachineScaleSetOsDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// +kubebuilder:validation:Optional
	DiffDiskSettings []OsDiskDiffDiskSettingsParameters `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineScaleSetParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalCapabilities []VirtualMachineScaleSetAdditionalCapabilitiesParameters `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`

	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AdminSSHKey []VirtualMachineScaleSetAdminSSHKeyParameters `json:"adminSshKey,omitempty" tf:"admin_ssh_key,omitempty"`

	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticInstanceRepair []AutomaticInstanceRepairParameters `json:"automaticInstanceRepair,omitempty" tf:"automatic_instance_repair,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticOsUpgradePolicy []AutomaticOsUpgradePolicyParameters `json:"automaticOsUpgradePolicy,omitempty" tf:"automatic_os_upgrade_policy,omitempty"`

	// +kubebuilder:validation:Optional
	BootDiagnostics []VirtualMachineScaleSetBootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`

	// +kubebuilder:validation:Optional
	ComputerNamePrefix *string `json:"computerNamePrefix,omitempty" tf:"computer_name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDataSecretRef *v1.SecretKeySelector `json:"customDataSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DataDisk []DataDiskParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// +kubebuilder:validation:Optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	DoNotRunExtensionsOnOverprovisionedMachines *bool `json:"doNotRunExtensionsOnOverprovisionedMachines,omitempty" tf:"do_not_run_extensions_on_overprovisioned_machines,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Extension []ExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// +kubebuilder:validation:Optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget,omitempty"`

	// +kubebuilder:validation:Optional
	HealthProbeID *string `json:"healthProbeId,omitempty" tf:"health_probe_id,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []VirtualMachineScaleSetIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Instances *int64 `json:"instances" tf:"instances,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Required
	OsDisk []VirtualMachineScaleSetOsDiskParameters `json:"osDisk" tf:"os_disk,omitempty"`

	// +kubebuilder:validation:Optional
	Overprovision *bool `json:"overprovision,omitempty" tf:"overprovision,omitempty"`

	// +kubebuilder:validation:Optional
	Plan []VirtualMachineScaleSetPlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformFaultDomainCount *int64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisionVMAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`

	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RollingUpgradePolicy []RollingUpgradePolicyParameters `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleInPolicy *string `json:"scaleInPolicy,omitempty" tf:"scale_in_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Secret []VirtualMachineScaleSetSecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty" tf:"single_placement_group,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImageID *string `json:"sourceImageId,omitempty" tf:"source_image_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImageReference []VirtualMachineScaleSetSourceImageReferenceParameters `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TerminateNotification []TerminateNotificationParameters `json:"terminateNotification,omitempty" tf:"terminate_notification,omitempty"`

	// +kubebuilder:validation:Optional
	UpgradeMode *string `json:"upgradeMode,omitempty" tf:"upgrade_mode,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneBalance *bool `json:"zoneBalance,omitempty" tf:"zone_balance,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type VirtualMachineScaleSetPlanObservation struct {
}

type VirtualMachineScaleSetPlanParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

type VirtualMachineScaleSetSecretObservation struct {
}

type VirtualMachineScaleSetSecretParameters struct {

	// +kubebuilder:validation:Required
	Certificate []SecretCertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`
}

type VirtualMachineScaleSetSourceImageReferenceObservation struct {
}

type VirtualMachineScaleSetSourceImageReferenceParameters struct {

	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// VirtualMachineScaleSetSpec defines the desired state of VirtualMachineScaleSet
type VirtualMachineScaleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualMachineScaleSetParameters `json:"forProvider"`
}

// VirtualMachineScaleSetStatus defines the observed state of VirtualMachineScaleSet.
type VirtualMachineScaleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualMachineScaleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineScaleSet is the Schema for the VirtualMachineScaleSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineScaleSetSpec   `json:"spec"`
	Status            VirtualMachineScaleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineScaleSetList contains a list of VirtualMachineScaleSets
type VirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineScaleSet `json:"items"`
}

// Repository type metadata.
var (
	VirtualMachineScaleSet_Kind             = "VirtualMachineScaleSet"
	VirtualMachineScaleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualMachineScaleSet_Kind}.String()
	VirtualMachineScaleSet_KindAPIVersion   = VirtualMachineScaleSet_Kind + "." + CRDGroupVersion.String()
	VirtualMachineScaleSet_GroupVersionKind = CRDGroupVersion.WithKind(VirtualMachineScaleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachineScaleSet{}, &VirtualMachineScaleSetList{})
}
