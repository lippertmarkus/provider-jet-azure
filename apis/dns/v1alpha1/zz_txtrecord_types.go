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

type TxtRecordObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TxtRecordParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Record []TxtRecordRecordParameters `json:"record" tf:"record,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ZoneName *string `json:"zoneName" tf:"zone_name,omitempty"`
}

type TxtRecordRecordObservation struct {
}

type TxtRecordRecordParameters struct {

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// TxtRecordSpec defines the desired state of TxtRecord
type TxtRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TxtRecordParameters `json:"forProvider"`
}

// TxtRecordStatus defines the observed state of TxtRecord.
type TxtRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TxtRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TxtRecord is the Schema for the TxtRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TxtRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TxtRecordSpec   `json:"spec"`
	Status            TxtRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TxtRecordList contains a list of TxtRecords
type TxtRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TxtRecord `json:"items"`
}

// Repository type metadata.
var (
	TxtRecord_Kind             = "TxtRecord"
	TxtRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TxtRecord_Kind}.String()
	TxtRecord_KindAPIVersion   = TxtRecord_Kind + "." + CRDGroupVersion.String()
	TxtRecord_GroupVersionKind = CRDGroupVersion.WithKind(TxtRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&TxtRecord{}, &TxtRecordList{})
}
