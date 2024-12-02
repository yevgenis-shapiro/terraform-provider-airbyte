// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceHardcodedRecordsRequest struct {
	SourceHardcodedRecordsPutRequest *shared.SourceHardcodedRecordsPutRequest `request:"mediaType=application/json"`
	SourceID                         string                                   `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceHardcodedRecordsRequest) GetSourceHardcodedRecordsPutRequest() *shared.SourceHardcodedRecordsPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceHardcodedRecordsPutRequest
}

func (o *PutSourceHardcodedRecordsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceHardcodedRecordsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceHardcodedRecordsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceHardcodedRecordsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceHardcodedRecordsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}