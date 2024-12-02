// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Pipedrive string

const (
	PipedrivePipedrive Pipedrive = "pipedrive"
)

func (e Pipedrive) ToPointer() *Pipedrive {
	return &e
}
func (e *Pipedrive) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pipedrive":
		*e = Pipedrive(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pipedrive: %v", v)
	}
}

type SourcePipedrive struct {
	// The Pipedrive API Token.
	APIToken string `json:"api_token"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. When specified and not None, then stream will behave as incremental
	ReplicationStartDate string    `json:"replication_start_date"`
	sourceType           Pipedrive `const:"pipedrive" json:"sourceType"`
}

func (s SourcePipedrive) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePipedrive) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePipedrive) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourcePipedrive) GetReplicationStartDate() string {
	if o == nil {
		return ""
	}
	return o.ReplicationStartDate
}

func (o *SourcePipedrive) GetSourceType() Pipedrive {
	return PipedrivePipedrive
}
