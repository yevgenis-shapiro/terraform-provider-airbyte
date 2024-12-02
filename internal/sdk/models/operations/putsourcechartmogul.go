// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceChartmogulRequest struct {
	SourceChartmogulPutRequest *shared.SourceChartmogulPutRequest `request:"mediaType=application/json"`
	SourceID                   string                             `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceChartmogulRequest) GetSourceChartmogulPutRequest() *shared.SourceChartmogulPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceChartmogulPutRequest
}

func (o *PutSourceChartmogulRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceChartmogulResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceChartmogulResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceChartmogulResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceChartmogulResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
