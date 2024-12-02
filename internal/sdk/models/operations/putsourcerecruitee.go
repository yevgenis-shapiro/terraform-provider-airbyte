// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceRecruiteeRequest struct {
	SourceRecruiteePutRequest *shared.SourceRecruiteePutRequest `request:"mediaType=application/json"`
	SourceID                  string                            `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceRecruiteeRequest) GetSourceRecruiteePutRequest() *shared.SourceRecruiteePutRequest {
	if o == nil {
		return nil
	}
	return o.SourceRecruiteePutRequest
}

func (o *PutSourceRecruiteeRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceRecruiteeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceRecruiteeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceRecruiteeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceRecruiteeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
