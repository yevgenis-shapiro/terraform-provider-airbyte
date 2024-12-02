// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcsParquetColumnarStorage struct {
	BlockSizeMb          types.Int64  `tfsdk:"block_size_mb"`
	CompressionCodec     types.String `tfsdk:"compression_codec"`
	DictionaryEncoding   types.Bool   `tfsdk:"dictionary_encoding"`
	DictionaryPageSizeKb types.Int64  `tfsdk:"dictionary_page_size_kb"`
	FormatType           types.String `tfsdk:"format_type"`
	MaxPaddingSizeMb     types.Int64  `tfsdk:"max_padding_size_mb"`
	PageSizeKb           types.Int64  `tfsdk:"page_size_kb"`
}
