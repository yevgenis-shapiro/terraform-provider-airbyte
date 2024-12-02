// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceJira struct {
	APIToken                  types.String   `tfsdk:"api_token"`
	Domain                    types.String   `tfsdk:"domain"`
	Email                     types.String   `tfsdk:"email"`
	EnableExperimentalStreams types.Bool     `tfsdk:"enable_experimental_streams"`
	LookbackWindowMinutes     types.Int64    `tfsdk:"lookback_window_minutes"`
	Projects                  []types.String `tfsdk:"projects"`
	StartDate                 types.String   `tfsdk:"start_date"`
}