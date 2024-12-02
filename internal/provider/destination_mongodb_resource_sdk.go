// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMongodbResourceModel) ToSharedDestinationMongodbCreateRequest() *shared.DestinationMongodbCreateRequest {
	var authType shared.DestinationMongodbAuthorizationType
	var destinationMongodbNone *shared.DestinationMongodbNone
	if r.Configuration.AuthType.None != nil {
		destinationMongodbNone = &shared.DestinationMongodbNone{}
	}
	if destinationMongodbNone != nil {
		authType = shared.DestinationMongodbAuthorizationType{
			DestinationMongodbNone: destinationMongodbNone,
		}
	}
	var destinationMongodbLoginPassword *shared.DestinationMongodbLoginPassword
	if r.Configuration.AuthType.LoginPassword != nil {
		password := r.Configuration.AuthType.LoginPassword.Password.ValueString()
		username := r.Configuration.AuthType.LoginPassword.Username.ValueString()
		destinationMongodbLoginPassword = &shared.DestinationMongodbLoginPassword{
			Password: password,
			Username: username,
		}
	}
	if destinationMongodbLoginPassword != nil {
		authType = shared.DestinationMongodbAuthorizationType{
			DestinationMongodbLoginPassword: destinationMongodbLoginPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	var instanceType *shared.DestinationMongodbMongoDbInstanceType
	if r.Configuration.InstanceType != nil {
		var destinationMongodbStandaloneMongoDbInstance *shared.DestinationMongodbStandaloneMongoDbInstance
		if r.Configuration.InstanceType.StandaloneMongoDbInstance != nil {
			host := r.Configuration.InstanceType.StandaloneMongoDbInstance.Host.ValueString()
			instance := new(shared.DestinationMongodbInstance)
			if !r.Configuration.InstanceType.StandaloneMongoDbInstance.Instance.IsUnknown() && !r.Configuration.InstanceType.StandaloneMongoDbInstance.Instance.IsNull() {
				*instance = shared.DestinationMongodbInstance(r.Configuration.InstanceType.StandaloneMongoDbInstance.Instance.ValueString())
			} else {
				instance = nil
			}
			port := new(int64)
			if !r.Configuration.InstanceType.StandaloneMongoDbInstance.Port.IsUnknown() && !r.Configuration.InstanceType.StandaloneMongoDbInstance.Port.IsNull() {
				*port = r.Configuration.InstanceType.StandaloneMongoDbInstance.Port.ValueInt64()
			} else {
				port = nil
			}
			tls := new(bool)
			if !r.Configuration.InstanceType.StandaloneMongoDbInstance.TLS.IsUnknown() && !r.Configuration.InstanceType.StandaloneMongoDbInstance.TLS.IsNull() {
				*tls = r.Configuration.InstanceType.StandaloneMongoDbInstance.TLS.ValueBool()
			} else {
				tls = nil
			}
			destinationMongodbStandaloneMongoDbInstance = &shared.DestinationMongodbStandaloneMongoDbInstance{
				Host:     host,
				Instance: instance,
				Port:     port,
				TLS:      tls,
			}
		}
		if destinationMongodbStandaloneMongoDbInstance != nil {
			instanceType = &shared.DestinationMongodbMongoDbInstanceType{
				DestinationMongodbStandaloneMongoDbInstance: destinationMongodbStandaloneMongoDbInstance,
			}
		}
		var destinationMongodbReplicaSet *shared.DestinationMongodbReplicaSet
		if r.Configuration.InstanceType.ReplicaSet != nil {
			instance1 := new(shared.DestinationMongodbSchemasInstance)
			if !r.Configuration.InstanceType.ReplicaSet.Instance.IsUnknown() && !r.Configuration.InstanceType.ReplicaSet.Instance.IsNull() {
				*instance1 = shared.DestinationMongodbSchemasInstance(r.Configuration.InstanceType.ReplicaSet.Instance.ValueString())
			} else {
				instance1 = nil
			}
			replicaSet := new(string)
			if !r.Configuration.InstanceType.ReplicaSet.ReplicaSet.IsUnknown() && !r.Configuration.InstanceType.ReplicaSet.ReplicaSet.IsNull() {
				*replicaSet = r.Configuration.InstanceType.ReplicaSet.ReplicaSet.ValueString()
			} else {
				replicaSet = nil
			}
			serverAddresses := r.Configuration.InstanceType.ReplicaSet.ServerAddresses.ValueString()
			destinationMongodbReplicaSet = &shared.DestinationMongodbReplicaSet{
				Instance:        instance1,
				ReplicaSet:      replicaSet,
				ServerAddresses: serverAddresses,
			}
		}
		if destinationMongodbReplicaSet != nil {
			instanceType = &shared.DestinationMongodbMongoDbInstanceType{
				DestinationMongodbReplicaSet: destinationMongodbReplicaSet,
			}
		}
		var destinationMongodbMongoDBAtlas *shared.DestinationMongodbMongoDBAtlas
		if r.Configuration.InstanceType.MongoDBAtlas != nil {
			clusterURL := r.Configuration.InstanceType.MongoDBAtlas.ClusterURL.ValueString()
			instance2 := new(shared.DestinationMongodbSchemasInstanceTypeInstance)
			if !r.Configuration.InstanceType.MongoDBAtlas.Instance.IsUnknown() && !r.Configuration.InstanceType.MongoDBAtlas.Instance.IsNull() {
				*instance2 = shared.DestinationMongodbSchemasInstanceTypeInstance(r.Configuration.InstanceType.MongoDBAtlas.Instance.ValueString())
			} else {
				instance2 = nil
			}
			destinationMongodbMongoDBAtlas = &shared.DestinationMongodbMongoDBAtlas{
				ClusterURL: clusterURL,
				Instance:   instance2,
			}
		}
		if destinationMongodbMongoDBAtlas != nil {
			instanceType = &shared.DestinationMongodbMongoDbInstanceType{
				DestinationMongodbMongoDBAtlas: destinationMongodbMongoDBAtlas,
			}
		}
	}
	var tunnelMethod *shared.DestinationMongodbSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMongodbNoTunnel *shared.DestinationMongodbNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationMongodbNoTunnel = &shared.DestinationMongodbNoTunnel{}
		}
		if destinationMongodbNoTunnel != nil {
			tunnelMethod = &shared.DestinationMongodbSSHTunnelMethod{
				DestinationMongodbNoTunnel: destinationMongodbNoTunnel,
			}
		}
		var destinationMongodbSSHKeyAuthentication *shared.DestinationMongodbSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationMongodbSSHKeyAuthentication = &shared.DestinationMongodbSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationMongodbSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMongodbSSHTunnelMethod{
				DestinationMongodbSSHKeyAuthentication: destinationMongodbSSHKeyAuthentication,
			}
		}
		var destinationMongodbPasswordAuthentication *shared.DestinationMongodbPasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationMongodbPasswordAuthentication = &shared.DestinationMongodbPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMongodbPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMongodbSSHTunnelMethod{
				DestinationMongodbPasswordAuthentication: destinationMongodbPasswordAuthentication,
			}
		}
	}
	configuration := shared.DestinationMongodb{
		AuthType:     authType,
		Database:     database,
		InstanceType: instanceType,
		TunnelMethod: tunnelMethod,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMongodbCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMongodbResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationMongodbResourceModel) ToSharedDestinationMongodbPutRequest() *shared.DestinationMongodbPutRequest {
	var authType shared.AuthorizationType
	var destinationMongodbUpdateNone *shared.DestinationMongodbUpdateNone
	if r.Configuration.AuthType.None != nil {
		destinationMongodbUpdateNone = &shared.DestinationMongodbUpdateNone{}
	}
	if destinationMongodbUpdateNone != nil {
		authType = shared.AuthorizationType{
			DestinationMongodbUpdateNone: destinationMongodbUpdateNone,
		}
	}
	var loginPassword *shared.LoginPassword
	if r.Configuration.AuthType.LoginPassword != nil {
		password := r.Configuration.AuthType.LoginPassword.Password.ValueString()
		username := r.Configuration.AuthType.LoginPassword.Username.ValueString()
		loginPassword = &shared.LoginPassword{
			Password: password,
			Username: username,
		}
	}
	if loginPassword != nil {
		authType = shared.AuthorizationType{
			LoginPassword: loginPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	var instanceType *shared.MongoDbInstanceType
	if r.Configuration.InstanceType != nil {
		var standaloneMongoDbInstance *shared.StandaloneMongoDbInstance
		if r.Configuration.InstanceType.StandaloneMongoDbInstance != nil {
			host := r.Configuration.InstanceType.StandaloneMongoDbInstance.Host.ValueString()
			instance := new(shared.Instance)
			if !r.Configuration.InstanceType.StandaloneMongoDbInstance.Instance.IsUnknown() && !r.Configuration.InstanceType.StandaloneMongoDbInstance.Instance.IsNull() {
				*instance = shared.Instance(r.Configuration.InstanceType.StandaloneMongoDbInstance.Instance.ValueString())
			} else {
				instance = nil
			}
			port := new(int64)
			if !r.Configuration.InstanceType.StandaloneMongoDbInstance.Port.IsUnknown() && !r.Configuration.InstanceType.StandaloneMongoDbInstance.Port.IsNull() {
				*port = r.Configuration.InstanceType.StandaloneMongoDbInstance.Port.ValueInt64()
			} else {
				port = nil
			}
			tls := new(bool)
			if !r.Configuration.InstanceType.StandaloneMongoDbInstance.TLS.IsUnknown() && !r.Configuration.InstanceType.StandaloneMongoDbInstance.TLS.IsNull() {
				*tls = r.Configuration.InstanceType.StandaloneMongoDbInstance.TLS.ValueBool()
			} else {
				tls = nil
			}
			standaloneMongoDbInstance = &shared.StandaloneMongoDbInstance{
				Host:     host,
				Instance: instance,
				Port:     port,
				TLS:      tls,
			}
		}
		if standaloneMongoDbInstance != nil {
			instanceType = &shared.MongoDbInstanceType{
				StandaloneMongoDbInstance: standaloneMongoDbInstance,
			}
		}
		var replicaSet *shared.ReplicaSet
		if r.Configuration.InstanceType.ReplicaSet != nil {
			instance1 := new(shared.DestinationMongodbUpdateInstance)
			if !r.Configuration.InstanceType.ReplicaSet.Instance.IsUnknown() && !r.Configuration.InstanceType.ReplicaSet.Instance.IsNull() {
				*instance1 = shared.DestinationMongodbUpdateInstance(r.Configuration.InstanceType.ReplicaSet.Instance.ValueString())
			} else {
				instance1 = nil
			}
			replicaSet1 := new(string)
			if !r.Configuration.InstanceType.ReplicaSet.ReplicaSet.IsUnknown() && !r.Configuration.InstanceType.ReplicaSet.ReplicaSet.IsNull() {
				*replicaSet1 = r.Configuration.InstanceType.ReplicaSet.ReplicaSet.ValueString()
			} else {
				replicaSet1 = nil
			}
			serverAddresses := r.Configuration.InstanceType.ReplicaSet.ServerAddresses.ValueString()
			replicaSet = &shared.ReplicaSet{
				Instance:        instance1,
				ReplicaSet:      replicaSet1,
				ServerAddresses: serverAddresses,
			}
		}
		if replicaSet != nil {
			instanceType = &shared.MongoDbInstanceType{
				ReplicaSet: replicaSet,
			}
		}
		var mongoDBAtlas *shared.MongoDBAtlas
		if r.Configuration.InstanceType.MongoDBAtlas != nil {
			clusterURL := r.Configuration.InstanceType.MongoDBAtlas.ClusterURL.ValueString()
			instance2 := new(shared.DestinationMongodbUpdateSchemasInstance)
			if !r.Configuration.InstanceType.MongoDBAtlas.Instance.IsUnknown() && !r.Configuration.InstanceType.MongoDBAtlas.Instance.IsNull() {
				*instance2 = shared.DestinationMongodbUpdateSchemasInstance(r.Configuration.InstanceType.MongoDBAtlas.Instance.ValueString())
			} else {
				instance2 = nil
			}
			mongoDBAtlas = &shared.MongoDBAtlas{
				ClusterURL: clusterURL,
				Instance:   instance2,
			}
		}
		if mongoDBAtlas != nil {
			instanceType = &shared.MongoDbInstanceType{
				MongoDBAtlas: mongoDBAtlas,
			}
		}
	}
	var tunnelMethod *shared.DestinationMongodbUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMongodbUpdateNoTunnel *shared.DestinationMongodbUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationMongodbUpdateNoTunnel = &shared.DestinationMongodbUpdateNoTunnel{}
		}
		if destinationMongodbUpdateNoTunnel != nil {
			tunnelMethod = &shared.DestinationMongodbUpdateSSHTunnelMethod{
				DestinationMongodbUpdateNoTunnel: destinationMongodbUpdateNoTunnel,
			}
		}
		var destinationMongodbUpdateSSHKeyAuthentication *shared.DestinationMongodbUpdateSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationMongodbUpdateSSHKeyAuthentication = &shared.DestinationMongodbUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationMongodbUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMongodbUpdateSSHTunnelMethod{
				DestinationMongodbUpdateSSHKeyAuthentication: destinationMongodbUpdateSSHKeyAuthentication,
			}
		}
		var destinationMongodbUpdatePasswordAuthentication *shared.DestinationMongodbUpdatePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationMongodbUpdatePasswordAuthentication = &shared.DestinationMongodbUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMongodbUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMongodbUpdateSSHTunnelMethod{
				DestinationMongodbUpdatePasswordAuthentication: destinationMongodbUpdatePasswordAuthentication,
			}
		}
	}
	configuration := shared.DestinationMongodbUpdate{
		AuthType:     authType,
		Database:     database,
		InstanceType: instanceType,
		TunnelMethod: tunnelMethod,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMongodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}