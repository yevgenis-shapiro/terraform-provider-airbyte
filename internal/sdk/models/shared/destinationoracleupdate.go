// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationOracleUpdateSchemasEncryptionMethod string

const (
	DestinationOracleUpdateSchemasEncryptionMethodEncryptedVerifyCertificate DestinationOracleUpdateSchemasEncryptionMethod = "encrypted_verify_certificate"
)

func (e DestinationOracleUpdateSchemasEncryptionMethod) ToPointer() *DestinationOracleUpdateSchemasEncryptionMethod {
	return &e
}
func (e *DestinationOracleUpdateSchemasEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = DestinationOracleUpdateSchemasEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleUpdateSchemasEncryptionMethod: %v", v)
	}
}

// TLSEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type TLSEncryptedVerifyCertificate struct {
	encryptionMethod *DestinationOracleUpdateSchemasEncryptionMethod `const:"encrypted_verify_certificate" json:"encryption_method"`
	// Privacy Enhanced Mail (PEM) files are concatenated certificate containers frequently used in certificate installations.
	SslCertificate string `json:"ssl_certificate"`
}

func (t TLSEncryptedVerifyCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSEncryptedVerifyCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TLSEncryptedVerifyCertificate) GetEncryptionMethod() *DestinationOracleUpdateSchemasEncryptionMethod {
	return DestinationOracleUpdateSchemasEncryptionMethodEncryptedVerifyCertificate.ToPointer()
}

func (o *TLSEncryptedVerifyCertificate) GetSslCertificate() string {
	if o == nil {
		return ""
	}
	return o.SslCertificate
}

// EncryptionAlgorithm - This parameter defines the database encryption algorithm.
type EncryptionAlgorithm string

const (
	EncryptionAlgorithmAes256      EncryptionAlgorithm = "AES256"
	EncryptionAlgorithmRc456       EncryptionAlgorithm = "RC4_56"
	EncryptionAlgorithmThreeDes168 EncryptionAlgorithm = "3DES168"
)

func (e EncryptionAlgorithm) ToPointer() *EncryptionAlgorithm {
	return &e
}
func (e *EncryptionAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AES256":
		fallthrough
	case "RC4_56":
		fallthrough
	case "3DES168":
		*e = EncryptionAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EncryptionAlgorithm: %v", v)
	}
}

type DestinationOracleUpdateEncryptionMethod string

const (
	DestinationOracleUpdateEncryptionMethodClientNne DestinationOracleUpdateEncryptionMethod = "client_nne"
)

func (e DestinationOracleUpdateEncryptionMethod) ToPointer() *DestinationOracleUpdateEncryptionMethod {
	return &e
}
func (e *DestinationOracleUpdateEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_nne":
		*e = DestinationOracleUpdateEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleUpdateEncryptionMethod: %v", v)
	}
}

// NativeNetworkEncryptionNNE - The native network encryption gives you the ability to encrypt database connections, without the configuration overhead of TCP/IP and SSL/TLS and without the need to open and listen on different ports.
type NativeNetworkEncryptionNNE struct {
	// This parameter defines the database encryption algorithm.
	EncryptionAlgorithm *EncryptionAlgorithm                     `default:"AES256" json:"encryption_algorithm"`
	encryptionMethod    *DestinationOracleUpdateEncryptionMethod `const:"client_nne" json:"encryption_method"`
}

func (n NativeNetworkEncryptionNNE) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NativeNetworkEncryptionNNE) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *NativeNetworkEncryptionNNE) GetEncryptionAlgorithm() *EncryptionAlgorithm {
	if o == nil {
		return nil
	}
	return o.EncryptionAlgorithm
}

func (o *NativeNetworkEncryptionNNE) GetEncryptionMethod() *DestinationOracleUpdateEncryptionMethod {
	return DestinationOracleUpdateEncryptionMethodClientNne.ToPointer()
}

type EncryptionMethod string

const (
	EncryptionMethodUnencrypted EncryptionMethod = "unencrypted"
)

func (e EncryptionMethod) ToPointer() *EncryptionMethod {
	return &e
}
func (e *EncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unencrypted":
		*e = EncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EncryptionMethod: %v", v)
	}
}

// DestinationOracleUpdateUnencrypted - Data transfer will not be encrypted.
type DestinationOracleUpdateUnencrypted struct {
	encryptionMethod *EncryptionMethod `const:"unencrypted" json:"encryption_method"`
}

func (d DestinationOracleUpdateUnencrypted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleUpdateUnencrypted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleUpdateUnencrypted) GetEncryptionMethod() *EncryptionMethod {
	return EncryptionMethodUnencrypted.ToPointer()
}

type EncryptionType string

const (
	EncryptionTypeDestinationOracleUpdateUnencrypted EncryptionType = "destination-oracle-update_Unencrypted"
	EncryptionTypeNativeNetworkEncryptionNNE         EncryptionType = "Native Network Encryption (NNE)"
	EncryptionTypeTLSEncryptedVerifyCertificate      EncryptionType = "TLS Encrypted (verify certificate)"
)

