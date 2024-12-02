// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceIntercomResourceModel) ToSharedSourceIntercomCreateRequest() *shared.SourceIntercomCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	activityLogsTimeStep := new(int64)
	if !r.Configuration.ActivityLogsTimeStep.IsUnknown() && !r.Configuration.ActivityLogsTimeStep.IsNull() {
		*activityLogsTimeStep = r.Configuration.ActivityLogsTimeStep.ValueInt64()
	} else {
		activityLogsTimeStep = nil
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceIntercom{
		AccessToken:          accessToken,
		ActivityLogsTimeStep: activityLogsTimeStep,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		StartDate:            startDate,
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
	out := shared.SourceIntercomCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceIntercomResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceIntercomResourceModel) ToSharedSourceIntercomPutRequest() *shared.SourceIntercomPutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	activityLogsTimeStep := new(int64)
	if !r.Configuration.ActivityLogsTimeStep.IsUnknown() && !r.Configuration.ActivityLogsTimeStep.IsNull() {
		*activityLogsTimeStep = r.Configuration.ActivityLogsTimeStep.ValueInt64()
	} else {
		activityLogsTimeStep = nil
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceIntercomUpdate{
		AccessToken:          accessToken,
		ActivityLogsTimeStep: activityLogsTimeStep,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		StartDate:            startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceIntercomPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}