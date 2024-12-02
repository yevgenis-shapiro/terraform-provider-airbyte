// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationSftpJSONResourceModel) ToSharedDestinationSftpJSONCreateRequest() *shared.DestinationSftpJSONCreateRequest {
	destinationPath := r.Configuration.DestinationPath.ValueString()
	host := r.Configuration.Host.ValueString()
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationSftpJSON{
		DestinationPath: destinationPath,
		Host:            host,
		Password:        password,
		Port:            port,
		Username:        username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSftpJSONCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSftpJSONResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationSftpJSONResourceModel) ToSharedDestinationSftpJSONPutRequest() *shared.DestinationSftpJSONPutRequest {
	destinationPath := r.Configuration.DestinationPath.ValueString()
	host := r.Configuration.Host.ValueString()
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationSftpJSONUpdate{
		DestinationPath: destinationPath,
		Host:            host,
		Password:        password,
		Port:            port,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSftpJSONPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}