// Encryption - The encryption method which is used when communicating with the database.
type Encryption struct {
	DestinationOracleUpdateUnencrypted *DestinationOracleUpdateUnencrypted
	NativeNetworkEncryptionNNE         *NativeNetworkEncryptionNNE
	TLSEncryptedVerifyCertificate      *TLSEncryptedVerifyCertificate

	Type EncryptionType
}

func CreateEncryptionDestinationOracleUpdateUnencrypted(destinationOracleUpdateUnencrypted DestinationOracleUpdateUnencrypted) Encryption {
	typ := EncryptionTypeDestinationOracleUpdateUnencrypted

	return Encryption{
		DestinationOracleUpdateUnencrypted: &destinationOracleUpdateUnencrypted,
		Type:                               typ,
	}
}

func CreateEncryptionNativeNetworkEncryptionNNE(nativeNetworkEncryptionNNE NativeNetworkEncryptionNNE) Encryption {
	typ := EncryptionTypeNativeNetworkEncryptionNNE

	return Encryption{
		NativeNetworkEncryptionNNE: &nativeNetworkEncryptionNNE,
		Type:                       typ,
	}
}

func CreateEncryptionTLSEncryptedVerifyCertificate(tlsEncryptedVerifyCertificate TLSEncryptedVerifyCertificate) Encryption {
	typ := EncryptionTypeTLSEncryptedVerifyCertificate

	return Encryption{
		TLSEncryptedVerifyCertificate: &tlsEncryptedVerifyCertificate,
		Type:                          typ,
	}
}

func (u *Encryption) UnmarshalJSON(data []byte) error {

	var destinationOracleUpdateUnencrypted DestinationOracleUpdateUnencrypted = DestinationOracleUpdateUnencrypted{}
	if err := utils.UnmarshalJSON(data, &destinationOracleUpdateUnencrypted, "", true, true); err == nil {
		u.DestinationOracleUpdateUnencrypted = &destinationOracleUpdateUnencrypted
		u.Type = EncryptionTypeDestinationOracleUpdateUnencrypted
		return nil
	}

	var nativeNetworkEncryptionNNE NativeNetworkEncryptionNNE = NativeNetworkEncryptionNNE{}
	if err := utils.UnmarshalJSON(data, &nativeNetworkEncryptionNNE, "", true, true); err == nil {
		u.NativeNetworkEncryptionNNE = &nativeNetworkEncryptionNNE
		u.Type = EncryptionTypeNativeNetworkEncryptionNNE
		return nil
	}

	var tlsEncryptedVerifyCertificate TLSEncryptedVerifyCertificate = TLSEncryptedVerifyCertificate{}
	if err := utils.UnmarshalJSON(data, &tlsEncryptedVerifyCertificate, "", true, true); err == nil {
		u.TLSEncryptedVerifyCertificate = &tlsEncryptedVerifyCertificate
		u.Type = EncryptionTypeTLSEncryptedVerifyCertificate
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Encryption", string(data))
}

func (u Encryption) MarshalJSON() ([]byte, error) {
	if u.DestinationOracleUpdateUnencrypted != nil {
		return utils.MarshalJSON(u.DestinationOracleUpdateUnencrypted, "", true)
	}

	if u.NativeNetworkEncryptionNNE != nil {
		return utils.MarshalJSON(u.NativeNetworkEncryptionNNE, "", true)
	}

	if u.TLSEncryptedVerifyCertificate != nil {
		return utils.MarshalJSON(u.TLSEncryptedVerifyCertificate, "", true)
	}

	return nil, errors.New("could not marshal union type Encryption: all fields are null")
}

// DestinationOracleUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationOracleUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationOracleUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationOracleUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationOracleUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationOracleUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationOracleUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationOracleUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationOracleUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationOracleUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationOracleUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationOracleUpdatePasswordAuthentication) GetTunnelMethod() DestinationOracleUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationOracleUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationOracleUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationOracleUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationOracleUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationOracleUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationOracleUpdateSchemasTunnelMethod string

const (
	DestinationOracleUpdateSchemasTunnelMethodSSHKeyAuth DestinationOracleUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationOracleUpdateSchemasTunnelMethod) ToPointer() *DestinationOracleUpdateSchemasTunnelMethod {
	return &e
}
func (e *DestinationOracleUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationOracleUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleUpdateSchemasTunnelMethod: %v", v)
	}
}

type DestinationOracleUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationOracleUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationOracleUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationOracleUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationOracleUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationOracleUpdateSchemasTunnelMethod {
	return DestinationOracleUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationOracleUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationOracleUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationOracleUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationOracleUpdateTunnelMethod string

const (
	DestinationOracleUpdateTunnelMethodNoTunnel DestinationOracleUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationOracleUpdateTunnelMethod) ToPointer() *DestinationOracleUpdateTunnelMethod {
	return &e
}
func (e *DestinationOracleUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationOracleUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleUpdateTunnelMethod: %v", v)
	}
}

type DestinationOracleUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationOracleUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationOracleUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleUpdateNoTunnel) GetTunnelMethod() DestinationOracleUpdateTunnelMethod {
	return DestinationOracleUpdateTunnelMethodNoTunnel
}

