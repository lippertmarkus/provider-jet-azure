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

type AnalyticsOutputEventhubObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AnalyticsOutputEventhubParameters struct {

	// +kubebuilder:validation:Required
	EventhubName *string `json:"eventhubName" tf:"eventhub_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Serialization []AnalyticsOutputEventhubSerializationParameters `json:"serialization" tf:"serialization,omitempty"`

	// +kubebuilder:validation:Required
	ServicebusNamespace *string `json:"servicebusNamespace" tf:"servicebus_namespace,omitempty"`

	// +kubebuilder:validation:Required
	SharedAccessPolicyKeySecretRef v1.SecretKeySelector `json:"sharedAccessPolicyKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName" tf:"shared_access_policy_name,omitempty"`

	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`
}

type AnalyticsOutputEventhubSerializationObservation struct {
}

type AnalyticsOutputEventhubSerializationParameters struct {

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// AnalyticsOutputEventhubSpec defines the desired state of AnalyticsOutputEventhub
type AnalyticsOutputEventhubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsOutputEventhubParameters `json:"forProvider"`
}

// AnalyticsOutputEventhubStatus defines the observed state of AnalyticsOutputEventhub.
type AnalyticsOutputEventhubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsOutputEventhubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsOutputEventhub is the Schema for the AnalyticsOutputEventhubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AnalyticsOutputEventhub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsOutputEventhubSpec   `json:"spec"`
	Status            AnalyticsOutputEventhubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsOutputEventhubList contains a list of AnalyticsOutputEventhubs
type AnalyticsOutputEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsOutputEventhub `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsOutputEventhub_Kind             = "AnalyticsOutputEventhub"
	AnalyticsOutputEventhub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsOutputEventhub_Kind}.String()
	AnalyticsOutputEventhub_KindAPIVersion   = AnalyticsOutputEventhub_Kind + "." + CRDGroupVersion.String()
	AnalyticsOutputEventhub_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsOutputEventhub_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsOutputEventhub{}, &AnalyticsOutputEventhubList{})
}
