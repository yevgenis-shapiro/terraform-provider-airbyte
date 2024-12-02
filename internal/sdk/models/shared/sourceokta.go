// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceOktaSchemasCredentialsAuthType string

const (
	SourceOktaSchemasCredentialsAuthTypeAPIToken SourceOktaSchemasCredentialsAuthType = "api_token"
)

func (e SourceOktaSchemasCredentialsAuthType) ToPointer() *SourceOktaSchemasCredentialsAuthType {
	return &e
}
func (e *SourceOktaSchemasCredentialsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceOktaSchemasCredentialsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOktaSchemasCredentialsAuthType: %v", v)
	}
}

type SourceOktaAPIToken struct {
	// An Okta token. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to generate it.
	APIToken string                               `json:"api_token"`
	authType SourceOktaSchemasCredentialsAuthType `const:"api_token" json:"auth_type"`
}

func (s SourceOktaAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOktaAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOktaAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceOktaAPIToken) GetAuthType() SourceOktaSchemasCredentialsAuthType {
	return SourceOktaSchemasCredentialsAuthTypeAPIToken
}

type SourceOktaSchemasAuthType string

const (
	SourceOktaSchemasAuthTypeOauth20PrivateKey SourceOktaSchemasAuthType = "oauth2.0_private_key"
)

func (e SourceOktaSchemasAuthType) ToPointer() *SourceOktaSchemasAuthType {
	return &e
}
func (e *SourceOktaSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0_private_key":
		*e = SourceOktaSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOktaSchemasAuthType: %v", v)
	}
}

type SourceOktaOAuth20WithPrivateKey struct {
	authType SourceOktaSchemasAuthType `const:"oauth2.0_private_key" json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The key ID (kid).
	KeyID string `json:"key_id"`
	// The private key in PEM format
	PrivateKey string `json:"private_key"`
	// The OAuth scope.
	Scope string `json:"scope"`
}

func (s SourceOktaOAuth20WithPrivateKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOktaOAuth20WithPrivateKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOktaOAuth20WithPrivateKey) GetAuthType() SourceOktaSchemasAuthType {
	return SourceOktaSchemasAuthTypeOauth20PrivateKey
}

func (o *SourceOktaOAuth20WithPrivateKey) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceOktaOAuth20WithPrivateKey) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *SourceOktaOAuth20WithPrivateKey) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *SourceOktaOAuth20WithPrivateKey) GetScope() string {
	if o == nil {
		return ""
	}
	return o.Scope
}

type SourceOktaAuthType string

const (
	SourceOktaAuthTypeOauth20 SourceOktaAuthType = "oauth2.0"
)

func (e SourceOktaAuthType) ToPointer() *SourceOktaAuthType {
	return &e
}
func (e *SourceOktaAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceOktaAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOktaAuthType: %v", v)
	}
}

type SourceOktaOAuth20 struct {
	authType SourceOktaAuthType `const:"oauth2.0" json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceOktaOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOktaOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOktaOAuth20) GetAuthType() SourceOktaAuthType {
	return SourceOktaAuthTypeOauth20
}

