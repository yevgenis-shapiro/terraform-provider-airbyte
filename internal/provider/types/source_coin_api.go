// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceCoinAPI struct {
	APIKey      types.String `tfsdk:"api_key"`
	EndDate     types.String `tfsdk:"end_date"`
	Environment types.String `tfsdk:"environment"`
	Limit       types.Int64  `tfsdk:"limit"`
	Period      types.String `tfsdk:"period"`
	StartDate   types.String `tfsdk:"start_date"`
	SymbolID    types.String `tfsdk:"symbol_id"`
}
