/*
Copyright 2021 The Crossplane Authors.

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

package sql

import (
	"fmt"

	"github.com/crossplane/terrajet/pkg/config"
	"github.com/pkg/errors"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

func msSQLConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) {
	dbName, ok := attr["name"].(string)
	if !ok {
		return nil, errors.New("cannot get name attribute")
	}
	username, ok := attr["administrator_login"].(string)
	if !ok {
		return nil, errors.New("cannot get administrator_login attribute")
	}

	endpoint, ok := attr["fully_qualified_domain_name"].(string)
	if !ok {
		return nil, errors.New("cannot get fully_qualified_domain_name attribute")
	}
	conn := map[string][]byte{
		xpv1.ResourceCredentialsSecretUserKey:     []byte(fmt.Sprintf("%s@%s", username, dbName)),
		xpv1.ResourceCredentialsSecretEndpointKey: []byte(endpoint),
	}

	// Note(turkenh): include password only if available, might not be available
	//  depending on the auth type.
	if password, ok := attr["administrator_login_password"].(string); ok {
		conn[xpv1.ResourceCredentialsSecretPasswordKey] = []byte(password)
	}

	return conn, nil
}

// Configure configures sql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_sql_server", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"threat_detection_policy"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.Sensitive.AdditionalConnectionDetailsFn = msSQLConnectionDetails
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Sql",
			"servers", "name",
		)
		r.UseAsync = true
	})
	p.AddResourceConfigurator("azurerm_mssql_server", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.Sensitive.AdditionalConnectionDetailsFn = msSQLConnectionDetails
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Sql",
			"servers", "name",
		)
		r.UseAsync = true
	})
}
