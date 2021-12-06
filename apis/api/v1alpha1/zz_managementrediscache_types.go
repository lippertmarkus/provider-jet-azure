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

type ManagementRedisCacheObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagementRedisCacheParameters struct {

	// +kubebuilder:validation:Required
	APIManagementID *string `json:"apiManagementId" tf:"api_management_id,omitempty"`

	// +kubebuilder:validation:Optional
	CacheLocation *string `json:"cacheLocation,omitempty" tf:"cache_location,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`
}

// ManagementRedisCacheSpec defines the desired state of ManagementRedisCache
type ManagementRedisCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementRedisCacheParameters `json:"forProvider"`
}

// ManagementRedisCacheStatus defines the observed state of ManagementRedisCache.
type ManagementRedisCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementRedisCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementRedisCache is the Schema for the ManagementRedisCaches API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagementRedisCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementRedisCacheSpec   `json:"spec"`
	Status            ManagementRedisCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementRedisCacheList contains a list of ManagementRedisCaches
type ManagementRedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementRedisCache `json:"items"`
}

// Repository type metadata.
var (
	ManagementRedisCache_Kind             = "ManagementRedisCache"
	ManagementRedisCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementRedisCache_Kind}.String()
	ManagementRedisCache_KindAPIVersion   = ManagementRedisCache_Kind + "." + CRDGroupVersion.String()
	ManagementRedisCache_GroupVersionKind = CRDGroupVersion.WithKind(ManagementRedisCache_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementRedisCache{}, &ManagementRedisCacheList{})
}
