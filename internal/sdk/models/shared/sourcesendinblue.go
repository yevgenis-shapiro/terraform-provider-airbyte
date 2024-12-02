// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Sendinblue string

const (
	SendinblueSendinblue Sendinblue = "sendinblue"
)

func (e Sendinblue) ToPointer() *Sendinblue {
	return &e
}
func (e *Sendinblue) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sendinblue":
		*e = Sendinblue(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sendinblue: %v", v)
	}
}

type SourceSendinblue struct {
	// Your API Key. See <a href="https://developers.sendinblue.com/docs/getting-started">here</a>.
	APIKey     string     `json:"api_key"`
	sourceType Sendinblue `const:"sendinblue" json:"sourceType"`
}

func (s SourceSendinblue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSendinblue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSendinblue) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceSendinblue) GetSourceType() Sendinblue {
	return SendinblueSendinblue
}