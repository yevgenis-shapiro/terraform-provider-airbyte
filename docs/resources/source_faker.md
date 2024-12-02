---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_faker Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceFaker Resource
---

# airbyte_source_faker (Resource)

SourceFaker Resource

## Example Usage

```terraform
resource "airbyte_source_faker" "my_source_faker" {
  configuration = {
    always_updated    = true
    count             = 8
    parallelism       = 2
    records_per_slice = 0
    seed              = 0
  }
  definition_id = "34e58876-cb03-40a1-a8ae-06a57c7c577a"
  name          = "Joe Weber"
  secret_id     = "...my_secret_id..."
  workspace_id  = "addd2747-bbc7-4f24-9709-ce4fe165bc48"
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

- `always_updated` (Boolean) Should the updated_at values for every record be new each sync?  Setting this to false will case the source to stop emitting records after COUNT records have been emitted. Default: true
- `count` (Number) How many users should be generated in total. The purchases table will be scaled to match, with 10 purchases created per 10 users. This setting does not apply to the products stream. Default: 1000
- `parallelism` (Number) How many parallel workers should we use to generate fake data?  Choose a value equal to the number of CPUs you will allocate to this source. Default: 4
- `records_per_slice` (Number) How many fake records will be in each page (stream slice), before a state message is emitted?. Default: 1000
- `seed` (Number) Manually control the faker random seed to return the same values on subsequent runs (leave -1 for random). Default: -1

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_faker.my_airbyte_source_faker ""
```