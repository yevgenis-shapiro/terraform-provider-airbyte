---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_sftp Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSftp Resource
---

# airbyte_source_sftp (Resource)

SourceSftp Resource

## Example Usage

```terraform
resource "airbyte_source_sftp" "my_source_sftp" {
  configuration = {
    credentials = {
      password_authentication = {
        auth_user_password = "...my_auth_user_password..."
      }
    }
    file_pattern = "log-([0-9]{4})([0-9]{2})([0-9]{2}) - This will filter files which  `log-yearmmdd`"
    file_types   = "csv"
    folder_path  = "/logs/2022"
    host         = "192.0.2.1"
    port         = 22
    user         = "...my_user..."
  }
  definition_id = "df9bf84b-784e-4daa-b2f4-24ed308606f0"
  name          = "Wendell Dare"
  secret_id     = "...my_secret_id..."
  workspace_id  = "2de7b1a9-3e59-415f-a584-4c8d7f9e67ba"
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

Required:

- `host` (String) The server host address
- `user` (String) The server user

Optional:

- `credentials` (Attributes) The server authentication method (see [below for nested schema](#nestedatt--configuration--credentials))
- `file_pattern` (String) The regular expression to specify files for sync in a chosen Folder Path. Default: ""
- `file_types` (String) Coma separated file types. Currently only 'csv' and 'json' types are supported. Default: "csv,json"
- `folder_path` (String) The directory to search files for sync. Default: ""
- `port` (Number) The server port. Default: 22

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `password_authentication` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--password_authentication))
- `ssh_key_authentication` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--ssh_key_authentication))

<a id="nestedatt--configuration--credentials--password_authentication"></a>
### Nested Schema for `configuration.credentials.password_authentication`

Required:

- `auth_user_password` (String, Sensitive) OS-level password for logging into the jump server host


<a id="nestedatt--configuration--credentials--ssh_key_authentication"></a>
### Nested Schema for `configuration.credentials.ssh_key_authentication`

Required:

- `auth_ssh_key` (String, Sensitive) OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_sftp.my_airbyte_source_sftp ""
```