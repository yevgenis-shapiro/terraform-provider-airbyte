---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_braintree Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceBraintree Resource
---

# airbyte_source_braintree (Resource)

SourceBraintree Resource

## Example Usage

```terraform
resource "airbyte_source_braintree" "my_source_braintree" {
  configuration = {
    environment = "Production"
    merchant_id = "...my_merchant_id..."
    private_key = "...my_private_key..."
    public_key  = "...my_public_key..."
    start_date  = "2020"
  }
  definition_id = "a849cd51-7f0e-432c-ae30-2e6de92b3e43"
  name          = "Daisy MacGyver"
  secret_id     = "...my_secret_id..."
  workspace_id  = "6f9835e6-422d-415b-8286-21a877d2e625"
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

- `environment` (String) Environment specifies where the data will come from. must be one of ["Development", "Sandbox", "Qa", "Production"]
- `merchant_id` (String) The unique identifier for your entire gateway account. See the <a href="https://docs.airbyte.com/integrations/sources/braintree">docs</a> for more information on how to obtain this ID.
- `private_key` (String, Sensitive) Braintree Private Key. See the <a href="https://docs.airbyte.com/integrations/sources/braintree">docs</a> for more information on how to obtain this key.
- `public_key` (String) Braintree Public Key. See the <a href="https://docs.airbyte.com/integrations/sources/braintree">docs</a> for more information on how to obtain this key.

Optional:

- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_braintree.my_airbyte_source_braintree ""
```