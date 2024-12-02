// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcs struct {
	Credential      DestinationGcsAuthentication `tfsdk:"credential"`
	Format          DestinationGcsOutputFormat   `tfsdk:"format"`
	GcsBucketName   types.String                 `tfsdk:"gcs_bucket_name"`
	GcsBucketPath   types.String                 `tfsdk:"gcs_bucket_path"`
	GcsBucketRegion types.String                 `tfsdk:"gcs_bucket_region"`
}
