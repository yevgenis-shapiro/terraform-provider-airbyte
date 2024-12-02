// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceClickupAPIUpdate struct {
	// Every ClickUp API call required authentication. This field is your personal API token. See <a href="https://clickup.com/api/developer-portal/authentication/#personal-token">here</a>.
	APIToken string `json:"api_token"`
	// Include or exclude closed tasks. By default, they are excluded. See <a https://clickup.com/api/clickupreference/operation/GetTasks/#!in=query&path=include_closed&t=request">here</a>.
	IncludeClosedTasks *bool `default:"false" json:"include_closed_tasks"`
}

func (s SourceClickupAPIUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClickupAPIUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceClickupAPIUpdate) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceClickupAPIUpdate) GetIncludeClosedTasks() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeClosedTasks
}