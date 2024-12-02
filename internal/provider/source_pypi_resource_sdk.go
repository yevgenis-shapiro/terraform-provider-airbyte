// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePypiResourceModel) ToSharedSourcePypiCreateRequest() *shared.SourcePypiCreateRequest {
	projectName := r.Configuration.ProjectName.ValueString()
	version := new(string)
	if !r.Configuration.Version.IsUnknown() && !r.Configuration.Version.IsNull() {
		*version = r.Configuration.Version.ValueString()
	} else {
		version = nil
	}
	configuration := shared.SourcePypi{
		ProjectName: projectName,
		Version:     version,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePypiCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePypiResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourcePypiResourceModel) ToSharedSourcePypiPutRequest() *shared.SourcePypiPutRequest {
	projectName := r.Configuration.ProjectName.ValueString()
	version := new(string)
	if !r.Configuration.Version.IsUnknown() && !r.Configuration.Version.IsNull() {
		*version = r.Configuration.Version.ValueString()
	} else {
		version = nil
	}
	configuration := shared.SourcePypiUpdate{
		ProjectName: projectName,
		Version:     version,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePypiPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}