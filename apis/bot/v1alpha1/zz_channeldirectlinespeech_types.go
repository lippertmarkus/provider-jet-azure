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

type ChannelDirectLineSpeechObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ChannelDirectLineSpeechParameters struct {

	// +kubebuilder:validation:Required
	BotName *string `json:"botName" tf:"bot_name,omitempty"`

	// +kubebuilder:validation:Required
	CognitiveServiceAccessKeySecretRef v1.SecretKeySelector `json:"cognitiveServiceAccessKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	CognitiveServiceLocation *string `json:"cognitiveServiceLocation" tf:"cognitive_service_location,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSpeechModelID *string `json:"customSpeechModelId,omitempty" tf:"custom_speech_model_id,omitempty"`

	// +kubebuilder:validation:Optional
	CustomVoiceDeploymentID *string `json:"customVoiceDeploymentId,omitempty" tf:"custom_voice_deployment_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// ChannelDirectLineSpeechSpec defines the desired state of ChannelDirectLineSpeech
type ChannelDirectLineSpeechSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChannelDirectLineSpeechParameters `json:"forProvider"`
}

// ChannelDirectLineSpeechStatus defines the observed state of ChannelDirectLineSpeech.
type ChannelDirectLineSpeechStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChannelDirectLineSpeechObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelDirectLineSpeech is the Schema for the ChannelDirectLineSpeechs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ChannelDirectLineSpeech struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelDirectLineSpeechSpec   `json:"spec"`
	Status            ChannelDirectLineSpeechStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelDirectLineSpeechList contains a list of ChannelDirectLineSpeechs
type ChannelDirectLineSpeechList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChannelDirectLineSpeech `json:"items"`
}

// Repository type metadata.
var (
	ChannelDirectLineSpeech_Kind             = "ChannelDirectLineSpeech"
	ChannelDirectLineSpeech_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ChannelDirectLineSpeech_Kind}.String()
	ChannelDirectLineSpeech_KindAPIVersion   = ChannelDirectLineSpeech_Kind + "." + CRDGroupVersion.String()
	ChannelDirectLineSpeech_GroupVersionKind = CRDGroupVersion.WithKind(ChannelDirectLineSpeech_Kind)
)

func init() {
	SchemeBuilder.Register(&ChannelDirectLineSpeech{}, &ChannelDirectLineSpeechList{})
}
