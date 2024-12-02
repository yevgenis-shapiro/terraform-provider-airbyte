// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSalesloftResourceModel) ToSharedSourceSalesloftCreateRequest() *shared.SourceSalesloftCreateRequest {
	var credentials shared.SourceSalesloftCredentials
	var sourceSalesloftAuthenticateViaOAuth *shared.SourceSalesloftAuthenticateViaOAuth
	if r.Configuration.Credentials.AuthenticateViaOAuth != nil {
		accessToken := r.Configuration.Credentials.AuthenticateViaOAuth.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.AuthenticateViaOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaOAuth.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.AuthenticateViaOAuth.TokenExpiryDate.ValueString())
		sourceSalesloftAuthenticateViaOAuth = &shared.SourceSalesloftAuthenticateViaOAuth{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceSalesloftAuthenticateViaOAuth != nil {
		credentials = shared.SourceSalesloftCredentials{
			SourceSalesloftAuthenticateViaOAuth: sourceSalesloftAuthenticateViaOAuth,
		}
	}
	var sourceSalesloftAuthenticateViaAPIKey *shared.SourceSalesloftAuthenticateViaAPIKey
	if r.Configuration.Credentials.AuthenticateViaAPIKey != nil {
		apiKey := r.Configuration.Credentials.AuthenticateViaAPIKey.APIKey.ValueString()
		sourceSalesloftAuthenticateViaAPIKey = &shared.SourceSalesloftAuthenticateViaAPIKey{
			APIKey: apiKey,
		}
	}
	if sourceSalesloftAuthenticateViaAPIKey != nil {
		credentials = shared.SourceSalesloftCredentials{
			SourceSalesloftAuthenticateViaAPIKey: sourceSalesloftAuthenticateViaAPIKey,
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSalesloft{
		Credentials: credentials,
		StartDate:   startDate,
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
	out := shared.SourceSalesloftCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSalesloftResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSalesloftResourceModel) ToSharedSourceSalesloftPutRequest() *shared.SourceSalesloftPutRequest {
	var credentials shared.SourceSalesloftUpdateCredentials
	var authenticateViaOAuth *shared.AuthenticateViaOAuth
	if r.Configuration.Credentials.AuthenticateViaOAuth != nil {
		accessToken := r.Configuration.Credentials.AuthenticateViaOAuth.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.AuthenticateViaOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaOAuth.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.AuthenticateViaOAuth.TokenExpiryDate.ValueString())
		authenticateViaOAuth = &shared.AuthenticateViaOAuth{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if authenticateViaOAuth != nil {
		credentials = shared.SourceSalesloftUpdateCredentials{
			AuthenticateViaOAuth: authenticateViaOAuth,
		}
	}
	var authenticateViaAPIKey *shared.AuthenticateViaAPIKey
	if r.Configuration.Credentials.AuthenticateViaAPIKey != nil {
		apiKey := r.Configuration.Credentials.AuthenticateViaAPIKey.APIKey.ValueString()
		authenticateViaAPIKey = &shared.AuthenticateViaAPIKey{
			APIKey: apiKey,
		}
	}
	if authenticateViaAPIKey != nil {
		credentials = shared.SourceSalesloftUpdateCredentials{
			AuthenticateViaAPIKey: authenticateViaAPIKey,
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSalesloftUpdate{
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSalesloftPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
