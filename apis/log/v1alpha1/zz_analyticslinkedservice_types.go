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

type AnalyticsLinkedServiceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AnalyticsLinkedServiceParameters struct {

	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Optional
	ReadAccessID *string `json:"readAccessId,omitempty" tf:"read_access_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// +kubebuilder:validation:Optional
	WorkspaceName *string `json:"workspaceName,omitempty" tf:"workspace_name,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAccessID *string `json:"writeAccessId,omitempty" tf:"write_access_id,omitempty"`
}

// AnalyticsLinkedServiceSpec defines the desired state of AnalyticsLinkedService
type AnalyticsLinkedServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsLinkedServiceParameters `json:"forProvider"`
}

// AnalyticsLinkedServiceStatus defines the observed state of AnalyticsLinkedService.
type AnalyticsLinkedServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsLinkedServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsLinkedService is the Schema for the AnalyticsLinkedServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AnalyticsLinkedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsLinkedServiceSpec   `json:"spec"`
	Status            AnalyticsLinkedServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsLinkedServiceList contains a list of AnalyticsLinkedServices
type AnalyticsLinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsLinkedService `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsLinkedService_Kind             = "AnalyticsLinkedService"
	AnalyticsLinkedService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsLinkedService_Kind}.String()
	AnalyticsLinkedService_KindAPIVersion   = AnalyticsLinkedService_Kind + "." + CRDGroupVersion.String()
	AnalyticsLinkedService_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsLinkedService_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsLinkedService{}, &AnalyticsLinkedServiceList{})
}
