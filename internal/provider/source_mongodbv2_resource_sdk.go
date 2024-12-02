// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMongodbV2ResourceModel) ToSharedSourceMongodbV2CreateRequest() *shared.SourceMongodbV2CreateRequest {
	var databaseConfig shared.SourceMongodbV2ClusterType
	var sourceMongodbV2MongoDBAtlasReplicaSet *shared.SourceMongodbV2MongoDBAtlasReplicaSet
	if r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet != nil {
		var additionalProperties interface{}
		if !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AdditionalProperties.IsUnknown() && !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AdditionalProperties.ValueString()), &additionalProperties)
		}
		authSource := new(string)
		if !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AuthSource.IsUnknown() && !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AuthSource.IsNull() {
			*authSource = r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AuthSource.ValueString()
		} else {
			authSource = nil
		}
		connectionString := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.ConnectionString.ValueString()
		database := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.Database.ValueString()
		password := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.Password.ValueString()
		schemaEnforced := new(bool)
		if !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.SchemaEnforced.IsUnknown() && !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.SchemaEnforced.IsNull() {
			*schemaEnforced = r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.SchemaEnforced.ValueBool()
		} else {
			schemaEnforced = nil
		}
		username := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.Username.ValueString()
		sourceMongodbV2MongoDBAtlasReplicaSet = &shared.SourceMongodbV2MongoDBAtlasReplicaSet{
			AdditionalProperties: additionalProperties,
			AuthSource:           authSource,
			ConnectionString:     connectionString,
			Database:             database,
			Password:             password,
			SchemaEnforced:       schemaEnforced,
			Username:             username,
		}
	}
	if sourceMongodbV2MongoDBAtlasReplicaSet != nil {
		databaseConfig = shared.SourceMongodbV2ClusterType{
			SourceMongodbV2MongoDBAtlasReplicaSet: sourceMongodbV2MongoDBAtlasReplicaSet,
		}
	}
	var sourceMongodbV2SelfManagedReplicaSet *shared.SourceMongodbV2SelfManagedReplicaSet
	if r.Configuration.DatabaseConfig.SelfManagedReplicaSet != nil {
		var additionalProperties1 interface{}
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AdditionalProperties.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AdditionalProperties.ValueString()), &additionalProperties1)
		}
		authSource1 := new(string)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AuthSource.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AuthSource.IsNull() {
			*authSource1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AuthSource.ValueString()
		} else {
			authSource1 = nil
		}
		connectionString1 := r.Configuration.DatabaseConfig.SelfManagedReplicaSet.ConnectionString.ValueString()
		database1 := r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Database.ValueString()
		password1 := new(string)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Password.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Password.IsNull() {
			*password1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Password.ValueString()
		} else {
			password1 = nil
		}
		schemaEnforced1 := new(bool)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.SchemaEnforced.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.SchemaEnforced.IsNull() {
			*schemaEnforced1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.SchemaEnforced.ValueBool()
		} else {
			schemaEnforced1 = nil
		}
		username1 := new(string)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Username.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Username.IsNull() {
			*username1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Username.ValueString()
		} else {
			username1 = nil
		}
		sourceMongodbV2SelfManagedReplicaSet = &shared.SourceMongodbV2SelfManagedReplicaSet{
			AdditionalProperties: additionalProperties1,
			AuthSource:           authSource1,
			ConnectionString:     connectionString1,
			Database:             database1,
			Password:             password1,
			SchemaEnforced:       schemaEnforced1,
			Username:             username1,
		}
	}
	if sourceMongodbV2SelfManagedReplicaSet != nil {
		databaseConfig = shared.SourceMongodbV2ClusterType{
			SourceMongodbV2SelfManagedReplicaSet: sourceMongodbV2SelfManagedReplicaSet,
		}
	}
	discoverSampleSize := new(int64)
	if !r.Configuration.DiscoverSampleSize.IsUnknown() && !r.Configuration.DiscoverSampleSize.IsNull() {
		*discoverSampleSize = r.Configuration.DiscoverSampleSize.ValueInt64()
	} else {
		discoverSampleSize = nil
	}
	initialLoadTimeoutHours := new(int64)
	if !r.Configuration.InitialLoadTimeoutHours.IsUnknown() && !r.Configuration.InitialLoadTimeoutHours.IsNull() {
		*initialLoadTimeoutHours = r.Configuration.InitialLoadTimeoutHours.ValueInt64()
	} else {
		initialLoadTimeoutHours = nil
	}
	initialWaitingSeconds := new(int64)
	if !r.Configuration.InitialWaitingSeconds.IsUnknown() && !r.Configuration.InitialWaitingSeconds.IsNull() {
		*initialWaitingSeconds = r.Configuration.InitialWaitingSeconds.ValueInt64()
	} else {
		initialWaitingSeconds = nil
	}
	invalidCdcCursorPositionBehavior := new(shared.SourceMongodbV2InvalidCDCPositionBehaviorAdvanced)
	if !r.Configuration.InvalidCdcCursorPositionBehavior.IsUnknown() && !r.Configuration.InvalidCdcCursorPositionBehavior.IsNull() {
		*invalidCdcCursorPositionBehavior = shared.SourceMongodbV2InvalidCDCPositionBehaviorAdvanced(r.Configuration.InvalidCdcCursorPositionBehavior.ValueString())
	} else {
		invalidCdcCursorPositionBehavior = nil
	}
	queueSize := new(int64)
	if !r.Configuration.QueueSize.IsUnknown() && !r.Configuration.QueueSize.IsNull() {
		*queueSize = r.Configuration.QueueSize.ValueInt64()
	} else {
		queueSize = nil
	}
	updateCaptureMode := new(shared.SourceMongodbV2CaptureModeAdvanced)
	if !r.Configuration.UpdateCaptureMode.IsUnknown() && !r.Configuration.UpdateCaptureMode.IsNull() {
		*updateCaptureMode = shared.SourceMongodbV2CaptureModeAdvanced(r.Configuration.UpdateCaptureMode.ValueString())
	} else {
		updateCaptureMode = nil
	}
	configuration := shared.SourceMongodbV2{
		DatabaseConfig:                   databaseConfig,
		DiscoverSampleSize:               discoverSampleSize,
		InitialLoadTimeoutHours:          initialLoadTimeoutHours,
		InitialWaitingSeconds:            initialWaitingSeconds,
		InvalidCdcCursorPositionBehavior: invalidCdcCursorPositionBehavior,
		QueueSize:                        queueSize,
		UpdateCaptureMode:                updateCaptureMode,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMongodbV2CreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMongodbV2ResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceMongodbV2ResourceModel) ToSharedSourceMongodbV2PutRequest() *shared.SourceMongodbV2PutRequest {
	var databaseConfig shared.ClusterType
	var mongoDBAtlasReplicaSet *shared.MongoDBAtlasReplicaSet
	if r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet != nil {
		var additionalProperties interface{}
		if !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AdditionalProperties.IsUnknown() && !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AdditionalProperties.ValueString()), &additionalProperties)
		}
		authSource := new(string)
		if !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AuthSource.IsUnknown() && !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AuthSource.IsNull() {
			*authSource = r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.AuthSource.ValueString()
		} else {
			authSource = nil
		}
		connectionString := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.ConnectionString.ValueString()
		database := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.Database.ValueString()
		password := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.Password.ValueString()
		schemaEnforced := new(bool)
		if !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.SchemaEnforced.IsUnknown() && !r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.SchemaEnforced.IsNull() {
			*schemaEnforced = r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.SchemaEnforced.ValueBool()
		} else {
			schemaEnforced = nil
		}
		username := r.Configuration.DatabaseConfig.MongoDBAtlasReplicaSet.Username.ValueString()
		mongoDBAtlasReplicaSet = &shared.MongoDBAtlasReplicaSet{
			AdditionalProperties: additionalProperties,
			AuthSource:           authSource,
			ConnectionString:     connectionString,
			Database:             database,
			Password:             password,
			SchemaEnforced:       schemaEnforced,
			Username:             username,
		}
	}
	if mongoDBAtlasReplicaSet != nil {
		databaseConfig = shared.ClusterType{
			MongoDBAtlasReplicaSet: mongoDBAtlasReplicaSet,
		}
	}
	var selfManagedReplicaSet *shared.SelfManagedReplicaSet
	if r.Configuration.DatabaseConfig.SelfManagedReplicaSet != nil {
		var additionalProperties1 interface{}
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AdditionalProperties.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AdditionalProperties.ValueString()), &additionalProperties1)
		}
		authSource1 := new(string)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AuthSource.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AuthSource.IsNull() {
			*authSource1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.AuthSource.ValueString()
		} else {
			authSource1 = nil
		}
		connectionString1 := r.Configuration.DatabaseConfig.SelfManagedReplicaSet.ConnectionString.ValueString()
		database1 := r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Database.ValueString()
		password1 := new(string)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Password.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Password.IsNull() {
			*password1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Password.ValueString()
		} else {
			password1 = nil
		}
		schemaEnforced1 := new(bool)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.SchemaEnforced.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.SchemaEnforced.IsNull() {
			*schemaEnforced1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.SchemaEnforced.ValueBool()
		} else {
			schemaEnforced1 = nil
		}
		username1 := new(string)
		if !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Username.IsUnknown() && !r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Username.IsNull() {
			*username1 = r.Configuration.DatabaseConfig.SelfManagedReplicaSet.Username.ValueString()
		} else {
			username1 = nil
		}
		selfManagedReplicaSet = &shared.SelfManagedReplicaSet{
			AdditionalProperties: additionalProperties1,
			AuthSource:           authSource1,
			ConnectionString:     connectionString1,
			Database:             database1,
			Password:             password1,
			SchemaEnforced:       schemaEnforced1,
			Username:             username1,
		}
	}
	if selfManagedReplicaSet != nil {
		databaseConfig = shared.ClusterType{
			SelfManagedReplicaSet: selfManagedReplicaSet,
		}
	}
	discoverSampleSize := new(int64)
	if !r.Configuration.DiscoverSampleSize.IsUnknown() && !r.Configuration.DiscoverSampleSize.IsNull() {
		*discoverSampleSize = r.Configuration.DiscoverSampleSize.ValueInt64()
	} else {
		discoverSampleSize = nil
	}
	initialLoadTimeoutHours := new(int64)
	if !r.Configuration.InitialLoadTimeoutHours.IsUnknown() && !r.Configuration.InitialLoadTimeoutHours.IsNull() {
		*initialLoadTimeoutHours = r.Configuration.InitialLoadTimeoutHours.ValueInt64()
	} else {
		initialLoadTimeoutHours = nil
	}
	initialWaitingSeconds := new(int64)
	if !r.Configuration.InitialWaitingSeconds.IsUnknown() && !r.Configuration.InitialWaitingSeconds.IsNull() {
		*initialWaitingSeconds = r.Configuration.InitialWaitingSeconds.ValueInt64()
	} else {
		initialWaitingSeconds = nil
	}
	invalidCdcCursorPositionBehavior := new(shared.InvalidCDCPositionBehaviorAdvanced)
	if !r.Configuration.InvalidCdcCursorPositionBehavior.IsUnknown() && !r.Configuration.InvalidCdcCursorPositionBehavior.IsNull() {
		*invalidCdcCursorPositionBehavior = shared.InvalidCDCPositionBehaviorAdvanced(r.Configuration.InvalidCdcCursorPositionBehavior.ValueString())
	} else {
		invalidCdcCursorPositionBehavior = nil
	}
	queueSize := new(int64)
	if !r.Configuration.QueueSize.IsUnknown() && !r.Configuration.QueueSize.IsNull() {
		*queueSize = r.Configuration.QueueSize.ValueInt64()
	} else {
		queueSize = nil
	}
	updateCaptureMode := new(shared.CaptureModeAdvanced)
	if !r.Configuration.UpdateCaptureMode.IsUnknown() && !r.Configuration.UpdateCaptureMode.IsNull() {
		*updateCaptureMode = shared.CaptureModeAdvanced(r.Configuration.UpdateCaptureMode.ValueString())
	} else {
		updateCaptureMode = nil
	}
	configuration := shared.SourceMongodbV2Update{
		DatabaseConfig:                   databaseConfig,
		DiscoverSampleSize:               discoverSampleSize,
		InitialLoadTimeoutHours:          initialLoadTimeoutHours,
		InitialWaitingSeconds:            initialWaitingSeconds,
		InvalidCdcCursorPositionBehavior: invalidCdcCursorPositionBehavior,
		QueueSize:                        queueSize,
		UpdateCaptureMode:                updateCaptureMode,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMongodbV2PutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}