// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationPineconeIndexing struct {
	Index               types.String `tfsdk:"index"`
	PineconeEnvironment types.String `tfsdk:"pinecone_environment"`
	PineconeKey         types.String `tfsdk:"pinecone_key"`
}
