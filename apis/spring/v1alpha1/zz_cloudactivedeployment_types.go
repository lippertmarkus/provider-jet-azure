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

type CloudActiveDeploymentObservation struct {
}

type CloudActiveDeploymentParameters struct {

	// +kubebuilder:validation:Required
	DeploymentName *string `json:"deploymentName" tf:"deployment_name,omitempty"`

	// +kubebuilder:validation:Required
	SpringCloudAppID *string `json:"springCloudAppId" tf:"spring_cloud_app_id,omitempty"`
}

// CloudActiveDeploymentSpec defines the desired state of CloudActiveDeployment
type CloudActiveDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudActiveDeploymentParameters `json:"forProvider"`
}

// CloudActiveDeploymentStatus defines the observed state of CloudActiveDeployment.
type CloudActiveDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudActiveDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudActiveDeployment is the Schema for the CloudActiveDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type CloudActiveDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudActiveDeploymentSpec   `json:"spec"`
	Status            CloudActiveDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudActiveDeploymentList contains a list of CloudActiveDeployments
type CloudActiveDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudActiveDeployment `json:"items"`
}

// Repository type metadata.
var (
	CloudActiveDeployment_Kind             = "CloudActiveDeployment"
	CloudActiveDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudActiveDeployment_Kind}.String()
	CloudActiveDeployment_KindAPIVersion   = CloudActiveDeployment_Kind + "." + CRDGroupVersion.String()
	CloudActiveDeployment_GroupVersionKind = CRDGroupVersion.WithKind(CloudActiveDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudActiveDeployment{}, &CloudActiveDeploymentList{})
}
