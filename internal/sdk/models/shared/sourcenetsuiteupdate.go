// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceNetsuiteUpdate struct {
	// Consumer key associated with your integration
	ConsumerKey string `json:"consumer_key"`
	// Consumer secret associated with your integration
	ConsumerSecret string `json:"consumer_secret"`
	// The API names of the Netsuite objects you want to sync. Setting this speeds up the connection setup process by limiting the number of schemas that need to be retrieved from Netsuite.
	ObjectTypes []string `json:"object_types,omitempty"`
	// Netsuite realm e.g. 2344535, as for `production` or 2344535_SB1, as for the `sandbox`
	Realm string `json:"realm"`
	// Starting point for your data replication, in format of "YYYY-MM-DDTHH:mm:ssZ"
	StartDatetime string `json:"start_datetime"`
	// Access token key
	TokenKey string `json:"token_key"`
	// Access token secret
	TokenSecret string `json:"token_secret"`
	// The amount of days used to query the data with date chunks. Set smaller value, if you have lots of data.
	WindowInDays *int64 `default:"30" json:"window_in_days"`
}

func (s SourceNetsuiteUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceNetsuiteUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceNetsuiteUpdate) GetConsumerKey() string {
	if o == nil {
		return ""
	}
	return o.ConsumerKey
}

func (o *SourceNetsuiteUpdate) GetConsumerSecret() string {
	if o == nil {
		return ""
	}
	return o.ConsumerSecret
}

func (o *SourceNetsuiteUpdate) GetObjectTypes() []string {
	if o == nil {
		return nil
	}
	return o.ObjectTypes
}

func (o *SourceNetsuiteUpdate) GetRealm() string {
	if o == nil {
		return ""
	}
	return o.Realm
}

func (o *SourceNetsuiteUpdate) GetStartDatetime() string {
	if o == nil {
		return ""
	}
	return o.StartDatetime
}

func (o *SourceNetsuiteUpdate) GetTokenKey() string {
	if o == nil {
		return ""
	}
	return o.TokenKey
}

func (o *SourceNetsuiteUpdate) GetTokenSecret() string {
	if o == nil {
		return ""
	}
	return o.TokenSecret
}

func (o *SourceNetsuiteUpdate) GetWindowInDays() *int64 {
	if o == nil {
		return nil
	}
	return o.WindowInDays
}
