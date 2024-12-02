// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceClickhouseResourceModel) ToSharedSourceClickhouseCreateRequest() *shared.SourceClickhouseCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	ssl := new(bool)
	if !r.Configuration.Ssl.IsUnknown() && !r.Configuration.Ssl.IsNull() {
		*ssl = r.Configuration.Ssl.ValueBool()
	} else {
		ssl = nil
	}
	var tunnelMethod *shared.SourceClickhouseSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var sourceClickhouseNoTunnel *shared.SourceClickhouseNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			sourceClickhouseNoTunnel = &shared.SourceClickhouseNoTunnel{}
		}
		if sourceClickhouseNoTunnel != nil {
			tunnelMethod = &shared.SourceClickhouseSSHTunnelMethod{
				SourceClickhouseNoTunnel: sourceClickhouseNoTunnel,
			}
		}
		var sourceClickhouseSSHKeyAuthentication *shared.SourceClickhouseSSHKeyAuthentication
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
			sourceClickhouseSSHKeyAuthentication = &shared.SourceClickhouseSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sourceClickhouseSSHKeyAuthentication != nil {
			tunnelMethod = &shared.SourceClickhouseSSHTunnelMethod{
				SourceClickhouseSSHKeyAuthentication: sourceClickhouseSSHKeyAuthentication,
			}
		}
		var sourceClickhousePasswordAuthentication *shared.SourceClickhousePasswordAuthentication
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
			sourceClickhousePasswordAuthentication = &shared.SourceClickhousePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if sourceClickhousePasswordAuthentication != nil {
			tunnelMethod = &shared.SourceClickhouseSSHTunnelMethod{
				SourceClickhousePasswordAuthentication: sourceClickhousePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceClickhouse{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Ssl:           ssl,
		TunnelMethod:  tunnelMethod,
		Username:      username,
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
	out := shared.SourceClickhouseCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceClickhouseResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceClickhouseResourceModel) ToSharedSourceClickhousePutRequest() *shared.SourceClickhousePutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	ssl := new(bool)
	if !r.Configuration.Ssl.IsUnknown() && !r.Configuration.Ssl.IsNull() {
		*ssl = r.Configuration.Ssl.ValueBool()
	} else {
		ssl = nil
	}
	var tunnelMethod *shared.SourceClickhouseUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var sourceClickhouseUpdateNoTunnel *shared.SourceClickhouseUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			sourceClickhouseUpdateNoTunnel = &shared.SourceClickhouseUpdateNoTunnel{}
		}
		if sourceClickhouseUpdateNoTunnel != nil {
			tunnelMethod = &shared.SourceClickhouseUpdateSSHTunnelMethod{
				SourceClickhouseUpdateNoTunnel: sourceClickhouseUpdateNoTunnel,
			}
		}
		var sourceClickhouseUpdateSSHKeyAuthentication *shared.SourceClickhouseUpdateSSHKeyAuthentication
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
			sourceClickhouseUpdateSSHKeyAuthentication = &shared.SourceClickhouseUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sourceClickhouseUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.SourceClickhouseUpdateSSHTunnelMethod{
				SourceClickhouseUpdateSSHKeyAuthentication: sourceClickhouseUpdateSSHKeyAuthentication,
			}
		}
		var sourceClickhouseUpdatePasswordAuthentication *shared.SourceClickhouseUpdatePasswordAuthentication
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
			sourceClickhouseUpdatePasswordAuthentication = &shared.SourceClickhouseUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if sourceClickhouseUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.SourceClickhouseUpdateSSHTunnelMethod{
				SourceClickhouseUpdatePasswordAuthentication: sourceClickhouseUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceClickhouseUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Ssl:           ssl,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceClickhousePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
