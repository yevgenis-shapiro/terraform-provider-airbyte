// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Oracle string

const (
	OracleOracle Oracle = "oracle"
)

func (e Oracle) ToPointer() *Oracle {
	return &e
}
func (e *Oracle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oracle":
		*e = Oracle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Oracle: %v", v)
	}
}

type DestinationOracleSchemasEncryptionEncryptionMethod string

const (
	DestinationOracleSchemasEncryptionEncryptionMethodEncryptedVerifyCertificate DestinationOracleSchemasEncryptionEncryptionMethod = "encrypted_verify_certificate"
)

func (e DestinationOracleSchemasEncryptionEncryptionMethod) ToPointer() *DestinationOracleSchemasEncryptionEncryptionMethod {
	return &e
}
func (e *DestinationOracleSchemasEncryptionEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = DestinationOracleSchemasEncryptionEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleSchemasEncryptionEncryptionMethod: %v", v)
	}
}

// DestinationOracleTLSEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type DestinationOracleTLSEncryptedVerifyCertificate struct {
	encryptionMethod *DestinationOracleSchemasEncryptionEncryptionMethod `const:"encrypted_verify_certificate" json:"encryption_method"`
	// Privacy Enhanced Mail (PEM) files are concatenated certificate containers frequently used in certificate installations.
	SslCertificate string `json:"ssl_certificate"`
}

func (d DestinationOracleTLSEncryptedVerifyCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleTLSEncryptedVerifyCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleTLSEncryptedVerifyCertificate) GetEncryptionMethod() *DestinationOracleSchemasEncryptionEncryptionMethod {
	return DestinationOracleSchemasEncryptionEncryptionMethodEncryptedVerifyCertificate.ToPointer()
}

func (o *DestinationOracleTLSEncryptedVerifyCertificate) GetSslCertificate() string {
	if o == nil {
		return ""
	}
	return o.SslCertificate
}

// DestinationOracleEncryptionAlgorithm - This parameter defines the database encryption algorithm.
type DestinationOracleEncryptionAlgorithm string

const (
	DestinationOracleEncryptionAlgorithmAes256      DestinationOracleEncryptionAlgorithm = "AES256"
	DestinationOracleEncryptionAlgorithmRc456       DestinationOracleEncryptionAlgorithm = "RC4_56"
	DestinationOracleEncryptionAlgorithmThreeDes168 DestinationOracleEncryptionAlgorithm = "3DES168"
)

func (e DestinationOracleEncryptionAlgorithm) ToPointer() *DestinationOracleEncryptionAlgorithm {
	return &e
}
func (e *DestinationOracleEncryptionAlgorithm) UnmarshalJSON(data []byte) error {
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
		*e = DestinationOracleEncryptionAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleEncryptionAlgorithm: %v", v)
	}
}

type DestinationOracleSchemasEncryptionMethod string

const (
	DestinationOracleSchemasEncryptionMethodClientNne DestinationOracleSchemasEncryptionMethod = "client_nne"
)

func (e DestinationOracleSchemasEncryptionMethod) ToPointer() *DestinationOracleSchemasEncryptionMethod {
	return &e
}
func (e *DestinationOracleSchemasEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_nne":
		*e = DestinationOracleSchemasEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleSchemasEncryptionMethod: %v", v)
	}
}

// DestinationOracleNativeNetworkEncryptionNNE - The native network encryption gives you the ability to encrypt database connections, without the configuration overhead of TCP/IP and SSL/TLS and without the need to open and listen on different ports.
type DestinationOracleNativeNetworkEncryptionNNE struct {
	// This parameter defines the database encryption algorithm.
	EncryptionAlgorithm *DestinationOracleEncryptionAlgorithm     `default:"AES256" json:"encryption_algorithm"`
	encryptionMethod    *DestinationOracleSchemasEncryptionMethod `const:"client_nne" json:"encryption_method"`
}

func (d DestinationOracleNativeNetworkEncryptionNNE) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleNativeNetworkEncryptionNNE) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleNativeNetworkEncryptionNNE) GetEncryptionAlgorithm() *DestinationOracleEncryptionAlgorithm {
	if o == nil {
		return nil
	}
	return o.EncryptionAlgorithm
}

