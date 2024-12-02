// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceMongodbV2UpdateSchemasClusterType string

const (
	SourceMongodbV2UpdateSchemasClusterTypeSelfManagedReplicaSet SourceMongodbV2UpdateSchemasClusterType = "SELF_MANAGED_REPLICA_SET"
)

func (e SourceMongodbV2UpdateSchemasClusterType) ToPointer() *SourceMongodbV2UpdateSchemasClusterType {
	return &e
}
func (e *SourceMongodbV2UpdateSchemasClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SELF_MANAGED_REPLICA_SET":
		*e = SourceMongodbV2UpdateSchemasClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2UpdateSchemasClusterType: %v", v)
	}
}

// SelfManagedReplicaSet - MongoDB self-hosted cluster configured as a replica set
type SelfManagedReplicaSet struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The authentication source where the user information is stored.
	AuthSource  *string                                 `default:"admin" json:"auth_source"`
	clusterType SourceMongodbV2UpdateSchemasClusterType `const:"SELF_MANAGED_REPLICA_SET" json:"cluster_type"`
	// The connection string of the cluster that you want to replicate.  https://www.mongodb.com/docs/manual/reference/connection-string/#find-your-self-hosted-deployment-s-connection-string for more information.
	ConnectionString string `json:"connection_string"`
	// The name of the MongoDB database that contains the collection(s) to replicate.
	Database string `json:"database"`
	// The password associated with this username.
	Password *string `json:"password,omitempty"`
	// When enabled, syncs will validate and structure records against the stream's schema.
	SchemaEnforced *bool `default:"true" json:"schema_enforced"`
	// The username which is used to access the database.
	Username *string `json:"username,omitempty"`
}

func (s SelfManagedReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SelfManagedReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SelfManagedReplicaSet) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SelfManagedReplicaSet) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *SelfManagedReplicaSet) GetClusterType() SourceMongodbV2UpdateSchemasClusterType {
	return SourceMongodbV2UpdateSchemasClusterTypeSelfManagedReplicaSet
}

func (o *SelfManagedReplicaSet) GetConnectionString() string {
	if o == nil {
		return ""
	}
	return o.ConnectionString
}

func (o *SelfManagedReplicaSet) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SelfManagedReplicaSet) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SelfManagedReplicaSet) GetSchemaEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaEnforced
}

func (o *SelfManagedReplicaSet) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type SourceMongodbV2UpdateClusterType string

const (
	SourceMongodbV2UpdateClusterTypeAtlasReplicaSet SourceMongodbV2UpdateClusterType = "ATLAS_REPLICA_SET"
)

func (e SourceMongodbV2UpdateClusterType) ToPointer() *SourceMongodbV2UpdateClusterType {
	return &e
}
func (e *SourceMongodbV2UpdateClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ATLAS_REPLICA_SET":
		*e = SourceMongodbV2UpdateClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2UpdateClusterType: %v", v)
	}
}

