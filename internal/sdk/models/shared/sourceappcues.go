// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Appcues string

const (
	AppcuesAppcues Appcues = "appcues"
)

func (e Appcues) ToPointer() *Appcues {
	return &e
}
func (e *Appcues) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "appcues":
		*e = Appcues(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Appcues: %v", v)
	}
}

type SourceAppcues struct {
	// Account ID of Appcues found in account settings page (https://studio.appcues.com/settings/account)
	AccountID  string    `json:"account_id"`
	Password   *string   `json:"password,omitempty"`
	sourceType Appcues   `const:"appcues" json:"sourceType"`
	StartDate  time.Time `json:"start_date"`
	Username   string    `json:"username"`
}

func (s SourceAppcues) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAppcues) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAppcues) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *SourceAppcues) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceAppcues) GetSourceType() Appcues {
	return AppcuesAppcues
}

func (o *SourceAppcues) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceAppcues) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
