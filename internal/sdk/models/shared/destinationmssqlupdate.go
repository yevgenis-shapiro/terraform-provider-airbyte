// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationMssqlUpdateSchemasSslMethodSslMethod string

const (
	DestinationMssqlUpdateSchemasSslMethodSslMethodEncryptedVerifyCertificate DestinationMssqlUpdateSchemasSslMethodSslMethod = "encrypted_verify_certificate"
)

func (e DestinationMssqlUpdateSchemasSslMethodSslMethod) ToPointer() *DestinationMssqlUpdateSchemasSslMethodSslMethod {
	return &e
}
func (e *DestinationMssqlUpdateSchemasSslMethodSslMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = DestinationMssqlUpdateSchemasSslMethodSslMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSchemasSslMethodSslMethod: %v", v)
	}
}

// EncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type EncryptedVerifyCertificate struct {
	// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
	HostNameInCertificate *string                                          `json:"hostNameInCertificate,omitempty"`
	sslMethod             *DestinationMssqlUpdateSchemasSslMethodSslMethod `const:"encrypted_verify_certificate" json:"ssl_method"`
}

func (e EncryptedVerifyCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EncryptedVerifyCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *EncryptedVerifyCertificate) GetHostNameInCertificate() *string {
	if o == nil {
		return nil
	}
	return o.HostNameInCertificate
}

func (o *EncryptedVerifyCertificate) GetSslMethod() *DestinationMssqlUpdateSchemasSslMethodSslMethod {
	return DestinationMssqlUpdateSchemasSslMethodSslMethodEncryptedVerifyCertificate.ToPointer()
}

type DestinationMssqlUpdateSchemasSslMethod string

const (
	DestinationMssqlUpdateSchemasSslMethodEncryptedTrustServerCertificate DestinationMssqlUpdateSchemasSslMethod = "encrypted_trust_server_certificate"
)

func (e DestinationMssqlUpdateSchemasSslMethod) ToPointer() *DestinationMssqlUpdateSchemasSslMethod {
	return &e
}
func (e *DestinationMssqlUpdateSchemasSslMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_trust_server_certificate":
		*e = DestinationMssqlUpdateSchemasSslMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSchemasSslMethod: %v", v)
	}
}

// EncryptedTrustServerCertificate - Use the certificate provided by the server without verification. (For testing purposes only!)
type EncryptedTrustServerCertificate struct {
	sslMethod *DestinationMssqlUpdateSchemasSslMethod `const:"encrypted_trust_server_certificate" json:"ssl_method"`
}

func (e EncryptedTrustServerCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EncryptedTrustServerCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *EncryptedTrustServerCertificate) GetSslMethod() *DestinationMssqlUpdateSchemasSslMethod {
	return DestinationMssqlUpdateSchemasSslMethodEncryptedTrustServerCertificate.ToPointer()
}

type DestinationMssqlUpdateSslMethod string

const (
	DestinationMssqlUpdateSslMethodUnencrypted DestinationMssqlUpdateSslMethod = "unencrypted"
)

func (e DestinationMssqlUpdateSslMethod) ToPointer() *DestinationMssqlUpdateSslMethod {
	return &e
}
func (e *DestinationMssqlUpdateSslMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unencrypted":
		*e = DestinationMssqlUpdateSslMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSslMethod: %v", v)
	}
}

// Unencrypted - The data transfer will not be encrypted.
type Unencrypted struct {
	sslMethod *DestinationMssqlUpdateSslMethod `const:"unencrypted" json:"ssl_method"`
}

func (u Unencrypted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *Unencrypted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Unencrypted) GetSslMethod() *DestinationMssqlUpdateSslMethod {
	return DestinationMssqlUpdateSslMethodUnencrypted.ToPointer()
}

type SSLMethodType string

const (
	SSLMethodTypeUnencrypted                     SSLMethodType = "Unencrypted"
	SSLMethodTypeEncryptedTrustServerCertificate SSLMethodType = "Encrypted (trust server certificate)"
	SSLMethodTypeEncryptedVerifyCertificate      SSLMethodType = "Encrypted (verify certificate)"
)

// SSLMethod - The encryption method which is used to communicate with the database.
type SSLMethod struct {
	Unencrypted                     *Unencrypted
	EncryptedTrustServerCertificate *EncryptedTrustServerCertificate
	EncryptedVerifyCertificate      *EncryptedVerifyCertificate

	Type SSLMethodType
}

