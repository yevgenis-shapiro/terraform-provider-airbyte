// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Clockify string

const (
	ClockifyClockify Clockify = "clockify"
)

func (e Clockify) ToPointer() *Clockify {
	return &e
}
func (e *Clockify) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clockify":
		*e = Clockify(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Clockify: %v", v)
	}
}

type SourceClockify struct {
	// You can get your api access_key <a href="https://app.clockify.me/user/settings">here</a> This API is Case Sensitive.
	APIKey string `json:"api_key"`
	// The URL for the Clockify API. This should only need to be modified if connecting to an enterprise version of Clockify.
	APIURL     *string  `default:"https://api.clockify.me" json:"api_url"`
	sourceType Clockify `const:"clockify" json:"sourceType"`
	// WorkSpace Id
	WorkspaceID string `json:"workspace_id"`
}

func (s SourceClockify) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClockify) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceClockify) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceClockify) GetAPIURL() *string {
	if o == nil {
		return nil
	}
	return o.APIURL
}

func (o *SourceClockify) GetSourceType() Clockify {
	return ClockifyClockify
}

func (o *SourceClockify) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
