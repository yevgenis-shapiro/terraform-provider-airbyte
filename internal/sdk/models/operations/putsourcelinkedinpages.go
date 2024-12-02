// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceLinkedinPagesRequest struct {
	SourceLinkedinPagesPutRequest *shared.SourceLinkedinPagesPutRequest `request:"mediaType=application/json"`
	SourceID                      string                                `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceLinkedinPagesRequest) GetSourceLinkedinPagesPutRequest() *shared.SourceLinkedinPagesPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceLinkedinPagesPutRequest
}

func (o *PutSourceLinkedinPagesRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceLinkedinPagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceLinkedinPagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceLinkedinPagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceLinkedinPagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
