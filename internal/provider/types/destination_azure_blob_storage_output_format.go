// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type DestinationAzureBlobStorageOutputFormat struct {
	CSVCommaSeparatedValues       *CSVCommaSeparatedValues                                  `tfsdk:"csv_comma_separated_values" tfPlanOnly:"true"`
	JSONLinesNewlineDelimitedJSON *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON `tfsdk:"json_lines_newline_delimited_json" tfPlanOnly:"true"`
}
