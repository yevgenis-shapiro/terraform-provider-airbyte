// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTwilioResourceModel) ToSharedSourceTwilioCreateRequest() *shared.SourceTwilioCreateRequest {
	accountSid := r.Configuration.AccountSid.ValueString()
	authToken := r.Configuration.AuthToken.ValueString()
	lookbackWindow := new(int64)
	if !r.Configuration.LookbackWindow.IsUnknown() && !r.Configuration.LookbackWindow.IsNull() {
		*lookbackWindow = r.Configuration.LookbackWindow.ValueInt64()
	} else {
		lookbackWindow = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceTwilio{
		AccountSid:     accountSid,
		AuthToken:      authToken,
		LookbackWindow: lookbackWindow,
		StartDate:      startDate,
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
	out := shared.SourceTwilioCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTwilioResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceTwilioResourceModel) ToSharedSourceTwilioPutRequest() *shared.SourceTwilioPutRequest {
	accountSid := r.Configuration.AccountSid.ValueString()
	authToken := r.Configuration.AuthToken.ValueString()
	lookbackWindow := new(int64)
	if !r.Configuration.LookbackWindow.IsUnknown() && !r.Configuration.LookbackWindow.IsNull() {
		*lookbackWindow = r.Configuration.LookbackWindow.ValueInt64()
	} else {
		lookbackWindow = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceTwilioUpdate{
		AccountSid:     accountSid,
		AuthToken:      authToken,
		LookbackWindow: lookbackWindow,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTwilioPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
