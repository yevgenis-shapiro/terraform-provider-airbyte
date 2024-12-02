// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceLookerRequest struct {
	SourceLookerPutRequest *shared.SourceLookerPutRequest `request:"mediaType=application/json"`
	SourceID               string                         `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceLookerRequest) GetSourceLookerPutRequest() *shared.SourceLookerPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceLookerPutRequest
}

func (o *PutSourceLookerRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceLookerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceLookerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceLookerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceLookerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
