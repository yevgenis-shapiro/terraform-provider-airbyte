// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceWhenIWorkPutRequest struct {
	Configuration SourceWhenIWorkUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourceWhenIWorkPutRequest) GetConfiguration() SourceWhenIWorkUpdate {
	if o == nil {
		return SourceWhenIWorkUpdate{}
	}
	return o.Configuration
}

func (o *SourceWhenIWorkPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceWhenIWorkPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}