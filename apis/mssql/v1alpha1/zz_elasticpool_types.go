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

type ElasticpoolObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ElasticpoolParameters struct {

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PerDatabaseSettings []PerDatabaseSettingsParameters `json:"perDatabaseSettings" tf:"per_database_settings,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type PerDatabaseSettingsObservation struct {
}

type PerDatabaseSettingsParameters struct {

	// +kubebuilder:validation:Required
	MaxCapacity *float64 `json:"maxCapacity" tf:"max_capacity,omitempty"`

	// +kubebuilder:validation:Required
	MinCapacity *float64 `json:"minCapacity" tf:"min_capacity,omitempty"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity *int64 `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

// ElasticpoolSpec defines the desired state of Elasticpool
type ElasticpoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ElasticpoolParameters `json:"forProvider"`
}

// ElasticpoolStatus defines the observed state of Elasticpool.
type ElasticpoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ElasticpoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Elasticpool is the Schema for the Elasticpools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Elasticpool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticpoolSpec   `json:"spec"`
	Status            ElasticpoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticpoolList contains a list of Elasticpools
type ElasticpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Elasticpool `json:"items"`
}

// Repository type metadata.
var (
	Elasticpool_Kind             = "Elasticpool"
	Elasticpool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Elasticpool_Kind}.String()
	Elasticpool_KindAPIVersion   = Elasticpool_Kind + "." + CRDGroupVersion.String()
	Elasticpool_GroupVersionKind = CRDGroupVersion.WithKind(Elasticpool_Kind)
)

func init() {
	SchemeBuilder.Register(&Elasticpool{}, &ElasticpoolList{})
}