func (o *DestinationOracleNativeNetworkEncryptionNNE) GetEncryptionMethod() *DestinationOracleSchemasEncryptionMethod {
	return DestinationOracleSchemasEncryptionMethodClientNne.ToPointer()
}

type DestinationOracleEncryptionMethod string

const (
	DestinationOracleEncryptionMethodUnencrypted DestinationOracleEncryptionMethod = "unencrypted"
)

func (e DestinationOracleEncryptionMethod) ToPointer() *DestinationOracleEncryptionMethod {
	return &e
}
func (e *DestinationOracleEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unencrypted":
		*e = DestinationOracleEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleEncryptionMethod: %v", v)
	}
}

// DestinationOracleUnencrypted - Data transfer will not be encrypted.
type DestinationOracleUnencrypted struct {
	encryptionMethod *DestinationOracleEncryptionMethod `const:"unencrypted" json:"encryption_method"`
}

func (d DestinationOracleUnencrypted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleUnencrypted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleUnencrypted) GetEncryptionMethod() *DestinationOracleEncryptionMethod {
	return DestinationOracleEncryptionMethodUnencrypted.ToPointer()
}

type DestinationOracleEncryptionType string

const (
	DestinationOracleEncryptionTypeDestinationOracleUnencrypted                   DestinationOracleEncryptionType = "destination-oracle_Unencrypted"
	DestinationOracleEncryptionTypeDestinationOracleNativeNetworkEncryptionNNE    DestinationOracleEncryptionType = "destination-oracle_Native Network Encryption (NNE)"
	DestinationOracleEncryptionTypeDestinationOracleTLSEncryptedVerifyCertificate DestinationOracleEncryptionType = "destination-oracle_TLS Encrypted (verify certificate)"
)

// DestinationOracleEncryption - The encryption method which is used when communicating with the database.
type DestinationOracleEncryption struct {
	DestinationOracleUnencrypted                   *DestinationOracleUnencrypted
	DestinationOracleNativeNetworkEncryptionNNE    *DestinationOracleNativeNetworkEncryptionNNE
	DestinationOracleTLSEncryptedVerifyCertificate *DestinationOracleTLSEncryptedVerifyCertificate

	Type DestinationOracleEncryptionType
}

func CreateDestinationOracleEncryptionDestinationOracleUnencrypted(destinationOracleUnencrypted DestinationOracleUnencrypted) DestinationOracleEncryption {
	typ := DestinationOracleEncryptionTypeDestinationOracleUnencrypted

	return DestinationOracleEncryption{
		DestinationOracleUnencrypted: &destinationOracleUnencrypted,
		Type:                         typ,
	}
}

func CreateDestinationOracleEncryptionDestinationOracleNativeNetworkEncryptionNNE(destinationOracleNativeNetworkEncryptionNNE DestinationOracleNativeNetworkEncryptionNNE) DestinationOracleEncryption {
	typ := DestinationOracleEncryptionTypeDestinationOracleNativeNetworkEncryptionNNE

	return DestinationOracleEncryption{
		DestinationOracleNativeNetworkEncryptionNNE: &destinationOracleNativeNetworkEncryptionNNE,
		Type: typ,
	}
}

