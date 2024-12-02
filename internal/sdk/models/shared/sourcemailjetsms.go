// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type MailjetSms string

const (
	MailjetSmsMailjetSms MailjetSms = "mailjet-sms"
)

func (e MailjetSms) ToPointer() *MailjetSms {
	return &e
}
func (e *MailjetSms) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mailjet-sms":
		*e = MailjetSms(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MailjetSms: %v", v)
	}
}

type SourceMailjetSms struct {
	// Retrieve SMS messages created before the specified timestamp. Required format - Unix timestamp.
	EndDate    *int64     `json:"end_date,omitempty"`
	sourceType MailjetSms `const:"mailjet-sms" json:"sourceType"`
	// Retrieve SMS messages created after the specified timestamp. Required format - Unix timestamp.
	StartDate *int64 `json:"start_date,omitempty"`
	// Your access token. See <a href="https://dev.mailjet.com/sms/reference/overview/authentication">here</a>.
	Token string `json:"token"`
}

func (s SourceMailjetSms) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailjetSms) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailjetSms) GetEndDate() *int64 {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceMailjetSms) GetSourceType() MailjetSms {
	return MailjetSmsMailjetSms
}

func (o *SourceMailjetSms) GetStartDate() *int64 {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceMailjetSms) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
