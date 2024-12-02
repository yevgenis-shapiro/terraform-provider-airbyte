// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourcePokeapiRequest struct {
	SourcePokeapiPutRequest *shared.SourcePokeapiPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourcePokeapiRequest) GetSourcePokeapiPutRequest() *shared.SourcePokeapiPutRequest {
	if o == nil {
		return nil
	}
	return o.SourcePokeapiPutRequest
}

func (o *PutSourcePokeapiRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourcePokeapiResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourcePokeapiResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourcePokeapiResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourcePokeapiResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