type DestinationOracleUpdateSSHTunnelMethodType string

const (
	DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdateNoTunnel               DestinationOracleUpdateSSHTunnelMethodType = "destination-oracle-update_No Tunnel"
	DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdateSSHKeyAuthentication   DestinationOracleUpdateSSHTunnelMethodType = "destination-oracle-update_SSH Key Authentication"
	DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdatePasswordAuthentication DestinationOracleUpdateSSHTunnelMethodType = "destination-oracle-update_Password Authentication"
)

// DestinationOracleUpdateSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationOracleUpdateSSHTunnelMethod struct {
	DestinationOracleUpdateNoTunnel               *DestinationOracleUpdateNoTunnel
	DestinationOracleUpdateSSHKeyAuthentication   *DestinationOracleUpdateSSHKeyAuthentication
	DestinationOracleUpdatePasswordAuthentication *DestinationOracleUpdatePasswordAuthentication

	Type DestinationOracleUpdateSSHTunnelMethodType
}

func CreateDestinationOracleUpdateSSHTunnelMethodDestinationOracleUpdateNoTunnel(destinationOracleUpdateNoTunnel DestinationOracleUpdateNoTunnel) DestinationOracleUpdateSSHTunnelMethod {
	typ := DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdateNoTunnel

	return DestinationOracleUpdateSSHTunnelMethod{
		DestinationOracleUpdateNoTunnel: &destinationOracleUpdateNoTunnel,
		Type:                            typ,
	}
}

func CreateDestinationOracleUpdateSSHTunnelMethodDestinationOracleUpdateSSHKeyAuthentication(destinationOracleUpdateSSHKeyAuthentication DestinationOracleUpdateSSHKeyAuthentication) DestinationOracleUpdateSSHTunnelMethod {
	typ := DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdateSSHKeyAuthentication

	return DestinationOracleUpdateSSHTunnelMethod{
		DestinationOracleUpdateSSHKeyAuthentication: &destinationOracleUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationOracleUpdateSSHTunnelMethodDestinationOracleUpdatePasswordAuthentication(destinationOracleUpdatePasswordAuthentication DestinationOracleUpdatePasswordAuthentication) DestinationOracleUpdateSSHTunnelMethod {
	typ := DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdatePasswordAuthentication

	return DestinationOracleUpdateSSHTunnelMethod{
		DestinationOracleUpdatePasswordAuthentication: &destinationOracleUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationOracleUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationOracleUpdateNoTunnel DestinationOracleUpdateNoTunnel = DestinationOracleUpdateNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationOracleUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationOracleUpdateNoTunnel = &destinationOracleUpdateNoTunnel
		u.Type = DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdateNoTunnel
		return nil
	}

	var destinationOracleUpdateSSHKeyAuthentication DestinationOracleUpdateSSHKeyAuthentication = DestinationOracleUpdateSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationOracleUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationOracleUpdateSSHKeyAuthentication = &destinationOracleUpdateSSHKeyAuthentication
		u.Type = DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdateSSHKeyAuthentication
		return nil
	}

	var destinationOracleUpdatePasswordAuthentication DestinationOracleUpdatePasswordAuthentication = DestinationOracleUpdatePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationOracleUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationOracleUpdatePasswordAuthentication = &destinationOracleUpdatePasswordAuthentication
		u.Type = DestinationOracleUpdateSSHTunnelMethodTypeDestinationOracleUpdatePasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationOracleUpdateSSHTunnelMethod", string(data))
}

func (u DestinationOracleUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationOracleUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationOracleUpdateNoTunnel, "", true)
	}

	if u.DestinationOracleUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationOracleUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationOracleUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationOracleUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationOracleUpdateSSHTunnelMethod: all fields are null")
}

type DestinationOracleUpdate struct {
	// The encryption method which is used when communicating with the database.
	Encryption *Encryption `json:"encryption,omitempty"`
	// The hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// The port of the database.
	Port *int64 `default:"1521" json:"port"`
	// The schema to write raw tables into (default: airbyte_internal)
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// The default schema is used as the target schema for all statements issued from the connection that do not explicitly specify a schema name. The usual value for this field is "airbyte".  In Oracle, schemas and users are the same thing, so the "user" parameter is used as the login credentials and this is used for the default Airbyte message schema.
	Schema *string `default:"airbyte" json:"schema"`
	// The System Identifier uniquely distinguishes the instance from any other instance on the same computer.
	Sid string `json:"sid"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationOracleUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username to access the database. This user must have CREATE USER privileges in the database.
	Username string `json:"username"`
}

func (d DestinationOracleUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleUpdate) GetEncryption() *Encryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *DestinationOracleUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationOracleUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationOracleUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationOracleUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationOracleUpdate) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationOracleUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationOracleUpdate) GetSid() string {
	if o == nil {
		return ""
	}
	return o.Sid
}

func (o *DestinationOracleUpdate) GetTunnelMethod() *DestinationOracleUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationOracleUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}