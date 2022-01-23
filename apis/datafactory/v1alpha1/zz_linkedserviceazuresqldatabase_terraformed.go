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

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this LinkedServiceAzureSQLDatabase
func (mg *LinkedServiceAzureSQLDatabase) GetTerraformResourceType() string {
	return "azurerm_data_factory_linked_service_azure_sql_database"
}

// GetConnectionDetailsMapping for this LinkedServiceAzureSQLDatabase
func (tr *LinkedServiceAzureSQLDatabase) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LinkedServiceAzureSQLDatabase
func (tr *LinkedServiceAzureSQLDatabase) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LinkedServiceAzureSQLDatabase
func (tr *LinkedServiceAzureSQLDatabase) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LinkedServiceAzureSQLDatabase
func (tr *LinkedServiceAzureSQLDatabase) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LinkedServiceAzureSQLDatabase
func (tr *LinkedServiceAzureSQLDatabase) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LinkedServiceAzureSQLDatabase
func (tr *LinkedServiceAzureSQLDatabase) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LinkedServiceAzureSQLDatabase using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LinkedServiceAzureSQLDatabase) LateInitialize(attrs []byte) (bool, error) {
	params := &LinkedServiceAzureSQLDatabaseParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LinkedServiceAzureSQLDatabase) GetTerraformSchemaVersion() int {
	return 0
}
