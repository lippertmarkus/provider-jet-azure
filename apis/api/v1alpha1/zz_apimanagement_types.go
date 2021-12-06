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

type APIManagementObservation struct {
	DeveloperPortalURL *string `json:"developerPortalUrl,omitempty" tf:"developer_portal_url,omitempty"`

	GatewayRegionalURL *string `json:"gatewayRegionalUrl,omitempty" tf:"gateway_regional_url,omitempty"`

	GatewayURL *string `json:"gatewayUrl,omitempty" tf:"gateway_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ManagementAPIURL *string `json:"managementApiUrl,omitempty" tf:"management_api_url,omitempty"`

	PortalURL *string `json:"portalUrl,omitempty" tf:"portal_url,omitempty"`

	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`

	ScmURL *string `json:"scmUrl,omitempty" tf:"scm_url,omitempty"`
}

type APIManagementParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLocation []AdditionalLocationParameters `json:"additionalLocation,omitempty" tf:"additional_location,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate []CertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCertificateEnabled *bool `json:"clientCertificateEnabled,omitempty" tf:"client_certificate_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayDisabled *bool `json:"gatewayDisabled,omitempty" tf:"gateway_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	HostnameConfiguration []HostnameConfigurationParameters `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MinAPIVersion *string `json:"minApiVersion,omitempty" tf:"min_api_version,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationSenderEmail *string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email,omitempty"`

	// +kubebuilder:validation:Optional
	Policy []PolicyParameters `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	Protocols []ProtocolsParameters `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +kubebuilder:validation:Required
	PublisherEmail *string `json:"publisherEmail" tf:"publisher_email,omitempty"`

	// +kubebuilder:validation:Required
	PublisherName *string `json:"publisherName" tf:"publisher_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Security []SecurityParameters `json:"security,omitempty" tf:"security,omitempty"`

	// +kubebuilder:validation:Optional
	SignIn []SignInParameters `json:"signIn,omitempty" tf:"sign_in,omitempty"`

	// +kubebuilder:validation:Optional
	SignUp []SignUpParameters `json:"signUp,omitempty" tf:"sign_up,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TenantAccess []TenantAccessParameters `json:"tenantAccess,omitempty" tf:"tenant_access,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration []APIManagementVirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkType *string `json:"virtualNetworkType,omitempty" tf:"virtual_network_type,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type APIManagementVirtualNetworkConfigurationObservation struct {
}

type APIManagementVirtualNetworkConfigurationParameters struct {

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

type AdditionalLocationObservation struct {
	GatewayRegionalURL *string `json:"gatewayRegionalUrl,omitempty" tf:"gateway_regional_url,omitempty"`

	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`
}

type AdditionalLocationParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration []VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`
}

type CertificateObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CertificateParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	EncodedCertificateSecretRef v1.SecretKeySelector `json:"encodedCertificateSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StoreName *string `json:"storeName" tf:"store_name,omitempty"`
}

type DeveloperPortalObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type DeveloperPortalParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type HostnameConfigurationObservation struct {
}

type HostnameConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	DeveloperPortal []DeveloperPortalParameters `json:"developerPortal,omitempty" tf:"developer_portal,omitempty"`

	// +kubebuilder:validation:Optional
	Management []ManagementParameters `json:"management,omitempty" tf:"management,omitempty"`

	// +kubebuilder:validation:Optional
	Portal []PortalParameters `json:"portal,omitempty" tf:"portal,omitempty"`

	// +kubebuilder:validation:Optional
	Proxy []ProxyParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// +kubebuilder:validation:Optional
	Scm []ScmParameters `json:"scm,omitempty" tf:"scm,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ManagementObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ManagementParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type PolicyObservation struct {
}

type PolicyParameters struct {

	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type PortalObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type PortalParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type ProtocolsObservation struct {
}

type ProtocolsParameters struct {

	// +kubebuilder:validation:Optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`
}

type ProxyObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ProxyParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultSslBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding,omitempty"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type ScmObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ScmParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type SecurityObservation struct {
}

type SecurityParameters struct {

	// +kubebuilder:validation:Optional
	EnableBackendSsl30 *bool `json:"enableBackendSsl30,omitempty" tf:"enable_backend_ssl30,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBackendTls10 *bool `json:"enableBackendTls10,omitempty" tf:"enable_backend_tls10,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBackendTls11 *bool `json:"enableBackendTls11,omitempty" tf:"enable_backend_tls11,omitempty"`

	// +kubebuilder:validation:Optional
	EnableFrontendSsl30 *bool `json:"enableFrontendSsl30,omitempty" tf:"enable_frontend_ssl30,omitempty"`

	// +kubebuilder:validation:Optional
	EnableFrontendTls10 *bool `json:"enableFrontendTls10,omitempty" tf:"enable_frontend_tls10,omitempty"`

	// +kubebuilder:validation:Optional
	EnableFrontendTls11 *bool `json:"enableFrontendTls11,omitempty" tf:"enable_frontend_tls11,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTripleDesCiphers *bool `json:"enableTripleDesCiphers,omitempty" tf:"enable_triple_des_ciphers,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEcdheEcdsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEcdheEcdsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEcdheRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEcdheRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSRsaWithAes128CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes128CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha256_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSRsaWithAes128GcmSha256CiphersEnabled *bool `json:"tlsRsaWithAes128GcmSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_gcm_sha256_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSRsaWithAes256CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes256CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha256_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TripleDesCiphersEnabled *bool `json:"tripleDesCiphersEnabled,omitempty" tf:"triple_des_ciphers_enabled,omitempty"`
}

type SignInObservation struct {
}

type SignInParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type SignUpObservation struct {
}

type SignUpParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	TermsOfService []TermsOfServiceParameters `json:"termsOfService" tf:"terms_of_service,omitempty"`
}

type TenantAccessObservation struct {
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type TenantAccessParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type TermsOfServiceObservation struct {
}

type TermsOfServiceParameters struct {

	// +kubebuilder:validation:Required
	ConsentRequired *bool `json:"consentRequired" tf:"consent_required,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type VirtualNetworkConfigurationObservation struct {
}

type VirtualNetworkConfigurationParameters struct {

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// APIManagementSpec defines the desired state of APIManagement
type APIManagementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIManagementParameters `json:"forProvider"`
}

// APIManagementStatus defines the observed state of APIManagement.
type APIManagementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIManagementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIManagement is the Schema for the APIManagements API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type APIManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIManagementSpec   `json:"spec"`
	Status            APIManagementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIManagementList contains a list of APIManagements
type APIManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIManagement `json:"items"`
}

// Repository type metadata.
var (
	APIManagement_Kind             = "APIManagement"
	APIManagement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIManagement_Kind}.String()
	APIManagement_KindAPIVersion   = APIManagement_Kind + "." + CRDGroupVersion.String()
	APIManagement_GroupVersionKind = CRDGroupVersion.WithKind(APIManagement_Kind)
)

func init() {
	SchemeBuilder.Register(&APIManagement{}, &APIManagementList{})
}
