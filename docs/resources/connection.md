---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_connection Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  Connection Resource
---

# airbyte_connection (Resource)

Connection Resource

## Example Usage

```terraform
resource "airbyte_connection" "my_connection" {
  data_residency                       = "eu"
  destination_id                       = "669dd1e3-6208-43ea-bc85-5914e0a570f6"
  name                                 = "Taylor Hagenes"
  namespace_definition                 = "custom_format"
  namespace_format                     = SOURCE_NAMESPACE
  non_breaking_schema_updates_behavior = "propagate_columns"
  prefix                               = "...my_prefix..."
  source_id                            = "3a555847-8358-4423-a5b6-c7b3fd2fd307"
  status                               = "deprecated"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String) Requires replacement if changed.
- `source_id` (String) Requires replacement if changed.

### Optional

- `configurations` (Attributes) A list of configured stream options for a connection. (see [below for nested schema](#nestedatt--configurations))
- `data_residency` (String) must be one of ["auto", "us", "eu"]; Default: "auto"
- `name` (String) Optional name of the connection
- `namespace_definition` (String) Define the location where the data will be stored in the destination. must be one of ["source", "destination", "custom_format"]; Default: "destination"
- `namespace_format` (String) Used when namespaceDefinition is 'custom_format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'. Default: null
- `non_breaking_schema_updates_behavior` (String) Set how Airbyte handles syncs when it detects a non-breaking schema change in the source. must be one of ["ignore", "disable_connection", "propagate_columns", "propagate_fully"]; Default: "ignore"
- `prefix` (String) Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte_” causes “projects” => “airbyte_projects”).
- `schedule` (Attributes) schedule for when the the connection should run, per the schedule type (see [below for nested schema](#nestedatt--schedule))
- `status` (String) must be one of ["active", "inactive", "deprecated"]

### Read-Only

- `connection_id` (String)
- `workspace_id` (String)

<a id="nestedatt--configurations"></a>
### Nested Schema for `configurations`

Optional:

- `streams` (Attributes Set) (see [below for nested schema](#nestedatt--configurations--streams))

<a id="nestedatt--configurations--streams"></a>
### Nested Schema for `configurations.streams`

Optional:

- `cursor_field` (List of String) Path to the field that will be used to determine if a record is new or modified since the last sync. This field is REQUIRED if `sync_mode` is `incremental` unless there is a default.
- `name` (String) Not Null
- `primary_key` (List of List of String) Paths to the fields that will be used as primary key. This field is REQUIRED if `destination_sync_mode` is `*_dedup` unless it is already supplied by the source schema.
- `selected_fields` (Attributes List) Paths to the fields that will be included in the configured catalog. (see [below for nested schema](#nestedatt--configurations--streams--selected_fields))
- `sync_mode` (String) must be one of ["full_refresh_overwrite", "full_refresh_append", "incremental_append", "incremental_deduped_history"]

<a id="nestedatt--configurations--streams--selected_fields"></a>
### Nested Schema for `configurations.streams.selected_fields`

Optional:

- `field_path` (List of String)




<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Optional:

- `cron_expression` (String)
- `schedule_type` (String) Not Null; must be one of ["manual", "cron"]

Read-Only:

- `basic_timing` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_connection.my_airbyte_connection ""
```