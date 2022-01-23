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

type DedicatedHostObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DedicatedHostParameters struct {

	// +kubebuilder:validation:Optional
	AutoReplaceOnFailure *bool `json:"autoReplaceOnFailure,omitempty" tf:"auto_replace_on_failure,omitempty"`

	// +kubebuilder:validation:Required
	DedicatedHostGroupID *string `json:"dedicatedHostGroupId" tf:"dedicated_host_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PlatformFaultDomain *int64 `json:"platformFaultDomain" tf:"platform_fault_domain,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DedicatedHostSpec defines the desired state of DedicatedHost
type DedicatedHostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedHostParameters `json:"forProvider"`
}

// DedicatedHostStatus defines the observed state of DedicatedHost.
type DedicatedHostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHost is the Schema for the DedicatedHosts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DedicatedHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHostSpec   `json:"spec"`
	Status            DedicatedHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHostList contains a list of DedicatedHosts
type DedicatedHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedHost `json:"items"`
}

// Repository type metadata.
var (
	DedicatedHost_Kind             = "DedicatedHost"
	DedicatedHost_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedHost_Kind}.String()
	DedicatedHost_KindAPIVersion   = DedicatedHost_Kind + "." + CRDGroupVersion.String()
	DedicatedHost_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedHost_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedHost{}, &DedicatedHostList{})
}
