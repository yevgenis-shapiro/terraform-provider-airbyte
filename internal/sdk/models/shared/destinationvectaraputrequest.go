// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationVectaraPutRequest struct {
	// Configuration to connect to the Vectara instance
	Configuration DestinationVectaraUpdate `json:"configuration"`
	Name          string                   `json:"name"`
	WorkspaceID   string                   `json:"workspaceId"`
}

func (o *DestinationVectaraPutRequest) GetConfiguration() DestinationVectaraUpdate {
	if o == nil {
		return DestinationVectaraUpdate{}
	}
	return o.Configuration
}

func (o *DestinationVectaraPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationVectaraPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
