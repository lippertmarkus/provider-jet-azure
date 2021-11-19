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

// GetTerraformResourceType returns Terraform resource type for this VirtualMachine
func (mg *VirtualMachine) GetTerraformResourceType() string {
	return "azurerm_mssql_virtual_machine"
}

// GetConnectionDetailsMapping for this VirtualMachine
func (tr *VirtualMachine) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"auto_backup[*].encryption_password": "spec.forProvider.autoBackup[*].encryptionPasswordSecretRef", "key_vault_credential[*].key_vault_url": "spec.forProvider.keyVaultCredential[*].keyVaultURLSecretRef", "key_vault_credential[*].service_principal_name": "spec.forProvider.keyVaultCredential[*].servicePrincipalNameSecretRef", "key_vault_credential[*].service_principal_secret": "spec.forProvider.keyVaultCredential[*].servicePrincipalSecretSecretRef", "sql_connectivity_update_password": "spec.forProvider.sqlConnectivityUpdatePasswordSecretRef", "sql_connectivity_update_username": "spec.forProvider.sqlConnectivityUpdateUsernameSecretRef"}
}

// GetObservation of this VirtualMachine
func (tr *VirtualMachine) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualMachine
func (tr *VirtualMachine) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this VirtualMachine
func (tr *VirtualMachine) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualMachine
func (tr *VirtualMachine) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualMachine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualMachine) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualMachineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualMachine) GetTerraformSchemaVersion() int {
	return 0
}
