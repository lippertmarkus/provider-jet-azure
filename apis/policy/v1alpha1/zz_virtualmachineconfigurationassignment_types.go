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

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AssignmentType *string `json:"assignmentType,omitempty" tf:"assignment_type,omitempty"`

	// +kubebuilder:validation:Optional
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// +kubebuilder:validation:Optional
	ContentURI *string `json:"contentUri,omitempty" tf:"content_uri,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ParameterObservation struct {
}

type ParameterParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type VirtualMachineConfigurationAssignmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualMachineConfigurationAssignmentParameters struct {

	// +kubebuilder:validation:Required
	Configuration []ConfigurationParameters `json:"configuration" tf:"configuration,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	VirtualMachineID *string `json:"virtualMachineId" tf:"virtual_machine_id,omitempty"`
}

// VirtualMachineConfigurationAssignmentSpec defines the desired state of VirtualMachineConfigurationAssignment
type VirtualMachineConfigurationAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualMachineConfigurationAssignmentParameters `json:"forProvider"`
}

// VirtualMachineConfigurationAssignmentStatus defines the observed state of VirtualMachineConfigurationAssignment.
type VirtualMachineConfigurationAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualMachineConfigurationAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineConfigurationAssignment is the Schema for the VirtualMachineConfigurationAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VirtualMachineConfigurationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineConfigurationAssignmentSpec   `json:"spec"`
	Status            VirtualMachineConfigurationAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineConfigurationAssignmentList contains a list of VirtualMachineConfigurationAssignments
type VirtualMachineConfigurationAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineConfigurationAssignment `json:"items"`
}

// Repository type metadata.
var (
	VirtualMachineConfigurationAssignment_Kind             = "VirtualMachineConfigurationAssignment"
	VirtualMachineConfigurationAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualMachineConfigurationAssignment_Kind}.String()
	VirtualMachineConfigurationAssignment_KindAPIVersion   = VirtualMachineConfigurationAssignment_Kind + "." + CRDGroupVersion.String()
	VirtualMachineConfigurationAssignment_GroupVersionKind = CRDGroupVersion.WithKind(VirtualMachineConfigurationAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachineConfigurationAssignment{}, &VirtualMachineConfigurationAssignmentList{})
}
