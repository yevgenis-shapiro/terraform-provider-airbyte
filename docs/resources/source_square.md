---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_square Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSquare Resource
---

# airbyte_source_square (Resource)

SourceSquare Resource

## Example Usage

```terraform
resource "airbyte_source_square" "my_source_square" {
  configuration = {
    credentials = {
      api_key = {
        api_key = "...my_api_key..."
      }
    }
    include_deleted_objects = false
    is_sandbox              = false
    start_date              = "2022-02-02"
  }
  definition_id = "4ba31afc-75cd-4175-a719-6a7541fa2a79"
  name          = "Rhonda Stehr I"
  secret_id     = "...my_secret_id..."
  workspace_id  = "eedefb0b-3778-4954-a670-34f3e4d7496d"
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

Optional:

- `credentials` (Attributes) Choose how to authenticate to Square. (see [below for nested schema](#nestedatt--configuration--credentials))
- `include_deleted_objects` (Boolean) In some streams there is an option to include deleted objects (Items, Categories, Discounts, Taxes). Default: false
- `is_sandbox` (Boolean) Determines whether to use the sandbox or production environment. Default: false
- `start_date` (String) UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. If not set, all data will be replicated. Default: "2021-01-01"

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `api_key` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--api_key))
- `oauth_authentication` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--oauth_authentication))

<a id="nestedatt--configuration--credentials--api_key"></a>
### Nested Schema for `configuration.credentials.api_key`

Required:

- `api_key` (String, Sensitive) The API key for a Square application


<a id="nestedatt--configuration--credentials--oauth_authentication"></a>
### Nested Schema for `configuration.credentials.oauth_authentication`

Required:

- `client_id` (String, Sensitive) The Square-issued ID of your application
- `client_secret` (String, Sensitive) The Square-issued application secret for your application
- `refresh_token` (String, Sensitive) A refresh token generated using the above client ID and secret

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_square.my_airbyte_source_square ""
```