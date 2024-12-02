// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Picqer string

const (
	PicqerPicqer Picqer = "picqer"
)

func (e Picqer) ToPointer() *Picqer {
	return &e
}
func (e *Picqer) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "picqer":
		*e = Picqer(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Picqer: %v", v)
	}
}

type SourcePicqer struct {
	// The organization name which is used to login to picqer
	OrganizationName string    `json:"organization_name"`
	Password         *string   `json:"password,omitempty"`
	sourceType       Picqer    `const:"picqer" json:"sourceType"`
	StartDate        time.Time `json:"start_date"`
	Username         string    `json:"username"`
}

func (s SourcePicqer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePicqer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePicqer) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *SourcePicqer) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourcePicqer) GetSourceType() Picqer {
	return PicqerPicqer
}

func (o *SourcePicqer) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourcePicqer) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}