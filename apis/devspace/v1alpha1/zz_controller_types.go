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

type ControllerObservation struct {
	DataPlaneFqdn *string `json:"dataPlaneFqdn,omitempty" tf:"data_plane_fqdn,omitempty"`

	HostSuffix *string `json:"hostSuffix,omitempty" tf:"host_suffix,omitempty"`
}

type ControllerParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetContainerHostCredentialsBase64SecretRef v1.SecretKeySelector `json:"targetContainerHostCredentialsBase64SecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	TargetContainerHostResourceID *string `json:"targetContainerHostResourceId" tf:"target_container_host_resource_id,omitempty"`
}

// ControllerSpec defines the desired state of Controller
type ControllerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ControllerParameters `json:"forProvider"`
}

// ControllerStatus defines the observed state of Controller.
type ControllerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ControllerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Controller is the Schema for the Controllers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Controller struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ControllerSpec   `json:"spec"`
	Status            ControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ControllerList contains a list of Controllers
type ControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Controller `json:"items"`
}

// Repository type metadata.
var (
	Controller_Kind             = "Controller"
	Controller_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Controller_Kind}.String()
	Controller_KindAPIVersion   = Controller_Kind + "." + CRDGroupVersion.String()
	Controller_GroupVersionKind = CRDGroupVersion.WithKind(Controller_Kind)
)

func init() {
	SchemeBuilder.Register(&Controller{}, &ControllerList{})
}
