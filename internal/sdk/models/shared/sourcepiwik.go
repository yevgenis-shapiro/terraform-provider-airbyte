// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Piwik string

const (
	PiwikPiwik Piwik = "piwik"
)

func (e Piwik) ToPointer() *Piwik {
	return &e
}
func (e *Piwik) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "piwik":
		*e = Piwik(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Piwik: %v", v)
	}
}

type SourcePiwik struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// The organization id appearing at URL of your piwik website
	OrganizationID string `json:"organization_id"`
	sourceType     Piwik  `const:"piwik" json:"sourceType"`
}

func (s SourcePiwik) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePiwik) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePiwik) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourcePiwik) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourcePiwik) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *SourcePiwik) GetSourceType() Piwik {
	return PiwikPiwik
}
