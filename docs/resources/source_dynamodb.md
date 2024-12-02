---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_dynamodb Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceDynamodb Resource
---

# airbyte_source_dynamodb (Resource)

SourceDynamodb Resource

## Example Usage

```terraform
resource "airbyte_source_dynamodb" "my_source_dynamodb" {
  configuration = {
    credentials = {
      authenticate_via_access_keys = {
        access_key_id         = "A012345678910EXAMPLE"
        additional_properties = "{ \"see\": \"documentation\" }"
        secret_access_key     = "a012345678910ABCDEFGH/AbCdEfGhEXAMPLEKEY"
      }
    }
    endpoint                               = "https://{aws_dynamo_db_url}.com"
    ignore_missing_read_permissions_tables = true
    region                                 = "il-central-1"
    reserved_attribute_names               = "name, field_name, field-name"
  }
  definition_id = "ad00caee-12c4-4e65-b57e-54a27b617a01"
  name          = "Lorena Huel"
  secret_id     = "...my_secret_id..."
  workspace_id  = "68e1922d-f283-4a61-8313-a52314031fd7"
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

- `credentials` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials))
- `endpoint` (String) the URL of the Dynamodb database. Default: ""
- `ignore_missing_read_permissions_tables` (Boolean) Ignore tables with missing scan/read permissions. Default: false
- `region` (String) The region of the Dynamodb database. must be one of ["", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-4", "ca-central-1", "ca-west-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-central-2", "eu-north-1", "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3", "il-central-1", "me-central-1", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"]; Default: ""
- `reserved_attribute_names` (String, Sensitive) Comma separated reserved attribute names present in your tables

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `authenticate_via_access_keys` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--authenticate_via_access_keys))
- `role_based_authentication` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--role_based_authentication))

<a id="nestedatt--configuration--credentials--authenticate_via_access_keys"></a>
### Nested Schema for `configuration.credentials.authenticate_via_access_keys`

Required:

- `access_key_id` (String, Sensitive) The access key id to access Dynamodb. Airbyte requires read permissions to the database
- `secret_access_key` (String, Sensitive) The corresponding secret to the access key id.

Optional:

- `additional_properties` (String) Parsed as JSON.


<a id="nestedatt--configuration--credentials--role_based_authentication"></a>
### Nested Schema for `configuration.credentials.role_based_authentication`

Optional:

- `additional_properties` (String) Parsed as JSON.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_dynamodb.my_airbyte_source_dynamodb ""
```