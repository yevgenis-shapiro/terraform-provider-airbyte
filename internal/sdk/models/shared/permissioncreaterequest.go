// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PermissionCreateRequest struct {
	OrganizationID *string `json:"organizationId,omitempty"`
	// Subset of `PermissionType` (removing `instance_admin`), could be used in public-api.
	PermissionType PublicPermissionType `json:"permissionType"`
	// Internal Airbyte user ID
	UserID      string  `json:"userId"`
	WorkspaceID *string `json:"workspaceId,omitempty"`
}

func (o *PermissionCreateRequest) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *PermissionCreateRequest) GetPermissionType() PublicPermissionType {
	if o == nil {
		return PublicPermissionType("")
	}
	return o.PermissionType
}

func (o *PermissionCreateRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *PermissionCreateRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}