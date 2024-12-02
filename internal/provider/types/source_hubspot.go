// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceHubspot struct {
	Credentials               SourceHubspotAuthentication `tfsdk:"credentials"`
	EnableExperimentalStreams types.Bool                  `tfsdk:"enable_experimental_streams"`
	StartDate                 types.String                `tfsdk:"start_date"`
}
