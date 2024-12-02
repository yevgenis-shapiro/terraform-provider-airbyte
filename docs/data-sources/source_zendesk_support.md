---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_zendesk_support Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceZendeskSupport DataSource
---

# airbyte_source_zendesk_support (Data Source)

SourceZendeskSupport DataSource

## Example Usage

```terraform
data "airbyte_source_zendesk_support" "my_source_zendesksupport" {
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Read-Only

- `configuration` (String) The values required to configure the source. Parsed as JSON.
- `name` (String)
- `source_type` (String)
- `workspace_id` (String)

