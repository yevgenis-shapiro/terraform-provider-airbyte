---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_postgres Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePostgres DataSource
---

# airbyte_source_postgres (Data Source)

SourcePostgres DataSource

## Example Usage

```terraform
data "airbyte_source_postgres" "my_source_postgres" {
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

