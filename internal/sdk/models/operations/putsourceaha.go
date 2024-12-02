// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceAhaRequest struct {
	SourceAhaPutRequest *shared.SourceAhaPutRequest `request:"mediaType=application/json"`
	SourceID            string                      `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceAhaRequest) GetSourceAhaPutRequest() *shared.SourceAhaPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceAhaPutRequest
}

func (o *PutSourceAhaRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceAhaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceAhaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceAhaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceAhaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
