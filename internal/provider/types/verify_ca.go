// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type VerifyCa struct {
	CaCertificate     types.String `tfsdk:"ca_certificate"`
	ClientKeyPassword types.String `tfsdk:"client_key_password"`
}
