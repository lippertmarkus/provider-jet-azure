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

type LearningSynapseSparkIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type LearningSynapseSparkIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LearningSynapseSparkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LearningSynapseSparkParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []LearningSynapseSparkIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId" tf:"machine_learning_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SynapseSparkPoolID *string `json:"synapseSparkPoolId" tf:"synapse_spark_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LearningSynapseSparkSpec defines the desired state of LearningSynapseSpark
type LearningSynapseSparkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LearningSynapseSparkParameters `json:"forProvider"`
}

// LearningSynapseSparkStatus defines the observed state of LearningSynapseSpark.
type LearningSynapseSparkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LearningSynapseSparkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LearningSynapseSpark is the Schema for the LearningSynapseSparks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LearningSynapseSpark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LearningSynapseSparkSpec   `json:"spec"`
	Status            LearningSynapseSparkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LearningSynapseSparkList contains a list of LearningSynapseSparks
type LearningSynapseSparkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LearningSynapseSpark `json:"items"`
}

// Repository type metadata.
var (
	LearningSynapseSpark_Kind             = "LearningSynapseSpark"
	LearningSynapseSpark_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LearningSynapseSpark_Kind}.String()
	LearningSynapseSpark_KindAPIVersion   = LearningSynapseSpark_Kind + "." + CRDGroupVersion.String()
	LearningSynapseSpark_GroupVersionKind = CRDGroupVersion.WithKind(LearningSynapseSpark_Kind)
)

func init() {
	SchemeBuilder.Register(&LearningSynapseSpark{}, &LearningSynapseSparkList{})
}