// MongoDBAtlasReplicaSet - MongoDB Atlas-hosted cluster configured as a replica set
type MongoDBAtlasReplicaSet struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The authentication source where the user information is stored.  See https://www.mongodb.com/docs/manual/reference/connection-string/#mongodb-urioption-urioption.authSource for more details.
	AuthSource  *string                          `default:"admin" json:"auth_source"`
	clusterType SourceMongodbV2UpdateClusterType `const:"ATLAS_REPLICA_SET" json:"cluster_type"`
	// The connection string of the cluster that you want to replicate.
	ConnectionString string `json:"connection_string"`
	// The name of the MongoDB database that contains the collection(s) to replicate.
	Database string `json:"database"`
	// The password associated with this username.
	Password string `json:"password"`
	// When enabled, syncs will validate and structure records against the stream's schema.
	SchemaEnforced *bool `default:"true" json:"schema_enforced"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (m MongoDBAtlasReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MongoDBAtlasReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MongoDBAtlasReplicaSet) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *MongoDBAtlasReplicaSet) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *MongoDBAtlasReplicaSet) GetClusterType() SourceMongodbV2UpdateClusterType {
	return SourceMongodbV2UpdateClusterTypeAtlasReplicaSet
}

func (o *MongoDBAtlasReplicaSet) GetConnectionString() string {
	if o == nil {
		return ""
	}
	return o.ConnectionString
}

func (o *MongoDBAtlasReplicaSet) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *MongoDBAtlasReplicaSet) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *MongoDBAtlasReplicaSet) GetSchemaEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaEnforced
}

func (o *MongoDBAtlasReplicaSet) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type ClusterTypeType string

const (
	ClusterTypeTypeMongoDBAtlasReplicaSet ClusterTypeType = "MongoDB Atlas Replica Set"
	ClusterTypeTypeSelfManagedReplicaSet  ClusterTypeType = "Self-Managed Replica Set"
)

// ClusterType - Configures the MongoDB cluster type.
type ClusterType struct {
	MongoDBAtlasReplicaSet *MongoDBAtlasReplicaSet
	SelfManagedReplicaSet  *SelfManagedReplicaSet

	Type ClusterTypeType
}

func CreateClusterTypeMongoDBAtlasReplicaSet(mongoDBAtlasReplicaSet MongoDBAtlasReplicaSet) ClusterType {
	typ := ClusterTypeTypeMongoDBAtlasReplicaSet

	return ClusterType{
		MongoDBAtlasReplicaSet: &mongoDBAtlasReplicaSet,
		Type:                   typ,
	}
}

func CreateClusterTypeSelfManagedReplicaSet(selfManagedReplicaSet SelfManagedReplicaSet) ClusterType {
	typ := ClusterTypeTypeSelfManagedReplicaSet

	return ClusterType{
		SelfManagedReplicaSet: &selfManagedReplicaSet,
		Type:                  typ,
	}
}

func (u *ClusterType) UnmarshalJSON(data []byte) error {

	var mongoDBAtlasReplicaSet MongoDBAtlasReplicaSet = MongoDBAtlasReplicaSet{}
	if err := utils.UnmarshalJSON(data, &mongoDBAtlasReplicaSet, "", true, true); err == nil {
		u.MongoDBAtlasReplicaSet = &mongoDBAtlasReplicaSet
		u.Type = ClusterTypeTypeMongoDBAtlasReplicaSet
		return nil
	}

	var selfManagedReplicaSet SelfManagedReplicaSet = SelfManagedReplicaSet{}
	if err := utils.UnmarshalJSON(data, &selfManagedReplicaSet, "", true, true); err == nil {
		u.SelfManagedReplicaSet = &selfManagedReplicaSet
		u.Type = ClusterTypeTypeSelfManagedReplicaSet
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ClusterType", string(data))
}

func (u ClusterType) MarshalJSON() ([]byte, error) {
	if u.MongoDBAtlasReplicaSet != nil {
		return utils.MarshalJSON(u.MongoDBAtlasReplicaSet, "", true)
	}

	if u.SelfManagedReplicaSet != nil {
		return utils.MarshalJSON(u.SelfManagedReplicaSet, "", true)
	}

	return nil, errors.New("could not marshal union type ClusterType: all fields are null")
}

// InvalidCDCPositionBehaviorAdvanced - Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss.
type InvalidCDCPositionBehaviorAdvanced string

const (
	InvalidCDCPositionBehaviorAdvancedFailSync   InvalidCDCPositionBehaviorAdvanced = "Fail sync"
	InvalidCDCPositionBehaviorAdvancedReSyncData InvalidCDCPositionBehaviorAdvanced = "Re-sync data"
)

func (e InvalidCDCPositionBehaviorAdvanced) ToPointer() *InvalidCDCPositionBehaviorAdvanced {
	return &e
}
func (e *InvalidCDCPositionBehaviorAdvanced) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Fail sync":
		fallthrough
	case "Re-sync data":
		*e = InvalidCDCPositionBehaviorAdvanced(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvalidCDCPositionBehaviorAdvanced: %v", v)
	}
}

// CaptureModeAdvanced - Determines how Airbyte looks up the value of an updated document. If 'Lookup' is chosen, the current value of the document will be read. If 'Post Image' is chosen, then the version of the document immediately after an update will be read. WARNING : Severe data loss will occur if this option is chosen and the appropriate settings are not set on your Mongo instance : https://www.mongodb.com/docs/manual/changeStreams/#change-streams-with-document-pre-and-post-images.
type CaptureModeAdvanced string

const (
	CaptureModeAdvancedLookup    CaptureModeAdvanced = "Lookup"
	CaptureModeAdvancedPostImage CaptureModeAdvanced = "Post Image"
)

func (e CaptureModeAdvanced) ToPointer() *CaptureModeAdvanced {
	return &e
}
func (e *CaptureModeAdvanced) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Lookup":
		fallthrough
	case "Post Image":
		*e = CaptureModeAdvanced(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CaptureModeAdvanced: %v", v)
	}
}

type SourceMongodbV2Update struct {
	// Configures the MongoDB cluster type.
	DatabaseConfig ClusterType `json:"database_config"`
	// The maximum number of documents to sample when attempting to discover the unique fields for a collection.
	DiscoverSampleSize *int64 `default:"10000" json:"discover_sample_size"`
	// The amount of time an initial load is allowed to continue for before catching up on CDC logs.
	InitialLoadTimeoutHours *int64 `default:"8" json:"initial_load_timeout_hours"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds.
	InitialWaitingSeconds *int64 `default:"300" json:"initial_waiting_seconds"`
	// Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss.
	InvalidCdcCursorPositionBehavior *InvalidCDCPositionBehaviorAdvanced `default:"Fail sync" json:"invalid_cdc_cursor_position_behavior"`
	// The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.
	QueueSize *int64 `default:"10000" json:"queue_size"`
	// Determines how Airbyte looks up the value of an updated document. If 'Lookup' is chosen, the current value of the document will be read. If 'Post Image' is chosen, then the version of the document immediately after an update will be read. WARNING : Severe data loss will occur if this option is chosen and the appropriate settings are not set on your Mongo instance : https://www.mongodb.com/docs/manual/changeStreams/#change-streams-with-document-pre-and-post-images.
	UpdateCaptureMode *CaptureModeAdvanced `default:"Lookup" json:"update_capture_mode"`
}

func (s SourceMongodbV2Update) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbV2Update) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbV2Update) GetDatabaseConfig() ClusterType {
	if o == nil {
		return ClusterType{}
	}
	return o.DatabaseConfig
}

func (o *SourceMongodbV2Update) GetDiscoverSampleSize() *int64 {
	if o == nil {
		return nil
	}
	return o.DiscoverSampleSize
}

func (o *SourceMongodbV2Update) GetInitialLoadTimeoutHours() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialLoadTimeoutHours
}

func (o *SourceMongodbV2Update) GetInitialWaitingSeconds() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialWaitingSeconds
}

func (o *SourceMongodbV2Update) GetInvalidCdcCursorPositionBehavior() *InvalidCDCPositionBehaviorAdvanced {
	if o == nil {
		return nil
	}
	return o.InvalidCdcCursorPositionBehavior
}

func (o *SourceMongodbV2Update) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *SourceMongodbV2Update) GetUpdateCaptureMode() *CaptureModeAdvanced {
	if o == nil {
		return nil
	}
	return o.UpdateCaptureMode
}
