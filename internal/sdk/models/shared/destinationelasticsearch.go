// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationElasticsearchSchemasAuthenticationMethodMethod string

const (
	DestinationElasticsearchSchemasAuthenticationMethodMethodBasic DestinationElasticsearchSchemasAuthenticationMethodMethod = "basic"
)

func (e DestinationElasticsearchSchemasAuthenticationMethodMethod) ToPointer() *DestinationElasticsearchSchemasAuthenticationMethodMethod {
	return &e
}
func (e *DestinationElasticsearchSchemasAuthenticationMethodMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = DestinationElasticsearchSchemasAuthenticationMethodMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchSchemasAuthenticationMethodMethod: %v", v)
	}
}

// DestinationElasticsearchUsernamePassword - Basic auth header with a username and password
type DestinationElasticsearchUsernamePassword struct {
	method DestinationElasticsearchSchemasAuthenticationMethodMethod `const:"basic" json:"method"`
	// Basic auth password to access a secure Elasticsearch server
	Password string `json:"password"`
	// Basic auth username to access a secure Elasticsearch server
	Username string `json:"username"`
}

func (d DestinationElasticsearchUsernamePassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchUsernamePassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchUsernamePassword) GetMethod() DestinationElasticsearchSchemasAuthenticationMethodMethod {
	return DestinationElasticsearchSchemasAuthenticationMethodMethodBasic
}

func (o *DestinationElasticsearchUsernamePassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *DestinationElasticsearchUsernamePassword) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type DestinationElasticsearchSchemasMethod string

const (
	DestinationElasticsearchSchemasMethodSecret DestinationElasticsearchSchemasMethod = "secret"
)

func (e DestinationElasticsearchSchemasMethod) ToPointer() *DestinationElasticsearchSchemasMethod {
	return &e
}
func (e *DestinationElasticsearchSchemasMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		*e = DestinationElasticsearchSchemasMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchSchemasMethod: %v", v)
	}
}

// DestinationElasticsearchAPIKeySecret - Use a api key and secret combination to authenticate
type DestinationElasticsearchAPIKeySecret struct {
	// The Key ID to used when accessing an enterprise Elasticsearch instance.
	APIKeyID string `json:"apiKeyId"`
	// The secret associated with the API Key ID.
	APIKeySecret string                                `json:"apiKeySecret"`
	method       DestinationElasticsearchSchemasMethod `const:"secret" json:"method"`
}

func (d DestinationElasticsearchAPIKeySecret) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchAPIKeySecret) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchAPIKeySecret) GetAPIKeyID() string {
	if o == nil {
		return ""
	}
	return o.APIKeyID
}

func (o *DestinationElasticsearchAPIKeySecret) GetAPIKeySecret() string {
	if o == nil {
		return ""
	}
	return o.APIKeySecret
}

func (o *DestinationElasticsearchAPIKeySecret) GetMethod() DestinationElasticsearchSchemasMethod {
	return DestinationElasticsearchSchemasMethodSecret
}

type DestinationElasticsearchMethod string

const (
	DestinationElasticsearchMethodNone DestinationElasticsearchMethod = "none"
)

func (e DestinationElasticsearchMethod) ToPointer() *DestinationElasticsearchMethod {
	return &e
}
func (e *DestinationElasticsearchMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = DestinationElasticsearchMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchMethod: %v", v)
	}
}

// DestinationElasticsearchNone - No authentication will be used
type DestinationElasticsearchNone struct {
	method DestinationElasticsearchMethod `const:"none" json:"method"`
}

func (d DestinationElasticsearchNone) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchNone) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchNone) GetMethod() DestinationElasticsearchMethod {
	return DestinationElasticsearchMethodNone
}

type DestinationElasticsearchAuthenticationMethodType string

const (
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchNone             DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_None"
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAPIKeySecret     DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_Api Key/Secret"
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchUsernamePassword DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_Username/Password"
)

// DestinationElasticsearchAuthenticationMethod - The type of authentication to be used
type DestinationElasticsearchAuthenticationMethod struct {
	DestinationElasticsearchNone             *DestinationElasticsearchNone
	DestinationElasticsearchAPIKeySecret     *DestinationElasticsearchAPIKeySecret
	DestinationElasticsearchUsernamePassword *DestinationElasticsearchUsernamePassword

	Type DestinationElasticsearchAuthenticationMethodType
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchNone(destinationElasticsearchNone DestinationElasticsearchNone) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchNone

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchNone: &destinationElasticsearchNone,
		Type:                         typ,
	}
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchAPIKeySecret(destinationElasticsearchAPIKeySecret DestinationElasticsearchAPIKeySecret) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAPIKeySecret

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchAPIKeySecret: &destinationElasticsearchAPIKeySecret,
		Type:                                 typ,
	}
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchUsernamePassword(destinationElasticsearchUsernamePassword DestinationElasticsearchUsernamePassword) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchUsernamePassword

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchUsernamePassword: &destinationElasticsearchUsernamePassword,
		Type:                                     typ,
	}
}

