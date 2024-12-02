---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_paypal_transaction Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePaypalTransaction Resource
---

# airbyte_source_paypal_transaction (Resource)

SourcePaypalTransaction Resource

## Example Usage

```terraform
resource "airbyte_source_paypal_transaction" "my_source_paypaltransaction" {
  configuration = {
    client_id          = "...my_client_id..."
    client_secret      = "...my_client_secret..."
    dispute_start_date = "2021-06-11T23:59:59.000Z"
    end_date           = "2021-06-11T23:59:59+00:00"
    is_sandbox         = false
    refresh_token      = "...my_refresh_token..."
    start_date         = "2021-06-11T23:59:59+00:00"
    time_window        = 9
  }
  definition_id = "10910de8-7dfe-4701-adbd-0d10cf57eb67"
  name          = "Melody Lowe"
  secret_id     = "...my_secret_id..."
  workspace_id  = "55d63fb2-a63d-4a09-97a6-151fac3e8ec6"
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

- `client_id` (String, Sensitive) The Client ID of your Paypal developer application.
- `client_secret` (String, Sensitive) The Client Secret of your Paypal developer application.
- `start_date` (String) Start Date for data extraction in <a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\">ISO format</a>. Date must be in range from 3 years till 12 hrs before present time.

Optional:

- `dispute_start_date` (String) Start Date parameter for the list dispute endpoint in <a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\">ISO format</a>. This Start Date must be in range within 180 days before present time, and requires ONLY 3 miliseconds(mandatory). If you don't use this option, it defaults to a start date set 180 days in the past.
- `end_date` (String) End Date for data extraction in <a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\">ISO format</a>. This can be help you select specific range of time, mainly for test purposes  or data integrity tests. When this is not used, now_utc() is used by the streams. This does not apply to Disputes and Product streams.
- `is_sandbox` (Boolean) Determines whether to use the sandbox or production environment. Default: false
- `refresh_token` (String, Sensitive) The key to refresh the expired access token.
- `time_window` (Number) The number of days per request. Must be a number between 1 and 31. Default: 7

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_paypal_transaction.my_airbyte_source_paypal_transaction ""
```