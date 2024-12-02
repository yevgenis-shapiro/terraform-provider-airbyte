// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type DestinationTeradataSSLModes struct {
	Allow      *Fake                        `tfsdk:"allow" tfPlanOnly:"true"`
	Disable    *Fake                        `tfsdk:"disable" tfPlanOnly:"true"`
	Prefer     *Fake                        `tfsdk:"prefer" tfPlanOnly:"true"`
	Require    *Fake                        `tfsdk:"require" tfPlanOnly:"true"`
	VerifyCa   *DestinationTeradataVerifyCa `tfsdk:"verify_ca" tfPlanOnly:"true"`
	VerifyFull *DestinationTeradataVerifyCa `tfsdk:"verify_full" tfPlanOnly:"true"`
}