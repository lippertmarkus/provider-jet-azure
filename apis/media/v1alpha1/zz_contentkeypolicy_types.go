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

type ContentKeyPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ContentKeyPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	MediaServicesAccountName *string `json:"mediaServicesAccountName" tf:"media_services_account_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PolicyOption []PolicyOptionParameters `json:"policyOption" tf:"policy_option,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type FairplayConfigurationObservation struct {
}

type FairplayConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AskSecretRef *v1.SecretKeySelector `json:"askSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OfflineRentalConfiguration []OfflineRentalConfigurationParameters `json:"offlineRentalConfiguration,omitempty" tf:"offline_rental_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	PfxPasswordSecretRef *v1.SecretKeySelector `json:"pfxPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PfxSecretRef *v1.SecretKeySelector `json:"pfxSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RentalAndLeaseKeyType *string `json:"rentalAndLeaseKeyType,omitempty" tf:"rental_and_lease_key_type,omitempty"`

	// +kubebuilder:validation:Optional
	RentalDurationSeconds *int64 `json:"rentalDurationSeconds,omitempty" tf:"rental_duration_seconds,omitempty"`
}

type OfflineRentalConfigurationObservation struct {
}

type OfflineRentalConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	PlaybackDurationSeconds *int64 `json:"playbackDurationSeconds,omitempty" tf:"playback_duration_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	StorageDurationSeconds *int64 `json:"storageDurationSeconds,omitempty" tf:"storage_duration_seconds,omitempty"`
}

type PlayRightObservation struct {
}

type PlayRightParameters struct {

	// +kubebuilder:validation:Optional
	AgcAndColorStripeRestriction *int64 `json:"agcAndColorStripeRestriction,omitempty" tf:"agc_and_color_stripe_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	AllowPassingVideoContentToUnknownOutput *string `json:"allowPassingVideoContentToUnknownOutput,omitempty" tf:"allow_passing_video_content_to_unknown_output,omitempty"`

	// +kubebuilder:validation:Optional
	AnalogVideoOpl *int64 `json:"analogVideoOpl,omitempty" tf:"analog_video_opl,omitempty"`

	// +kubebuilder:validation:Optional
	CompressedDigitalAudioOpl *int64 `json:"compressedDigitalAudioOpl,omitempty" tf:"compressed_digital_audio_opl,omitempty"`

	// +kubebuilder:validation:Optional
	DigitalVideoOnlyContentRestriction *bool `json:"digitalVideoOnlyContentRestriction,omitempty" tf:"digital_video_only_content_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	FirstPlayExpiration *string `json:"firstPlayExpiration,omitempty" tf:"first_play_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	ImageConstraintForAnalogComponentVideoRestriction *bool `json:"imageConstraintForAnalogComponentVideoRestriction,omitempty" tf:"image_constraint_for_analog_component_video_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	ImageConstraintForAnalogComputerMonitorRestriction *bool `json:"imageConstraintForAnalogComputerMonitorRestriction,omitempty" tf:"image_constraint_for_analog_computer_monitor_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	ScmsRestriction *int64 `json:"scmsRestriction,omitempty" tf:"scms_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	UncompressedDigitalAudioOpl *int64 `json:"uncompressedDigitalAudioOpl,omitempty" tf:"uncompressed_digital_audio_opl,omitempty"`

	// +kubebuilder:validation:Optional
	UncompressedDigitalVideoOpl *int64 `json:"uncompressedDigitalVideoOpl,omitempty" tf:"uncompressed_digital_video_opl,omitempty"`
}

type PlayreadyConfigurationLicenseObservation struct {
}

