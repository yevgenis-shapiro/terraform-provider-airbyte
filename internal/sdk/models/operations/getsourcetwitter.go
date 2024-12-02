// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type GetSourceTwitterRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *GetSourceTwitterRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type GetSourceTwitterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Get a Source by the id in the path.
	SourceResponse *shared.SourceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSourceTwitterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSourceTwitterResponse) GetSourceResponse() *shared.SourceResponse {
	if o == nil {
		return nil
	}
	return o.SourceResponse
}

func (o *GetSourceTwitterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSourceTwitterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
