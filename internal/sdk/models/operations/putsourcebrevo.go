// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceBrevoRequest struct {
	SourceBrevoPutRequest *shared.SourceBrevoPutRequest `request:"mediaType=application/json"`
	SourceID              string                        `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceBrevoRequest) GetSourceBrevoPutRequest() *shared.SourceBrevoPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceBrevoPutRequest
}

func (o *PutSourceBrevoRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceBrevoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceBrevoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceBrevoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceBrevoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