func CreateSSLMethodUnencrypted(unencrypted Unencrypted) SSLMethod {
	typ := SSLMethodTypeUnencrypted

	return SSLMethod{
		Unencrypted: &unencrypted,
		Type:        typ,
	}
}

func CreateSSLMethodEncryptedTrustServerCertificate(encryptedTrustServerCertificate EncryptedTrustServerCertificate) SSLMethod {
	typ := SSLMethodTypeEncryptedTrustServerCertificate

	return SSLMethod{
		EncryptedTrustServerCertificate: &encryptedTrustServerCertificate,
		Type:                            typ,
	}
}

func CreateSSLMethodEncryptedVerifyCertificate(encryptedVerifyCertificate EncryptedVerifyCertificate) SSLMethod {
	typ := SSLMethodTypeEncryptedVerifyCertificate

	return SSLMethod{
		EncryptedVerifyCertificate: &encryptedVerifyCertificate,
		Type:                       typ,
	}
}

func (u *SSLMethod) UnmarshalJSON(data []byte) error {

	var unencrypted Unencrypted = Unencrypted{}
	if err := utils.UnmarshalJSON(data, &unencrypted, "", true, true); err == nil {
		u.Unencrypted = &unencrypted
		u.Type = SSLMethodTypeUnencrypted
		return nil
	}

	var encryptedTrustServerCertificate EncryptedTrustServerCertificate = EncryptedTrustServerCertificate{}
	if err := utils.UnmarshalJSON(data, &encryptedTrustServerCertificate, "", true, true); err == nil {
		u.EncryptedTrustServerCertificate = &encryptedTrustServerCertificate
		u.Type = SSLMethodTypeEncryptedTrustServerCertificate
		return nil
	}

	var encryptedVerifyCertificate EncryptedVerifyCertificate = EncryptedVerifyCertificate{}
	if err := utils.UnmarshalJSON(data, &encryptedVerifyCertificate, "", true, true); err == nil {
		u.EncryptedVerifyCertificate = &encryptedVerifyCertificate
		u.Type = SSLMethodTypeEncryptedVerifyCertificate
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SSLMethod", string(data))
}

func (u SSLMethod) MarshalJSON() ([]byte, error) {
	if u.Unencrypted != nil {
		return utils.MarshalJSON(u.Unencrypted, "", true)
	}

	if u.EncryptedTrustServerCertificate != nil {
		return utils.MarshalJSON(u.EncryptedTrustServerCertificate, "", true)
	}

	if u.EncryptedVerifyCertificate != nil {
		return utils.MarshalJSON(u.EncryptedVerifyCertificate, "", true)
	}

	return nil, errors.New("could not marshal union type SSLMethod: all fields are null")
}

// DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationMssqlUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationMssqlUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationMssqlUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMssqlUpdatePasswordAuthentication) GetTunnelMethod() DestinationMssqlUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationMssqlUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationMssqlUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMssqlUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationMssqlUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationMssqlUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMssqlUpdateSchemasTunnelMethod string

const (
	DestinationMssqlUpdateSchemasTunnelMethodSSHKeyAuth DestinationMssqlUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMssqlUpdateSchemasTunnelMethod) ToPointer() *DestinationMssqlUpdateSchemasTunnelMethod {
	return &e
}
func (e *DestinationMssqlUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMssqlUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSchemasTunnelMethod: %v", v)
	}
}

type DestinationMssqlUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationMssqlUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationMssqlUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationMssqlUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMssqlUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationMssqlUpdateSchemasTunnelMethod {
	return DestinationMssqlUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationMssqlUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMssqlUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationMssqlUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMssqlUpdateTunnelMethod string

const (
	DestinationMssqlUpdateTunnelMethodNoTunnel DestinationMssqlUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMssqlUpdateTunnelMethod) ToPointer() *DestinationMssqlUpdateTunnelMethod {
	return &e
}
func (e *DestinationMssqlUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMssqlUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateTunnelMethod: %v", v)
	}
}

type DestinationMssqlUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationMssqlUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationMssqlUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateNoTunnel) GetTunnelMethod() DestinationMssqlUpdateTunnelMethod {
	return DestinationMssqlUpdateTunnelMethodNoTunnel
}

type DestinationMssqlUpdateSSHTunnelMethodType string

