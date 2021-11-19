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

type FactoryLinkedServiceKustoObservation struct {
}

type FactoryLinkedServiceKustoParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryID *string `json:"dataFactoryId" tf:"data_factory_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	KustoDatabaseName *string `json:"kustoDatabaseName" tf:"kusto_database_name,omitempty"`

	// +kubebuilder:validation:Required
	KustoEndpoint *string `json:"kustoEndpoint" tf:"kusto_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServicePrincipalKeySecretRef *v1.SecretKeySelector `json:"servicePrincipalKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// +kubebuilder:validation:Optional
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

// FactoryLinkedServiceKustoSpec defines the desired state of FactoryLinkedServiceKusto
type FactoryLinkedServiceKustoSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceKustoParameters `json:"forProvider"`
}

// FactoryLinkedServiceKustoStatus defines the observed state of FactoryLinkedServiceKusto.
type FactoryLinkedServiceKustoStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceKustoObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceKusto is the Schema for the FactoryLinkedServiceKustos API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryLinkedServiceKusto struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceKustoSpec   `json:"spec"`
	Status            FactoryLinkedServiceKustoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceKustoList contains a list of FactoryLinkedServiceKustos
type FactoryLinkedServiceKustoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceKusto `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceKusto_Kind             = "FactoryLinkedServiceKusto"
	FactoryLinkedServiceKusto_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryLinkedServiceKusto_Kind}.String()
	FactoryLinkedServiceKusto_KindAPIVersion   = FactoryLinkedServiceKusto_Kind + "." + CRDGroupVersion.String()
	FactoryLinkedServiceKusto_GroupVersionKind = CRDGroupVersion.WithKind(FactoryLinkedServiceKusto_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceKusto{}, &FactoryLinkedServiceKustoList{})
}
