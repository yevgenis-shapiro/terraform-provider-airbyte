// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceStripeUpdate struct {
	// Your Stripe account ID (starts with 'acct_', find yours <a href="https://dashboard.stripe.com/settings/account">here</a>).
	AccountID string `json:"account_id"`
	// The number of API calls per second that you allow connector to make. This value can not be bigger than real API call rate limit (https://stripe.com/docs/rate-limits). If not specified the default maximum is 25 and 100 calls per second for test and production tokens respectively.
	CallRateLimit *int64 `json:"call_rate_limit,omitempty"`
	// Stripe API key (usually starts with 'sk_live_'; find yours <a href="https://dashboard.stripe.com/apikeys">here</a>).
	ClientSecret string `json:"client_secret"`
	// When set, the connector will always re-export data from the past N days, where N is the value set here. This is useful if your data is frequently updated after creation. The Lookback Window only applies to streams that do not support event-based incremental syncs: Events, SetupAttempts, ShippingRates, BalanceTransactions, Files, FileLinks, Refunds. More info <a href="https://docs.airbyte.com/integrations/sources/stripe#requirements">here</a>
	LookbackWindowDays *int64 `default:"0" json:"lookback_window_days"`
	// The number of worker thread to use for the sync. The performance upper boundary depends on call_rate_limit setting and type of account.
	NumWorkers *int64 `default:"10" json:"num_workers"`
	// The time increment used by the connector when requesting data from the Stripe API. The bigger the value is, the less requests will be made and faster the sync will be. On the other hand, the more seldom the state is persisted.
	SliceRange *int64 `default:"365" json:"slice_range"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Only data generated after this date will be replicated.
	StartDate *time.Time `default:"2017-01-25T00:00:00Z" json:"start_date"`
}

func (s SourceStripeUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceStripeUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceStripeUpdate) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *SourceStripeUpdate) GetCallRateLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.CallRateLimit
}

func (o *SourceStripeUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceStripeUpdate) GetLookbackWindowDays() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindowDays
}

func (o *SourceStripeUpdate) GetNumWorkers() *int64 {
	if o == nil {
		return nil
	}
	return o.NumWorkers
}

func (o *SourceStripeUpdate) GetSliceRange() *int64 {
	if o == nil {
		return nil
	}
	return o.SliceRange
}

func (o *SourceStripeUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
