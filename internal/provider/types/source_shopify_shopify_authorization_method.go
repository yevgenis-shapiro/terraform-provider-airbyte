// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceShopifyShopifyAuthorizationMethod struct {
	APIPassword *APIPassword          `tfsdk:"api_password" tfPlanOnly:"true"`
	OAuth20     *SourceShopifyOAuth20 `tfsdk:"o_auth20" tfPlanOnly:"true"`
}