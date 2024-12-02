// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAsana struct {
	Credentials           *SourceAsanaAuthenticationMechanism `tfsdk:"credentials"`
	OrganizationExportIds []types.String                      `tfsdk:"organization_export_ids"`
}