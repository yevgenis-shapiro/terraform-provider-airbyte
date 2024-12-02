// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationAstraProcessingConfigModel struct {
	ChunkOverlap      types.Int64                   `tfsdk:"chunk_overlap"`
	ChunkSize         types.Int64                   `tfsdk:"chunk_size"`
	FieldNameMappings []FieldNameMappingConfigModel `tfsdk:"field_name_mappings"`
	MetadataFields    []types.String                `tfsdk:"metadata_fields"`
	TextFields        []types.String                `tfsdk:"text_fields"`
	TextSplitter      *DestinationAstraTextSplitter `tfsdk:"text_splitter"`
}
