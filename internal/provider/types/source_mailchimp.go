// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMailchimp struct {
	Credentials *SourceMailchimpAuthentication `tfsdk:"credentials"`
	StartDate   types.String                   `tfsdk:"start_date"`
}