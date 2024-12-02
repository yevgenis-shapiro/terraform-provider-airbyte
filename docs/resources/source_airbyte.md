---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_airbyte Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAirbyte Resource
---

# airbyte_source_airbyte (Resource)

SourceAirbyte Resource

## Example Usage

```terraform
resource "airbyte_source_airbyte" "my_source_airbyte" {
  configuration = {
    client_id     = "...my_client_id..."
    client_secret = "...my_client_secret..."
    start_date    = "2022-03-19T07:46:59.910Z"
  }
  definition_id = "0dabba6e-f9fc-43c3-b44f-d252e57aa673"
  name          = "Steve Luettgen"
  secret_id     = "...my_secret_id..."
  workspace_id  = "9f0c220e-39e1-40d6-af09-fb849b0bdf3d"
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

- `client_id` (String)
- `client_secret` (String, Sensitive)
- `start_date` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_airbyte.my_airbyte_source_airbyte ""
```