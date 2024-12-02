---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_paystack Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePaystack DataSource
---

# airbyte_source_paystack (Data Source)

SourcePaystack DataSource

## Example Usage

```terraform
data "airbyte_source_paystack" "my_source_paystack" {
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

