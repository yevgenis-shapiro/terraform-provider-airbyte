---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_azure_blob_storage Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationAzureBlobStorage Resource
---

# airbyte_destination_azure_blob_storage (Resource)

DestinationAzureBlobStorage Resource

## Example Usage

```terraform
resource "airbyte_destination_azure_blob_storage" "my_destination_azureblobstorage" {
  configuration = {
    azure_blob_storage_account_key          = "Z8ZkZpteggFx394vm+PJHnGTvdRncaYS+JhLKdj789YNmD+iyGTnG+PV+POiuYNhBg/ACS+LKjd%4FG3FHGN12Nd=="
    azure_blob_storage_account_name         = "airbyte5storage"
    azure_blob_storage_container_name       = "airbytetescontainername"
    azure_blob_storage_endpoint_domain_name = "blob.core.windows.net"
    azure_blob_storage_output_buffer_size   = 5
    azure_blob_storage_spill_size           = 500
    format = {
      csv_comma_separated_values = {
        file_extension = true
        flattening     = "No flattening"
      }
    }
  }
  definition_id = "63ca2e92-d142-4842-85e9-75e40d11a3c6"
  name          = "Virgil Dickens"
  workspace_id  = "c05b91a7-2d27-400d-8d43-ac809ede88b1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the destination e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `azure_blob_storage_account_key` (String, Sensitive) The Azure blob storage account key.
- `azure_blob_storage_account_name` (String) The account's name of the Azure Blob Storage.
- `format` (Attributes) Output data format (see [below for nested schema](#nestedatt--configuration--format))

Optional:

- `azure_blob_storage_container_name` (String) The name of the Azure blob storage container. If not exists - will be created automatically. May be empty, then will be created automatically airbytecontainer+timestamp
- `azure_blob_storage_endpoint_domain_name` (String) This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example. Default: "blob.core.windows.net"
- `azure_blob_storage_output_buffer_size` (Number) The amount of megabytes to buffer for the output stream to Azure. This will impact memory footprint on workers, but may need adjustment for performance and appropriate block size in Azure. Default: 5
- `azure_blob_storage_spill_size` (Number) The amount of megabytes after which the connector should spill the records in a new blob object. Make sure to configure size greater than individual records. Enter 0 if not applicable. Default: 500

<a id="nestedatt--configuration--format"></a>
### Nested Schema for `configuration.format`

Optional:

- `csv_comma_separated_values` (Attributes) (see [below for nested schema](#nestedatt--configuration--format--csv_comma_separated_values))
- `json_lines_newline_delimited_json` (Attributes) (see [below for nested schema](#nestedatt--configuration--format--json_lines_newline_delimited_json))

<a id="nestedatt--configuration--format--csv_comma_separated_values"></a>
### Nested Schema for `configuration.format.csv_comma_separated_values`

Optional:

- `file_extension` (Boolean) Add file extensions to the output file. Default: false
- `flattening` (String) Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details. must be one of ["No flattening", "Root level flattening"]; Default: "No flattening"


<a id="nestedatt--configuration--format--json_lines_newline_delimited_json"></a>
### Nested Schema for `configuration.format.json_lines_newline_delimited_json`

Optional:

- `file_extension` (Boolean) Add file extensions to the output file. Default: false

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_destination_azure_blob_storage.my_airbyte_destination_azure_blob_storage ""
```