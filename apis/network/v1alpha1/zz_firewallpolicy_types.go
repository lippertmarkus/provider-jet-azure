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

type DNSObservation struct {
}

type DNSParameters struct {

	// +kubebuilder:validation:Optional
	NetworkRuleFqdnEnabled *bool `json:"networkRuleFqdnEnabled,omitempty" tf:"network_rule_fqdn_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyEnabled *bool `json:"proxyEnabled,omitempty" tf:"proxy_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Servers []*string `json:"servers,omitempty" tf:"servers,omitempty"`
}

type FirewallPolicyIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type FirewallPolicyIdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UserAssignedIdentityIds []*string `json:"userAssignedIdentityIds,omitempty" tf:"user_assigned_identity_ids,omitempty"`
}

type FirewallPolicyObservation struct {
	ChildPolicies []*string `json:"childPolicies,omitempty" tf:"child_policies,omitempty"`

	Firewalls []*string `json:"firewalls,omitempty" tf:"firewalls,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RuleCollectionGroups []*string `json:"ruleCollectionGroups,omitempty" tf:"rule_collection_groups,omitempty"`
}

type FirewallPolicyParameters struct {

	// +kubebuilder:validation:Optional
	BasePolicyID *string `json:"basePolicyId,omitempty" tf:"base_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []FirewallPolicyIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	IntrusionDetection []IntrusionDetectionParameters `json:"intrusionDetection,omitempty" tf:"intrusion_detection,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPRanges []*string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	TLSCertificate []TLSCertificateParameters `json:"tlsCertificate,omitempty" tf:"tls_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThreatIntelligenceAllowlist []ThreatIntelligenceAllowlistParameters `json:"threatIntelligenceAllowlist,omitempty" tf:"threat_intelligence_allowlist,omitempty"`

	// +kubebuilder:validation:Optional
	ThreatIntelligenceMode *string `json:"threatIntelligenceMode,omitempty" tf:"threat_intelligence_mode,omitempty"`
}

type IntrusionDetectionObservation struct {
}

type IntrusionDetectionParameters struct {

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	SignatureOverrides []SignatureOverridesParameters `json:"signatureOverrides,omitempty" tf:"signature_overrides,omitempty"`

	// +kubebuilder:validation:Optional
	TrafficBypass []TrafficBypassParameters `json:"trafficBypass,omitempty" tf:"traffic_bypass,omitempty"`
}

type SignatureOverridesObservation struct {
}

type SignatureOverridesParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type TLSCertificateObservation struct {
}

type TLSCertificateParameters struct {

	// +kubebuilder:validation:Required
	KeyVaultSecretID *string `json:"keyVaultSecretId" tf:"key_vault_secret_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ThreatIntelligenceAllowlistObservation struct {
}

type ThreatIntelligenceAllowlistParameters struct {

	// +kubebuilder:validation:Optional
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
}

type TrafficBypassObservation struct {
}

type TrafficBypassParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationAddresses []*string `json:"destinationAddresses,omitempty" tf:"destination_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationIPGroups []*string `json:"destinationIpGroups,omitempty" tf:"destination_ip_groups,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationPorts []*string `json:"destinationPorts,omitempty" tf:"destination_ports,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`
}

// FirewallPolicySpec defines the desired state of FirewallPolicy
type FirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyParameters `json:"forProvider"`
}

// FirewallPolicyStatus defines the observed state of FirewallPolicy.
type FirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicy is the Schema for the FirewallPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicySpec   `json:"spec"`
	Status            FirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyList contains a list of FirewallPolicys
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicy_Kind             = "FirewallPolicy"
	FirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicy_Kind}.String()
	FirewallPolicy_KindAPIVersion   = FirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicy{}, &FirewallPolicyList{})
}
