// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMetabase struct {
	InstanceAPIURL types.String `tfsdk:"instance_api_url"`
	Password       types.String `tfsdk:"password"`
	SessionToken   types.String `tfsdk:"session_token"`
	Username       types.String `tfsdk:"username"`
}
