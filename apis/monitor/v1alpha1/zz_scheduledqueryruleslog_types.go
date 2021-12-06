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

type CriteriaDimensionObservation struct {
}

type CriteriaDimensionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ScheduledQueryRulesLogCriteriaObservation struct {
}

type ScheduledQueryRulesLogCriteriaParameters struct {

	// +kubebuilder:validation:Required
	Dimension []CriteriaDimensionParameters `json:"dimension" tf:"dimension,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`
}

type ScheduledQueryRulesLogObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScheduledQueryRulesLogParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// +kubebuilder:validation:Required
	Criteria []ScheduledQueryRulesLogCriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Required
	DataSourceID *string `json:"dataSourceId" tf:"data_source_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ScheduledQueryRulesLogSpec defines the desired state of ScheduledQueryRulesLog
type ScheduledQueryRulesLogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduledQueryRulesLogParameters `json:"forProvider"`
}

// ScheduledQueryRulesLogStatus defines the observed state of ScheduledQueryRulesLog.
type ScheduledQueryRulesLogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduledQueryRulesLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledQueryRulesLog is the Schema for the ScheduledQueryRulesLogs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ScheduledQueryRulesLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduledQueryRulesLogSpec   `json:"spec"`
	Status            ScheduledQueryRulesLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledQueryRulesLogList contains a list of ScheduledQueryRulesLogs
type ScheduledQueryRulesLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledQueryRulesLog `json:"items"`
}

// Repository type metadata.
var (
	ScheduledQueryRulesLog_Kind             = "ScheduledQueryRulesLog"
	ScheduledQueryRulesLog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScheduledQueryRulesLog_Kind}.String()
	ScheduledQueryRulesLog_KindAPIVersion   = ScheduledQueryRulesLog_Kind + "." + CRDGroupVersion.String()
	ScheduledQueryRulesLog_GroupVersionKind = CRDGroupVersion.WithKind(ScheduledQueryRulesLog_Kind)
)

func init() {
	SchemeBuilder.Register(&ScheduledQueryRulesLog{}, &ScheduledQueryRulesLogList{})
}
