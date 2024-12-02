---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_hardcoded_records Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceHardcodedRecords Resource
---

# airbyte_source_hardcoded_records (Resource)

SourceHardcodedRecords Resource

## Example Usage

```terraform
resource "airbyte_source_hardcoded_records" "my_source_hardcodedrecords" {
  configuration = {
    count = 4
  }
  definition_id = "ca13b1e8-9c14-488f-aa41-1d9d922269c9"
  name          = "Ricardo Gutmann"
  secret_id     = "...my_secret_id..."
  workspace_id  = "0bccdd2e-95af-46ed-bc47-c14160113c2d"
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

- `count` (Number) How many records per stream should be generated. Default: 1000

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_hardcoded_records.my_airbyte_source_hardcoded_records ""
```