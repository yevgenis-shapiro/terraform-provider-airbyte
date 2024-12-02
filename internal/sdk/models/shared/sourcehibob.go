// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Hibob string

const (
	HibobHibob Hibob = "hibob"
)

func (e Hibob) ToPointer() *Hibob {
	return &e
}
func (e *Hibob) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hibob":
		*e = Hibob(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hibob: %v", v)
	}
}

type SourceHibob struct {
	// Toggle true if this instance is a HiBob sandbox
	IsSandbox  bool    `json:"is_sandbox"`
	Password   *string `json:"password,omitempty"`
	sourceType Hibob   `const:"hibob" json:"sourceType"`
	Username   string  `json:"username"`
}

func (s SourceHibob) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHibob) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceHibob) GetIsSandbox() bool {
	if o == nil {
		return false
	}
	return o.IsSandbox
}

func (o *SourceHibob) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceHibob) GetSourceType() Hibob {
	return HibobHibob
}

func (o *SourceHibob) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}