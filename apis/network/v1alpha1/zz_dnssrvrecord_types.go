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

type DNSSRVRecordObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSSRVRecordParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Record []DNSSRVRecordRecordParameters `json:"record" tf:"record,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ZoneName *string `json:"zoneName" tf:"zone_name,omitempty"`
}

type DNSSRVRecordRecordObservation struct {
}

type DNSSRVRecordRecordParameters struct {

	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// +kubebuilder:validation:Required
	Weight *int64 `json:"weight" tf:"weight,omitempty"`
}

// DNSSRVRecordSpec defines the desired state of DNSSRVRecord
type DNSSRVRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSSRVRecordParameters `json:"forProvider"`
}

// DNSSRVRecordStatus defines the observed state of DNSSRVRecord.
type DNSSRVRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSSRVRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSSRVRecord is the Schema for the DNSSRVRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DNSSRVRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DNSSRVRecordSpec   `json:"spec"`
	Status            DNSSRVRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSSRVRecordList contains a list of DNSSRVRecords
type DNSSRVRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSSRVRecord `json:"items"`
}

// Repository type metadata.
var (
	DNSSRVRecord_Kind             = "DNSSRVRecord"
	DNSSRVRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSSRVRecord_Kind}.String()
	DNSSRVRecord_KindAPIVersion   = DNSSRVRecord_Kind + "." + CRDGroupVersion.String()
	DNSSRVRecord_GroupVersionKind = CRDGroupVersion.WithKind(DNSSRVRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSSRVRecord{}, &DNSSRVRecordList{})
}