func CreateDestinationOracleEncryptionDestinationOracleTLSEncryptedVerifyCertificate(destinationOracleTLSEncryptedVerifyCertificate DestinationOracleTLSEncryptedVerifyCertificate) DestinationOracleEncryption {
	typ := DestinationOracleEncryptionTypeDestinationOracleTLSEncryptedVerifyCertificate

	return DestinationOracleEncryption{
		DestinationOracleTLSEncryptedVerifyCertificate: &destinationOracleTLSEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *DestinationOracleEncryption) UnmarshalJSON(data []byte) error {

	var destinationOracleUnencrypted DestinationOracleUnencrypted = DestinationOracleUnencrypted{}
	if err := utils.UnmarshalJSON(data, &destinationOracleUnencrypted, "", true, true); err == nil {
		u.DestinationOracleUnencrypted = &destinationOracleUnencrypted
		u.Type = DestinationOracleEncryptionTypeDestinationOracleUnencrypted
		return nil
	}

	var destinationOracleNativeNetworkEncryptionNNE DestinationOracleNativeNetworkEncryptionNNE = DestinationOracleNativeNetworkEncryptionNNE{}
	if err := utils.UnmarshalJSON(data, &destinationOracleNativeNetworkEncryptionNNE, "", true, true); err == nil {
		u.DestinationOracleNativeNetworkEncryptionNNE = &destinationOracleNativeNetworkEncryptionNNE
		u.Type = DestinationOracleEncryptionTypeDestinationOracleNativeNetworkEncryptionNNE
		return nil
	}

	var destinationOracleTLSEncryptedVerifyCertificate DestinationOracleTLSEncryptedVerifyCertificate = DestinationOracleTLSEncryptedVerifyCertificate{}
	if err := utils.UnmarshalJSON(data, &destinationOracleTLSEncryptedVerifyCertificate, "", true, true); err == nil {
		u.DestinationOracleTLSEncryptedVerifyCertificate = &destinationOracleTLSEncryptedVerifyCertificate
		u.Type = DestinationOracleEncryptionTypeDestinationOracleTLSEncryptedVerifyCertificate
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationOracleEncryption", string(data))
}

func (u DestinationOracleEncryption) MarshalJSON() ([]byte, error) {
	if u.DestinationOracleUnencrypted != nil {
		return utils.MarshalJSON(u.DestinationOracleUnencrypted, "", true)
	}

	if u.DestinationOracleNativeNetworkEncryptionNNE != nil {
		return utils.MarshalJSON(u.DestinationOracleNativeNetworkEncryptionNNE, "", true)
	}

	if u.DestinationOracleTLSEncryptedVerifyCertificate != nil {
		return utils.MarshalJSON(u.DestinationOracleTLSEncryptedVerifyCertificate, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationOracleEncryption: all fields are null")
}

// DestinationOracleSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationOracleSchemasTunnelMethodTunnelMethod string

const (
	DestinationOracleSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationOracleSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationOracleSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationOracleSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationOracleSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationOracleSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationOraclePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationOracleSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationOraclePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOraclePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOraclePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationOraclePasswordAuthentication) GetTunnelMethod() DestinationOracleSchemasTunnelMethodTunnelMethod {
	return DestinationOracleSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationOraclePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationOraclePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationOraclePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationOracleSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationOracleSchemasTunnelMethod string

const (
	DestinationOracleSchemasTunnelMethodSSHKeyAuth DestinationOracleSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationOracleSchemasTunnelMethod) ToPointer() *DestinationOracleSchemasTunnelMethod {
	return &e
}
func (e *DestinationOracleSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationOracleSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleSchemasTunnelMethod: %v", v)
	}
}

type DestinationOracleSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationOracleSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationOracleSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationOracleSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationOracleSSHKeyAuthentication) GetTunnelMethod() DestinationOracleSchemasTunnelMethod {
	return DestinationOracleSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationOracleSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationOracleSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationOracleTunnelMethod - No ssh tunnel needed to connect to database
type DestinationOracleTunnelMethod string

const (
	DestinationOracleTunnelMethodNoTunnel DestinationOracleTunnelMethod = "NO_TUNNEL"
)

func (e DestinationOracleTunnelMethod) ToPointer() *DestinationOracleTunnelMethod {
	return &e
}
func (e *DestinationOracleTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationOracleTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationOracleTunnelMethod: %v", v)
	}
}

type DestinationOracleNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationOracleTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationOracleNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracleNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracleNoTunnel) GetTunnelMethod() DestinationOracleTunnelMethod {
	return DestinationOracleTunnelMethodNoTunnel
}

type DestinationOracleSSHTunnelMethodType string

const (
	DestinationOracleSSHTunnelMethodTypeDestinationOracleNoTunnel               DestinationOracleSSHTunnelMethodType = "destination-oracle_No Tunnel"
	DestinationOracleSSHTunnelMethodTypeDestinationOracleSSHKeyAuthentication   DestinationOracleSSHTunnelMethodType = "destination-oracle_SSH Key Authentication"
	DestinationOracleSSHTunnelMethodTypeDestinationOraclePasswordAuthentication DestinationOracleSSHTunnelMethodType = "destination-oracle_Password Authentication"
)

// DestinationOracleSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationOracleSSHTunnelMethod struct {
	DestinationOracleNoTunnel               *DestinationOracleNoTunnel
	DestinationOracleSSHKeyAuthentication   *DestinationOracleSSHKeyAuthentication
	DestinationOraclePasswordAuthentication *DestinationOraclePasswordAuthentication

	Type DestinationOracleSSHTunnelMethodType
}

func CreateDestinationOracleSSHTunnelMethodDestinationOracleNoTunnel(destinationOracleNoTunnel DestinationOracleNoTunnel) DestinationOracleSSHTunnelMethod {
	typ := DestinationOracleSSHTunnelMethodTypeDestinationOracleNoTunnel

	return DestinationOracleSSHTunnelMethod{
		DestinationOracleNoTunnel: &destinationOracleNoTunnel,
		Type:                      typ,
	}
}

func CreateDestinationOracleSSHTunnelMethodDestinationOracleSSHKeyAuthentication(destinationOracleSSHKeyAuthentication DestinationOracleSSHKeyAuthentication) DestinationOracleSSHTunnelMethod {
	typ := DestinationOracleSSHTunnelMethodTypeDestinationOracleSSHKeyAuthentication

	return DestinationOracleSSHTunnelMethod{
		DestinationOracleSSHKeyAuthentication: &destinationOracleSSHKeyAuthentication,
		Type:                                  typ,
	}
}

func CreateDestinationOracleSSHTunnelMethodDestinationOraclePasswordAuthentication(destinationOraclePasswordAuthentication DestinationOraclePasswordAuthentication) DestinationOracleSSHTunnelMethod {
	typ := DestinationOracleSSHTunnelMethodTypeDestinationOraclePasswordAuthentication

	return DestinationOracleSSHTunnelMethod{
		DestinationOraclePasswordAuthentication: &destinationOraclePasswordAuthentication,
		Type:                                    typ,
	}
}

func (u *DestinationOracleSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationOracleNoTunnel DestinationOracleNoTunnel = DestinationOracleNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationOracleNoTunnel, "", true, true); err == nil {
		u.DestinationOracleNoTunnel = &destinationOracleNoTunnel
		u.Type = DestinationOracleSSHTunnelMethodTypeDestinationOracleNoTunnel
		return nil
	}

	var destinationOracleSSHKeyAuthentication DestinationOracleSSHKeyAuthentication = DestinationOracleSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationOracleSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationOracleSSHKeyAuthentication = &destinationOracleSSHKeyAuthentication
		u.Type = DestinationOracleSSHTunnelMethodTypeDestinationOracleSSHKeyAuthentication
		return nil
	}

	var destinationOraclePasswordAuthentication DestinationOraclePasswordAuthentication = DestinationOraclePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationOraclePasswordAuthentication, "", true, true); err == nil {
		u.DestinationOraclePasswordAuthentication = &destinationOraclePasswordAuthentication
		u.Type = DestinationOracleSSHTunnelMethodTypeDestinationOraclePasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationOracleSSHTunnelMethod", string(data))
}

func (u DestinationOracleSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationOracleNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationOracleNoTunnel, "", true)
	}

	if u.DestinationOracleSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationOracleSSHKeyAuthentication, "", true)
	}

	if u.DestinationOraclePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationOraclePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationOracleSSHTunnelMethod: all fields are null")
}

type DestinationOracle struct {
	destinationType Oracle `const:"oracle" json:"destinationType"`
	// The encryption method which is used when communicating with the database.
	Encryption *DestinationOracleEncryption `json:"encryption,omitempty"`
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
	TunnelMethod *DestinationOracleSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username to access the database. This user must have CREATE USER privileges in the database.
	Username string `json:"username"`
}

func (d DestinationOracle) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationOracle) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationOracle) GetDestinationType() Oracle {
	return OracleOracle
}

func (o *DestinationOracle) GetEncryption() *DestinationOracleEncryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *DestinationOracle) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationOracle) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationOracle) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationOracle) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationOracle) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationOracle) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationOracle) GetSid() string {
	if o == nil {
		return ""
	}
	return o.Sid
}

func (o *DestinationOracle) GetTunnelMethod() *DestinationOracleSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationOracle) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}