const (
	DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateNoTunnel               DestinationMssqlUpdateSSHTunnelMethodType = "destination-mssql-update_No Tunnel"
	DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHKeyAuthentication   DestinationMssqlUpdateSSHTunnelMethodType = "destination-mssql-update_SSH Key Authentication"
	DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdatePasswordAuthentication DestinationMssqlUpdateSSHTunnelMethodType = "destination-mssql-update_Password Authentication"
)

// DestinationMssqlUpdateSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlUpdateSSHTunnelMethod struct {
	DestinationMssqlUpdateNoTunnel               *DestinationMssqlUpdateNoTunnel
	DestinationMssqlUpdateSSHKeyAuthentication   *DestinationMssqlUpdateSSHKeyAuthentication
	DestinationMssqlUpdatePasswordAuthentication *DestinationMssqlUpdatePasswordAuthentication

	Type DestinationMssqlUpdateSSHTunnelMethodType
}

func CreateDestinationMssqlUpdateSSHTunnelMethodDestinationMssqlUpdateNoTunnel(destinationMssqlUpdateNoTunnel DestinationMssqlUpdateNoTunnel) DestinationMssqlUpdateSSHTunnelMethod {
	typ := DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateNoTunnel

	return DestinationMssqlUpdateSSHTunnelMethod{
		DestinationMssqlUpdateNoTunnel: &destinationMssqlUpdateNoTunnel,
		Type:                           typ,
	}
}

func CreateDestinationMssqlUpdateSSHTunnelMethodDestinationMssqlUpdateSSHKeyAuthentication(destinationMssqlUpdateSSHKeyAuthentication DestinationMssqlUpdateSSHKeyAuthentication) DestinationMssqlUpdateSSHTunnelMethod {
	typ := DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHKeyAuthentication

	return DestinationMssqlUpdateSSHTunnelMethod{
		DestinationMssqlUpdateSSHKeyAuthentication: &destinationMssqlUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMssqlUpdateSSHTunnelMethodDestinationMssqlUpdatePasswordAuthentication(destinationMssqlUpdatePasswordAuthentication DestinationMssqlUpdatePasswordAuthentication) DestinationMssqlUpdateSSHTunnelMethod {
	typ := DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdatePasswordAuthentication

	return DestinationMssqlUpdateSSHTunnelMethod{
		DestinationMssqlUpdatePasswordAuthentication: &destinationMssqlUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMssqlUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationMssqlUpdateNoTunnel DestinationMssqlUpdateNoTunnel = DestinationMssqlUpdateNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationMssqlUpdateNoTunnel = &destinationMssqlUpdateNoTunnel
		u.Type = DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateNoTunnel
		return nil
	}

	var destinationMssqlUpdateSSHKeyAuthentication DestinationMssqlUpdateSSHKeyAuthentication = DestinationMssqlUpdateSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationMssqlUpdateSSHKeyAuthentication = &destinationMssqlUpdateSSHKeyAuthentication
		u.Type = DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHKeyAuthentication
		return nil
	}

	var destinationMssqlUpdatePasswordAuthentication DestinationMssqlUpdatePasswordAuthentication = DestinationMssqlUpdatePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationMssqlUpdatePasswordAuthentication = &destinationMssqlUpdatePasswordAuthentication
		u.Type = DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdatePasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationMssqlUpdateSSHTunnelMethod", string(data))
}

func (u DestinationMssqlUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMssqlUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateNoTunnel, "", true)
	}

	if u.DestinationMssqlUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationMssqlUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationMssqlUpdateSSHTunnelMethod: all fields are null")
}

type DestinationMssqlUpdate struct {
	// The name of the MSSQL database.
	Database string `json:"database"`
	// The host name of the MSSQL database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with this username.
	Password *string `json:"password,omitempty"`
	// The port of the MSSQL database.
	Port *int64 `default:"1433" json:"port"`
	// The schema to write raw tables into (default: airbyte_internal)
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// The encryption method which is used to communicate with the database.
	SslMethod *SSLMethod `json:"ssl_method,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMssqlUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (d DestinationMssqlUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationMssqlUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationMssqlUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationMssqlUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationMssqlUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationMssqlUpdate) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationMssqlUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationMssqlUpdate) GetSslMethod() *SSLMethod {
	if o == nil {
		return nil
	}
	return o.SslMethod
}

func (o *DestinationMssqlUpdate) GetTunnelMethod() *DestinationMssqlUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationMssqlUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
