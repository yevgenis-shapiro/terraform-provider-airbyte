// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Expression struct {
	FieldName types.String                                                               `tfsdk:"field_name"`
	Filter    SourceGoogleAnalyticsDataAPISchemasCustomReportsArrayDimensionFilterFilter `tfsdk:"filter"`
}
