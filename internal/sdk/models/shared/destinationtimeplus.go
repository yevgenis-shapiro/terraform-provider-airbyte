// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Timeplus string

const (
	TimeplusTimeplus Timeplus = "timeplus"
)

func (e Timeplus) ToPointer() *Timeplus {
	return &e
}
func (e *Timeplus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "timeplus":
		*e = Timeplus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Timeplus: %v", v)
	}
}

type DestinationTimeplus struct {
	// Personal API key
	Apikey          string   `json:"apikey"`
	destinationType Timeplus `const:"timeplus" json:"destinationType"`
	// Timeplus workspace endpoint
	Endpoint *string `default:"https://us-west-2.timeplus.cloud/<workspace_id>" json:"endpoint"`
}

func (d DestinationTimeplus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTimeplus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTimeplus) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *DestinationTimeplus) GetDestinationType() Timeplus {
	return TimeplusTimeplus
}

func (o *DestinationTimeplus) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}
