// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceSftpBulkRequest struct {
	SourceSftpBulkPutRequest *shared.SourceSftpBulkPutRequest `request:"mediaType=application/json"`
	SourceID                 string                           `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceSftpBulkRequest) GetSourceSftpBulkPutRequest() *shared.SourceSftpBulkPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceSftpBulkPutRequest
}

func (o *PutSourceSftpBulkRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceSftpBulkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceSftpBulkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceSftpBulkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceSftpBulkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
