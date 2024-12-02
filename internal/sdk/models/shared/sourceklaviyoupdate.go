// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceKlaviyoUpdate struct {
	// Klaviyo API Key. See our <a href="https://docs.airbyte.com/integrations/sources/klaviyo">docs</a> if you need help finding this key.
	APIKey string `json:"api_key"`
	// Certain streams like the profiles stream can retrieve predictive analytics data from Klaviyo's API. However, at high volume, this can lead to service availability issues on the API which can be improved by not fetching this field. WARNING: Enabling this setting will stop the  "predictive_analytics" column from being populated in your downstream destination.
	DisableFetchingPredictiveAnalytics *bool `json:"disable_fetching_predictive_analytics,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. This field is optional - if not provided, all data will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceKlaviyoUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceKlaviyoUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceKlaviyoUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceKlaviyoUpdate) GetDisableFetchingPredictiveAnalytics() *bool {
	if o == nil {
		return nil
	}
	return o.DisableFetchingPredictiveAnalytics
}

func (o *SourceKlaviyoUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
