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

type CaptureObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CaptureParameters struct {

	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumBytesPerPacket *int64 `json:"maximumBytesPerPacket,omitempty" tf:"maximum_bytes_per_packet,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumBytesPerSession *int64 `json:"maximumBytesPerSession,omitempty" tf:"maximum_bytes_per_session,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumCaptureDuration *int64 `json:"maximumCaptureDuration,omitempty" tf:"maximum_capture_duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkWatcherName *string `json:"networkWatcherName" tf:"network_watcher_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageLocation []StorageLocationParameters `json:"storageLocation" tf:"storage_location,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceID *string `json:"targetResourceId" tf:"target_resource_id,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	LocalIPAddress *string `json:"localIpAddress,omitempty" tf:"local_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	LocalPort *string `json:"localPort,omitempty" tf:"local_port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteIPAddress *string `json:"remoteIpAddress,omitempty" tf:"remote_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	RemotePort *string `json:"remotePort,omitempty" tf:"remote_port,omitempty"`
}

type StorageLocationObservation struct {
	StoragePath *string `json:"storagePath,omitempty" tf:"storage_path,omitempty"`
}

type StorageLocationParameters struct {

	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

// CaptureSpec defines the desired state of Capture
type CaptureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CaptureParameters `json:"forProvider"`
}

// CaptureStatus defines the observed state of Capture.
type CaptureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CaptureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Capture is the Schema for the Captures API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Capture struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CaptureSpec   `json:"spec"`
	Status            CaptureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaptureList contains a list of Captures
type CaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Capture `json:"items"`
}

// Repository type metadata.
var (
	Capture_Kind             = "Capture"
	Capture_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Capture_Kind}.String()
	Capture_KindAPIVersion   = Capture_Kind + "." + CRDGroupVersion.String()
	Capture_GroupVersionKind = CRDGroupVersion.WithKind(Capture_Kind)
)

func init() {
	SchemeBuilder.Register(&Capture{}, &CaptureList{})
}
