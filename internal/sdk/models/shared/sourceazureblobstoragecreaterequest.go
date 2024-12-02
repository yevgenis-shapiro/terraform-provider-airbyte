// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceAzureBlobStorageCreateRequest struct {
	// NOTE: When this Spec is changed, legacy_config_transformer.py must also be modified to uptake the changes
	// because it is responsible for converting legacy Azure Blob Storage v0 configs into v1 configs using the File-Based CDK.
	Configuration SourceAzureBlobStorage `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceAzureBlobStorageCreateRequest) GetConfiguration() SourceAzureBlobStorage {
	if o == nil {
		return SourceAzureBlobStorage{}
	}
	return o.Configuration
}

func (o *SourceAzureBlobStorageCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceAzureBlobStorageCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAzureBlobStorageCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceAzureBlobStorageCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
