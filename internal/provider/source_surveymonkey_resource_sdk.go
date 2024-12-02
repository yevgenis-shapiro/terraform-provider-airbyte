// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSurveymonkeyResourceModel) ToSharedSourceSurveymonkeyCreateRequest() *shared.SourceSurveymonkeyCreateRequest {
	accessToken := r.Configuration.Credentials.AccessToken.ValueString()
	clientID := new(string)
	if !r.Configuration.Credentials.ClientID.IsUnknown() && !r.Configuration.Credentials.ClientID.IsNull() {
		*clientID = r.Configuration.Credentials.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.Credentials.ClientSecret.IsUnknown() && !r.Configuration.Credentials.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.Credentials.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	credentials := shared.SourceSurveymonkeySurveyMonkeyAuthorizationMethod{
		AccessToken:  accessToken,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	origin := new(shared.SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccount)
	if !r.Configuration.Origin.IsUnknown() && !r.Configuration.Origin.IsNull() {
		*origin = shared.SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccount(r.Configuration.Origin.ValueString())
	} else {
		origin = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	var surveyIds []string = []string{}
	for _, surveyIdsItem := range r.Configuration.SurveyIds {
		surveyIds = append(surveyIds, surveyIdsItem.ValueString())
	}
	configuration := shared.SourceSurveymonkey{
		Credentials: credentials,
		Origin:      origin,
		StartDate:   startDate,
		SurveyIds:   surveyIds,
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
	out := shared.SourceSurveymonkeyCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSurveymonkeyResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSurveymonkeyResourceModel) ToSharedSourceSurveymonkeyPutRequest() *shared.SourceSurveymonkeyPutRequest {
	accessToken := r.Configuration.Credentials.AccessToken.ValueString()
	clientID := new(string)
	if !r.Configuration.Credentials.ClientID.IsUnknown() && !r.Configuration.Credentials.ClientID.IsNull() {
		*clientID = r.Configuration.Credentials.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.Credentials.ClientSecret.IsUnknown() && !r.Configuration.Credentials.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.Credentials.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	credentials := shared.SurveyMonkeyAuthorizationMethod{
		AccessToken:  accessToken,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	origin := new(shared.OriginDatacenterOfTheSurveyMonkeyAccount)
	if !r.Configuration.Origin.IsUnknown() && !r.Configuration.Origin.IsNull() {
		*origin = shared.OriginDatacenterOfTheSurveyMonkeyAccount(r.Configuration.Origin.ValueString())
	} else {
		origin = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	var surveyIds []string = []string{}
	for _, surveyIdsItem := range r.Configuration.SurveyIds {
		surveyIds = append(surveyIds, surveyIdsItem.ValueString())
	}
	configuration := shared.SourceSurveymonkeyUpdate{
		Credentials: credentials,
		Origin:      origin,
		StartDate:   startDate,
		SurveyIds:   surveyIds,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSurveymonkeyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}