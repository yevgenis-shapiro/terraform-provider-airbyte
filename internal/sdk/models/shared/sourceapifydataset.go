// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type ApifyDataset string

const (
	ApifyDatasetApifyDataset ApifyDataset = "apify-dataset"
)

func (e ApifyDataset) ToPointer() *ApifyDataset {
	return &e
}
func (e *ApifyDataset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apify-dataset":
		*e = ApifyDataset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApifyDataset: %v", v)
	}
}

type SourceApifyDataset struct {
	// ID of the dataset you would like to load to Airbyte. In Apify Console, you can view your datasets in the <a href="https://console.apify.com/storage/datasets">Storage section under the Datasets tab</a> after you login. See the <a href="https://docs.apify.com/platform/storage/dataset">Apify Docs</a> for more information.
	DatasetID  string       `json:"dataset_id"`
	sourceType ApifyDataset `const:"apify-dataset" json:"sourceType"`
	// Personal API token of your Apify account. In Apify Console, you can find your API token in the <a href="https://console.apify.com/account/integrations">Settings section under the Integrations tab</a> after you login. See the <a href="https://docs.apify.com/platform/integrations/api#api-token">Apify Docs</a> for more information.
	Token string `json:"token"`
}

func (s SourceApifyDataset) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceApifyDataset) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceApifyDataset) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

func (o *SourceApifyDataset) GetSourceType() ApifyDataset {
	return ApifyDatasetApifyDataset
}

func (o *SourceApifyDataset) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}