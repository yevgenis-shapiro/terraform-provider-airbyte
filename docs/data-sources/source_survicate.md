---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_survicate Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSurvicate DataSource
---

# airbyte_source_survicate (Data Source)

SourceSurvicate DataSource

## Example Usage

```terraform
data "airbyte_source_survicate" "my_source_survicate" {
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

