// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMssqlEncryptedVerifyCertificate struct {
	Certificate           types.String `tfsdk:"certificate"`
	HostNameInCertificate types.String `tfsdk:"host_name_in_certificate"`
}
