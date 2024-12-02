// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceAmplitudeRequest struct {
	SourceAmplitudePutRequest *shared.SourceAmplitudePutRequest `request:"mediaType=application/json"`
	SourceID                  string                            `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceAmplitudeRequest) GetSourceAmplitudePutRequest() *shared.SourceAmplitudePutRequest {
	if o == nil {
		return nil
	}
	return o.SourceAmplitudePutRequest
}

func (o *PutSourceAmplitudeRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceAmplitudeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceAmplitudeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceAmplitudeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceAmplitudeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
