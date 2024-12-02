---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_trustpilot Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceTrustpilot Resource
---

# airbyte_source_trustpilot (Resource)

SourceTrustpilot Resource

## Example Usage

```terraform
resource "airbyte_source_trustpilot" "my_source_trustpilot" {
  configuration = {
    business_units = [
      "...",
    ]
    credentials = {
      api_key = {
        client_id = "...my_client_id..."
      }
    }
    start_date = "%Y-%m-%dT%H:%M:%SZ"
  }
  definition_id = "a76308e2-8188-410e-b61b-40ea12469466"
  name          = "Crystal Rutherford"
  secret_id     = "...my_secret_id..."
  workspace_id  = "f85dbdf8-ab2f-4267-87a5-b6047904d383"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `business_units` (List of String) The names of business units which shall be synchronized. Some streams e.g. configured_business_units or private_reviews use this configuration.
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `start_date` (String) For streams with sync. method incremental the start date time to be used

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `api_key` (Attributes) The API key authentication method gives you access to only the streams which are part of the Public API. When you want to get streams available via the Consumer API (e.g. the private reviews) you need to use authentication method OAuth 2.0. (see [below for nested schema](#nestedatt--configuration--credentials--api_key))
- `o_auth20` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--o_auth20))

<a id="nestedatt--configuration--credentials--api_key"></a>
### Nested Schema for `configuration.credentials.api_key`

Required:

- `client_id` (String, Sensitive) The API key of the Trustpilot API application.


<a id="nestedatt--configuration--credentials--o_auth20"></a>
### Nested Schema for `configuration.credentials.o_auth20`

Required:

- `access_token` (String, Sensitive) Access Token for making authenticated requests.
- `client_id` (String, Sensitive) The API key of the Trustpilot API application. (represents the OAuth Client ID)
- `client_secret` (String, Sensitive) The Secret of the Trustpilot API application. (represents the OAuth Client Secret)
- `refresh_token` (String, Sensitive) The key to refresh the expired access_token.
- `token_expiry_date` (String) The date-time when the access token should be refreshed.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_trustpilot.my_airbyte_source_trustpilot ""
```