// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceBuildkiteRequest struct {
	SourceBuildkitePutRequest *shared.SourceBuildkitePutRequest `request:"mediaType=application/json"`
	SourceID                  string                            `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceBuildkiteRequest) GetSourceBuildkitePutRequest() *shared.SourceBuildkitePutRequest {
	if o == nil {
		return nil
	}
	return o.SourceBuildkitePutRequest
}

func (o *PutSourceBuildkiteRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceBuildkiteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceBuildkiteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceBuildkiteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceBuildkiteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