func (u *DestinationElasticsearchAuthenticationMethod) UnmarshalJSON(data []byte) error {

	var destinationElasticsearchNone DestinationElasticsearchNone = DestinationElasticsearchNone{}
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchNone, "", true, true); err == nil {
		u.DestinationElasticsearchNone = &destinationElasticsearchNone
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchNone
		return nil
	}

	var destinationElasticsearchAPIKeySecret DestinationElasticsearchAPIKeySecret = DestinationElasticsearchAPIKeySecret{}
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchAPIKeySecret, "", true, true); err == nil {
		u.DestinationElasticsearchAPIKeySecret = &destinationElasticsearchAPIKeySecret
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAPIKeySecret
		return nil
	}

	var destinationElasticsearchUsernamePassword DestinationElasticsearchUsernamePassword = DestinationElasticsearchUsernamePassword{}
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchUsernamePassword, "", true, true); err == nil {
		u.DestinationElasticsearchUsernamePassword = &destinationElasticsearchUsernamePassword
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchUsernamePassword
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationElasticsearchAuthenticationMethod", string(data))
}

func (u DestinationElasticsearchAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationElasticsearchNone != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchNone, "", true)
	}

	if u.DestinationElasticsearchAPIKeySecret != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchAPIKeySecret, "", true)
	}

	if u.DestinationElasticsearchUsernamePassword != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchUsernamePassword, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationElasticsearchAuthenticationMethod: all fields are null")
}

type Elasticsearch string

const (
	ElasticsearchElasticsearch Elasticsearch = "elasticsearch"
)

func (e Elasticsearch) ToPointer() *Elasticsearch {
	return &e
}
func (e *Elasticsearch) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "elasticsearch":
		*e = Elasticsearch(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Elasticsearch: %v", v)
	}
}

// DestinationElasticsearchSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationElasticsearchSchemasTunnelMethodTunnelMethod string

const (
	DestinationElasticsearchSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationElasticsearchSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationElasticsearchSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationElasticsearchSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationElasticsearchSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationElasticsearchSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationElasticsearchPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationElasticsearchSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationElasticsearchPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationElasticsearchPasswordAuthentication) GetTunnelMethod() DestinationElasticsearchSchemasTunnelMethodTunnelMethod {
	return DestinationElasticsearchSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationElasticsearchPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationElasticsearchPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationElasticsearchPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationElasticsearchSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationElasticsearchSchemasTunnelMethod string

const (
	DestinationElasticsearchSchemasTunnelMethodSSHKeyAuth DestinationElasticsearchSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationElasticsearchSchemasTunnelMethod) ToPointer() *DestinationElasticsearchSchemasTunnelMethod {
	return &e
}
func (e *DestinationElasticsearchSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationElasticsearchSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchSchemasTunnelMethod: %v", v)
	}
}

type DestinationElasticsearchSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationElasticsearchSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationElasticsearchSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationElasticsearchSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationElasticsearchSSHKeyAuthentication) GetTunnelMethod() DestinationElasticsearchSchemasTunnelMethod {
	return DestinationElasticsearchSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationElasticsearchSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationElasticsearchSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationElasticsearchTunnelMethod - No ssh tunnel needed to connect to database
type DestinationElasticsearchTunnelMethod string

const (
	DestinationElasticsearchTunnelMethodNoTunnel DestinationElasticsearchTunnelMethod = "NO_TUNNEL"
)

func (e DestinationElasticsearchTunnelMethod) ToPointer() *DestinationElasticsearchTunnelMethod {
	return &e
}
func (e *DestinationElasticsearchTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationElasticsearchTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchTunnelMethod: %v", v)
	}
}

type DestinationElasticsearchNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationElasticsearchTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationElasticsearchNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchNoTunnel) GetTunnelMethod() DestinationElasticsearchTunnelMethod {
	return DestinationElasticsearchTunnelMethodNoTunnel
}

type DestinationElasticsearchSSHTunnelMethodType string

const (
	DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchNoTunnel               DestinationElasticsearchSSHTunnelMethodType = "destination-elasticsearch_No Tunnel"
	DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchSSHKeyAuthentication   DestinationElasticsearchSSHTunnelMethodType = "destination-elasticsearch_SSH Key Authentication"
	DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchPasswordAuthentication DestinationElasticsearchSSHTunnelMethodType = "destination-elasticsearch_Password Authentication"
)

// DestinationElasticsearchSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationElasticsearchSSHTunnelMethod struct {
	DestinationElasticsearchNoTunnel               *DestinationElasticsearchNoTunnel
	DestinationElasticsearchSSHKeyAuthentication   *DestinationElasticsearchSSHKeyAuthentication
	DestinationElasticsearchPasswordAuthentication *DestinationElasticsearchPasswordAuthentication

	Type DestinationElasticsearchSSHTunnelMethodType
}

func CreateDestinationElasticsearchSSHTunnelMethodDestinationElasticsearchNoTunnel(destinationElasticsearchNoTunnel DestinationElasticsearchNoTunnel) DestinationElasticsearchSSHTunnelMethod {
	typ := DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchNoTunnel

	return DestinationElasticsearchSSHTunnelMethod{
		DestinationElasticsearchNoTunnel: &destinationElasticsearchNoTunnel,
		Type:                             typ,
	}
}

func CreateDestinationElasticsearchSSHTunnelMethodDestinationElasticsearchSSHKeyAuthentication(destinationElasticsearchSSHKeyAuthentication DestinationElasticsearchSSHKeyAuthentication) DestinationElasticsearchSSHTunnelMethod {
	typ := DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchSSHKeyAuthentication

	return DestinationElasticsearchSSHTunnelMethod{
		DestinationElasticsearchSSHKeyAuthentication: &destinationElasticsearchSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationElasticsearchSSHTunnelMethodDestinationElasticsearchPasswordAuthentication(destinationElasticsearchPasswordAuthentication DestinationElasticsearchPasswordAuthentication) DestinationElasticsearchSSHTunnelMethod {
	typ := DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchPasswordAuthentication

	return DestinationElasticsearchSSHTunnelMethod{
		DestinationElasticsearchPasswordAuthentication: &destinationElasticsearchPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationElasticsearchSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationElasticsearchNoTunnel DestinationElasticsearchNoTunnel = DestinationElasticsearchNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchNoTunnel, "", true, true); err == nil {
		u.DestinationElasticsearchNoTunnel = &destinationElasticsearchNoTunnel
		u.Type = DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchNoTunnel
		return nil
	}

	var destinationElasticsearchSSHKeyAuthentication DestinationElasticsearchSSHKeyAuthentication = DestinationElasticsearchSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationElasticsearchSSHKeyAuthentication = &destinationElasticsearchSSHKeyAuthentication
		u.Type = DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchSSHKeyAuthentication
		return nil
	}

	var destinationElasticsearchPasswordAuthentication DestinationElasticsearchPasswordAuthentication = DestinationElasticsearchPasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchPasswordAuthentication, "", true, true); err == nil {
		u.DestinationElasticsearchPasswordAuthentication = &destinationElasticsearchPasswordAuthentication
		u.Type = DestinationElasticsearchSSHTunnelMethodTypeDestinationElasticsearchPasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationElasticsearchSSHTunnelMethod", string(data))
}

func (u DestinationElasticsearchSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationElasticsearchNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchNoTunnel, "", true)
	}

	if u.DestinationElasticsearchSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchSSHKeyAuthentication, "", true)
	}

	if u.DestinationElasticsearchPasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationElasticsearchSSHTunnelMethod: all fields are null")
}

type DestinationElasticsearch struct {
	// The type of authentication to be used
	AuthenticationMethod *DestinationElasticsearchAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// CA certificate
	CaCertificate   *string       `json:"ca_certificate,omitempty"`
	destinationType Elasticsearch `const:"elasticsearch" json:"destinationType"`
	// The full url of the Elasticsearch server
	Endpoint string `json:"endpoint"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationElasticsearchSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// If a primary key identifier is defined in the source, an upsert will be performed using the primary key value as the elasticsearch doc id. Does not support composite primary keys.
	Upsert *bool `default:"true" json:"upsert"`
}

func (d DestinationElasticsearch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearch) GetAuthenticationMethod() *DestinationElasticsearchAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthenticationMethod
}

func (o *DestinationElasticsearch) GetCaCertificate() *string {
	if o == nil {
		return nil
	}
	return o.CaCertificate
}

func (o *DestinationElasticsearch) GetDestinationType() Elasticsearch {
	return ElasticsearchElasticsearch
}

func (o *DestinationElasticsearch) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *DestinationElasticsearch) GetTunnelMethod() *DestinationElasticsearchSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationElasticsearch) GetUpsert() *bool {
	if o == nil {
		return nil
	}
	return o.Upsert
}