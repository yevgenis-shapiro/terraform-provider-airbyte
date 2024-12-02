---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_auth0 Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAuth0 DataSource
---

# airbyte_source_auth0 (Data Source)

SourceAuth0 DataSource

## Example Usage

```terraform
data "airbyte_source_auth0" "my_source_auth0" {
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

