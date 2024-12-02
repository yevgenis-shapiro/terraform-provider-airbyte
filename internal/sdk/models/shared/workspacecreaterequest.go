// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WorkspaceCreateRequest struct {
	// Name of the workspace
	Name string `json:"name"`
	// ID of organization to add workspace to.
	OrganizationID *string `json:"organizationId,omitempty"`
}

func (o *WorkspaceCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WorkspaceCreateRequest) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}