// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceNylas struct {
	APIKey    types.String `tfsdk:"api_key"`
	APIServer types.String `tfsdk:"api_server"`
	EndDate   types.String `tfsdk:"end_date"`
	StartDate types.String `tfsdk:"start_date"`
}
