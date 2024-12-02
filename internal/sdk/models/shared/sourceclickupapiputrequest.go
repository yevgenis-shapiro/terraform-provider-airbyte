// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceClickupAPIPutRequest struct {
	Configuration SourceClickupAPIUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceClickupAPIPutRequest) GetConfiguration() SourceClickupAPIUpdate {
	if o == nil {
		return SourceClickupAPIUpdate{}
	}
	return o.Configuration
}

func (o *SourceClickupAPIPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceClickupAPIPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}