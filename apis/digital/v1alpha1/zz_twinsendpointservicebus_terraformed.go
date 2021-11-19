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
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/resource"
	"github.com/crossplane-contrib/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this TwinsEndpointServicebus
func (mg *TwinsEndpointServicebus) GetTerraformResourceType() string {
	return "azurerm_digital_twins_endpoint_servicebus"
}

// GetConnectionDetailsMapping for this TwinsEndpointServicebus
func (tr *TwinsEndpointServicebus) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"dead_letter_storage_secret": "spec.forProvider.deadLetterStorageSecretSecretRef", "servicebus_primary_connection_string": "spec.forProvider.servicebusPrimaryConnectionStringSecretRef", "servicebus_secondary_connection_string": "spec.forProvider.servicebusSecondaryConnectionStringSecretRef"}
}

// GetObservation of this TwinsEndpointServicebus
func (tr *TwinsEndpointServicebus) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TwinsEndpointServicebus
func (tr *TwinsEndpointServicebus) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this TwinsEndpointServicebus
func (tr *TwinsEndpointServicebus) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TwinsEndpointServicebus
func (tr *TwinsEndpointServicebus) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this TwinsEndpointServicebus using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TwinsEndpointServicebus) LateInitialize(attrs []byte) (bool, error) {
	params := &TwinsEndpointServicebusParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TwinsEndpointServicebus) GetTerraformSchemaVersion() int {
	return 0
}
