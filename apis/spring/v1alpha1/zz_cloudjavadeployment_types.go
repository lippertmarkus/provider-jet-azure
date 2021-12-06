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

type CloudJavaDeploymentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloudJavaDeploymentParameters struct {

	// +kubebuilder:validation:Optional
	CPU *int64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	JvmOptions *string `json:"jvmOptions,omitempty" tf:"jvm_options,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryInGb *int64 `json:"memoryInGb,omitempty" tf:"memory_in_gb,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`

	// +kubebuilder:validation:Required
	SpringCloudAppID *string `json:"springCloudAppId" tf:"spring_cloud_app_id,omitempty"`
}

// CloudJavaDeploymentSpec defines the desired state of CloudJavaDeployment
type CloudJavaDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudJavaDeploymentParameters `json:"forProvider"`
}

// CloudJavaDeploymentStatus defines the observed state of CloudJavaDeployment.
type CloudJavaDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudJavaDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudJavaDeployment is the Schema for the CloudJavaDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type CloudJavaDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudJavaDeploymentSpec   `json:"spec"`
	Status            CloudJavaDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudJavaDeploymentList contains a list of CloudJavaDeployments
type CloudJavaDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudJavaDeployment `json:"items"`
}

// Repository type metadata.
var (
	CloudJavaDeployment_Kind             = "CloudJavaDeployment"
	CloudJavaDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudJavaDeployment_Kind}.String()
	CloudJavaDeployment_KindAPIVersion   = CloudJavaDeployment_Kind + "." + CRDGroupVersion.String()
	CloudJavaDeployment_GroupVersionKind = CRDGroupVersion.WithKind(CloudJavaDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudJavaDeployment{}, &CloudJavaDeploymentList{})
}