func (o *SourceOktaOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceOktaOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceOktaOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceOktaAuthorizationMethodType string

const (
	SourceOktaAuthorizationMethodTypeSourceOktaOAuth20               SourceOktaAuthorizationMethodType = "source-okta_OAuth2.0"
	SourceOktaAuthorizationMethodTypeSourceOktaOAuth20WithPrivateKey SourceOktaAuthorizationMethodType = "source-okta_OAuth 2.0 with private key"
	SourceOktaAuthorizationMethodTypeSourceOktaAPIToken              SourceOktaAuthorizationMethodType = "source-okta_API Token"
)

type SourceOktaAuthorizationMethod struct {
	SourceOktaOAuth20               *SourceOktaOAuth20
	SourceOktaOAuth20WithPrivateKey *SourceOktaOAuth20WithPrivateKey
	SourceOktaAPIToken              *SourceOktaAPIToken

	Type SourceOktaAuthorizationMethodType
}

func CreateSourceOktaAuthorizationMethodSourceOktaOAuth20(sourceOktaOAuth20 SourceOktaOAuth20) SourceOktaAuthorizationMethod {
	typ := SourceOktaAuthorizationMethodTypeSourceOktaOAuth20

	return SourceOktaAuthorizationMethod{
		SourceOktaOAuth20: &sourceOktaOAuth20,
		Type:              typ,
	}
}

func CreateSourceOktaAuthorizationMethodSourceOktaOAuth20WithPrivateKey(sourceOktaOAuth20WithPrivateKey SourceOktaOAuth20WithPrivateKey) SourceOktaAuthorizationMethod {
	typ := SourceOktaAuthorizationMethodTypeSourceOktaOAuth20WithPrivateKey

	return SourceOktaAuthorizationMethod{
		SourceOktaOAuth20WithPrivateKey: &sourceOktaOAuth20WithPrivateKey,
		Type:                            typ,
	}
}

func CreateSourceOktaAuthorizationMethodSourceOktaAPIToken(sourceOktaAPIToken SourceOktaAPIToken) SourceOktaAuthorizationMethod {
	typ := SourceOktaAuthorizationMethodTypeSourceOktaAPIToken

	return SourceOktaAuthorizationMethod{
		SourceOktaAPIToken: &sourceOktaAPIToken,
		Type:               typ,
	}
}

func (u *SourceOktaAuthorizationMethod) UnmarshalJSON(data []byte) error {

	var sourceOktaAPIToken SourceOktaAPIToken = SourceOktaAPIToken{}
	if err := utils.UnmarshalJSON(data, &sourceOktaAPIToken, "", true, true); err == nil {
		u.SourceOktaAPIToken = &sourceOktaAPIToken
		u.Type = SourceOktaAuthorizationMethodTypeSourceOktaAPIToken
		return nil
	}

	var sourceOktaOAuth20 SourceOktaOAuth20 = SourceOktaOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceOktaOAuth20, "", true, true); err == nil {
		u.SourceOktaOAuth20 = &sourceOktaOAuth20
		u.Type = SourceOktaAuthorizationMethodTypeSourceOktaOAuth20
		return nil
	}

	var sourceOktaOAuth20WithPrivateKey SourceOktaOAuth20WithPrivateKey = SourceOktaOAuth20WithPrivateKey{}
	if err := utils.UnmarshalJSON(data, &sourceOktaOAuth20WithPrivateKey, "", true, true); err == nil {
		u.SourceOktaOAuth20WithPrivateKey = &sourceOktaOAuth20WithPrivateKey
		u.Type = SourceOktaAuthorizationMethodTypeSourceOktaOAuth20WithPrivateKey
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceOktaAuthorizationMethod", string(data))
}

func (u SourceOktaAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceOktaOAuth20 != nil {
		return utils.MarshalJSON(u.SourceOktaOAuth20, "", true)
	}

	if u.SourceOktaOAuth20WithPrivateKey != nil {
		return utils.MarshalJSON(u.SourceOktaOAuth20WithPrivateKey, "", true)
	}

	if u.SourceOktaAPIToken != nil {
		return utils.MarshalJSON(u.SourceOktaAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceOktaAuthorizationMethod: all fields are null")
}

type Okta string

const (
	OktaOkta Okta = "okta"
)

func (e Okta) ToPointer() *Okta {
	return &e
}
func (e *Okta) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "okta":
		*e = Okta(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Okta: %v", v)
	}
}

type SourceOkta struct {
	Credentials *SourceOktaAuthorizationMethod `json:"credentials,omitempty"`
	// The Okta domain. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to find it.
	Domain     *string `json:"domain,omitempty"`
	sourceType Okta    `const:"okta" json:"sourceType"`
	// UTC date and time in the format YYYY-MM-DDTHH:MM:SSZ. Any data before this date will not be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceOkta) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOkta) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOkta) GetCredentials() *SourceOktaAuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceOkta) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *SourceOkta) GetSourceType() Okta {
	return OktaOkta
}

func (o *SourceOkta) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}