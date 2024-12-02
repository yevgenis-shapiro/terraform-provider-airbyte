---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_planhat Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePlanhat Resource
---

# airbyte_source_planhat (Resource)

SourcePlanhat Resource

## Example Usage

```terraform
resource "airbyte_source_planhat" "my_source_planhat" {
  configuration = {
    api_token = "...my_api_token..."
  }
  definition_id = "a6744848-ac2b-404b-aae9-e175304065f6"
  name          = "Tara King"
  secret_id     = "...my_secret_id..."
  workspace_id  = "901f87c9-df1a-4f8f-9013-d5d0cf403b28"
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

- `api_token` (String, Sensitive) Your Planhat <a href="https://docs.planhat.com/#authentication">API Access Token</a>

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_planhat.my_airbyte_source_planhat ""
```