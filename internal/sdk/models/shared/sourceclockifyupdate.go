// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceClockifyUpdate struct {
	// You can get your api access_key <a href="https://app.clockify.me/user/settings">here</a> This API is Case Sensitive.
	APIKey string `json:"api_key"`
	// The URL for the Clockify API. This should only need to be modified if connecting to an enterprise version of Clockify.
	APIURL *string `default:"https://api.clockify.me" json:"api_url"`
	// WorkSpace Id
	WorkspaceID string `json:"workspace_id"`
}

func (s SourceClockifyUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClockifyUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceClockifyUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceClockifyUpdate) GetAPIURL() *string {
	if o == nil {
		return nil
	}
	return o.APIURL
}

func (o *SourceClockifyUpdate) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
