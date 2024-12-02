// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceStravaAuthType string

const (
	SourceStravaAuthTypeClient SourceStravaAuthType = "Client"
)

func (e SourceStravaAuthType) ToPointer() *SourceStravaAuthType {
	return &e
}
func (e *SourceStravaAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceStravaAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceStravaAuthType: %v", v)
	}
}

type Strava string

const (
	StravaStrava Strava = "strava"
)

func (e Strava) ToPointer() *Strava {
	return &e
}
func (e *Strava) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "strava":
		*e = Strava(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Strava: %v", v)
	}
}

type SourceStrava struct {
	// The Athlete ID of your Strava developer application.
	AthleteID int64                 `json:"athlete_id"`
	authType  *SourceStravaAuthType `const:"Client" json:"auth_type"`
	// The Client ID of your Strava developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Strava developer application.
	ClientSecret string `json:"client_secret"`
	// The Refresh Token with the activity: read_all permissions.
	RefreshToken string `json:"refresh_token"`
	sourceType   Strava `const:"strava" json:"sourceType"`
	// UTC date and time. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceStrava) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceStrava) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceStrava) GetAthleteID() int64 {
	if o == nil {
		return 0
	}
	return o.AthleteID
}

func (o *SourceStrava) GetAuthType() *SourceStravaAuthType {
	return SourceStravaAuthTypeClient.ToPointer()
}

func (o *SourceStrava) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceStrava) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceStrava) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceStrava) GetSourceType() Strava {
	return StravaStrava
}

func (o *SourceStrava) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
