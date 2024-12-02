// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceChameleonResourceModel) ToSharedSourceChameleonCreateRequest() *shared.SourceChameleonCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	endDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	filter := new(shared.SourceChameleonFilter)
	if !r.Configuration.Filter.IsUnknown() && !r.Configuration.Filter.IsNull() {
		*filter = shared.SourceChameleonFilter(r.Configuration.Filter.ValueString())
	} else {
		filter = nil
	}
	limit := new(string)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueString()
	} else {
		limit = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceChameleon{
		APIKey:    apiKey,
		EndDate:   endDate,
		Filter:    filter,
		Limit:     limit,
		StartDate: startDate,
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
	out := shared.SourceChameleonCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceChameleonResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceChameleonResourceModel) ToSharedSourceChameleonPutRequest() *shared.SourceChameleonPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	endDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	filter := new(shared.Filter)
	if !r.Configuration.Filter.IsUnknown() && !r.Configuration.Filter.IsNull() {
		*filter = shared.Filter(r.Configuration.Filter.ValueString())
	} else {
		filter = nil
	}
	limit := new(string)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueString()
	} else {
		limit = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceChameleonUpdate{
		APIKey:    apiKey,
		EndDate:   endDate,
		Filter:    filter,
		Limit:     limit,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceChameleonPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
