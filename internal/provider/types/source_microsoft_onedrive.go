// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMicrosoftOnedrive struct {
	Credentials SourceMicrosoftOnedriveAuthentication          `tfsdk:"credentials"`
	DriveName   types.String                                   `tfsdk:"drive_name"`
	FolderPath  types.String                                   `tfsdk:"folder_path"`
	SearchScope types.String                                   `tfsdk:"search_scope"`
	StartDate   types.String                                   `tfsdk:"start_date"`
	Streams     []SourceMicrosoftOnedriveFileBasedStreamConfig `tfsdk:"streams"`
}