type PlayreadyConfigurationLicenseParameters struct {

	// +kubebuilder:validation:Optional
	AllowTestDevices *bool `json:"allowTestDevices,omitempty" tf:"allow_test_devices,omitempty"`

	// +kubebuilder:validation:Optional
	BeginDate *string `json:"beginDate,omitempty" tf:"begin_date,omitempty"`

	// +kubebuilder:validation:Optional
	ContentKeyLocationFromHeaderEnabled *bool `json:"contentKeyLocationFromHeaderEnabled,omitempty" tf:"content_key_location_from_header_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ContentKeyLocationFromKeyID *string `json:"contentKeyLocationFromKeyId,omitempty" tf:"content_key_location_from_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// +kubebuilder:validation:Optional
	GracePeriodSecretRef *v1.SecretKeySelector `json:"gracePeriodSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// +kubebuilder:validation:Optional
	PlayRight []PlayRightParameters `json:"playRight,omitempty" tf:"play_right,omitempty"`

	// +kubebuilder:validation:Optional
	RelativeBeginDate *string `json:"relativeBeginDate,omitempty" tf:"relative_begin_date,omitempty"`

	// +kubebuilder:validation:Optional
	RelativeExpirationDate *string `json:"relativeExpirationDate,omitempty" tf:"relative_expiration_date,omitempty"`
}

type PolicyOptionObservation struct {
}

type PolicyOptionParameters struct {

	// +kubebuilder:validation:Optional
	ClearKeyConfigurationEnabled *bool `json:"clearKeyConfigurationEnabled,omitempty" tf:"clear_key_configuration_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	FairplayConfiguration []FairplayConfigurationParameters `json:"fairplayConfiguration,omitempty" tf:"fairplay_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OpenRestrictionEnabled *bool `json:"openRestrictionEnabled,omitempty" tf:"open_restriction_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PlayreadyConfigurationLicense []PlayreadyConfigurationLicenseParameters `json:"playreadyConfigurationLicense,omitempty" tf:"playready_configuration_license,omitempty"`

	// +kubebuilder:validation:Optional
	TokenRestriction []TokenRestrictionParameters `json:"tokenRestriction,omitempty" tf:"token_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	WidevineConfigurationTemplate *string `json:"widevineConfigurationTemplate,omitempty" tf:"widevine_configuration_template,omitempty"`
}

type RequiredClaimObservation struct {
}

type RequiredClaimParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TokenRestrictionObservation struct {
}

type TokenRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// +kubebuilder:validation:Optional
	OpenIDConnectDiscoveryDocument *string `json:"openIdConnectDiscoveryDocument,omitempty" tf:"open_id_connect_discovery_document,omitempty"`

	// +kubebuilder:validation:Optional
	PrimaryRsaTokenKeyExponentSecretRef *v1.SecretKeySelector `json:"primaryRsaTokenKeyExponentSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrimaryRsaTokenKeyModulusSecretRef *v1.SecretKeySelector `json:"primaryRsaTokenKeyModulusSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrimarySymmetricTokenKeySecretRef *v1.SecretKeySelector `json:"primarySymmetricTokenKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrimaryX509TokenKeyRawSecretRef *v1.SecretKeySelector `json:"primaryX509TokenKeyRawSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RequiredClaim []RequiredClaimParameters `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// ContentKeyPolicySpec defines the desired state of ContentKeyPolicy
type ContentKeyPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContentKeyPolicyParameters `json:"forProvider"`
}

// ContentKeyPolicyStatus defines the observed state of ContentKeyPolicy.
type ContentKeyPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContentKeyPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContentKeyPolicy is the Schema for the ContentKeyPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ContentKeyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContentKeyPolicySpec   `json:"spec"`
	Status            ContentKeyPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentKeyPolicyList contains a list of ContentKeyPolicys
type ContentKeyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentKeyPolicy `json:"items"`
}

// Repository type metadata.
var (
	ContentKeyPolicy_Kind             = "ContentKeyPolicy"
	ContentKeyPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContentKeyPolicy_Kind}.String()
	ContentKeyPolicy_KindAPIVersion   = ContentKeyPolicy_Kind + "." + CRDGroupVersion.String()
	ContentKeyPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ContentKeyPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ContentKeyPolicy{}, &ContentKeyPolicyList{})
}
