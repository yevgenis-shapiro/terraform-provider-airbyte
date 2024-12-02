---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_marketo Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMarketo Resource
---

# airbyte_source_marketo (Resource)

SourceMarketo Resource

## Example Usage

```terraform
resource "airbyte_source_marketo" "my_source_marketo" {
  configuration = {
    client_id     = "...my_client_id..."
    client_secret = "...my_client_secret..."
    domain_url    = "https://000-AAA-000.mktorest.com"
    start_date    = "2020-09-25T00:00:00Z"
  }
  definition_id = "f008f118-d815-472f-b24d-1e0e7e708b9f"
  name          = "Ralph Hermiston"
  secret_id     = "...my_secret_id..."
  workspace_id  = "9f1370c2-8b27-48d2-9e4e-e4a51abe7bbe"
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

- `client_id` (String, Sensitive) The Client ID of your Marketo developer application. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
- `client_secret` (String, Sensitive) The Client Secret of your Marketo developer application. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
- `domain_url` (String, Sensitive) Your Marketo Base URL. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_marketo.my_airbyte_source_marketo ""
```