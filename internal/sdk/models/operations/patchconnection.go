// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PatchConnectionRequest struct {
	ConnectionPatchRequest shared.ConnectionPatchRequest `request:"mediaType=application/json"`
	ConnectionID           string                        `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *PatchConnectionRequest) GetConnectionPatchRequest() shared.ConnectionPatchRequest {
	if o == nil {
		return shared.ConnectionPatchRequest{}
	}
	return o.ConnectionPatchRequest
}

func (o *PatchConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PatchConnectionResponse struct {
	// Update a Connection by the id in the path.
	ConnectionResponse *shared.ConnectionResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchConnectionResponse) GetConnectionResponse() *shared.ConnectionResponse {
	if o == nil {
		return nil
	}
	return o.ConnectionResponse
}

func (o *PatchConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
