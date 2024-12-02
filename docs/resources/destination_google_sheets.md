---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_google_sheets Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationGoogleSheets Resource
---

# airbyte_destination_google_sheets (Resource)

DestinationGoogleSheets Resource

## Example Usage

```terraform
resource "airbyte_destination_google_sheets" "my_destination_googlesheets" {
  configuration = {
    credentials = {
      client_id     = "...my_client_id..."
      client_secret = "...my_client_secret..."
      refresh_token = "...my_refresh_token..."
    }
    spreadsheet_id = "https://docs.google.com/spreadsheets/d/1hLd9Qqti3UyLXZB2aFfUWDT7BG/edit"
  }
  definition_id = "64a237e4-a59e-47bf-91d4-96bd14d08d4a"
  name          = "Marcella Hand"
  workspace_id  = "b2ee153b-42c3-42f4-8f6e-543a0f0f39a6"
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

- `credentials` (Attributes) Google API Credentials for connecting to Google Sheets and Google Drive APIs (see [below for nested schema](#nestedatt--configuration--credentials))
- `spreadsheet_id` (String) The link to your spreadsheet. See <a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'>this guide</a> for more details.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Required:

- `client_id` (String, Sensitive) The Client ID of your Google Sheets developer application.
- `client_secret` (String, Sensitive) The Client Secret of your Google Sheets developer application.
- `refresh_token` (String, Sensitive) The token for obtaining new access token.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_destination_google_sheets.my_airbyte_destination_google_sheets ""
```