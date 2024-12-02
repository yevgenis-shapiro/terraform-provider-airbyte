---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_wikipedia_pageviews Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceWikipediaPageviews Resource
---

# airbyte_source_wikipedia_pageviews (Resource)

SourceWikipediaPageviews Resource

## Example Usage

```terraform
resource "airbyte_source_wikipedia_pageviews" "my_source_wikipediapageviews" {
  configuration = {
    access  = "mobile-web"
    agent   = "all-agents"
    article = "Are_You_the_One%3F"
    country = "FR"
    end     = "...my_end..."
    project = "www.mediawiki.org"
    start   = "...my_start..."
  }
  definition_id = "c999b10a-73ba-40c6-911f-95cb65e559b3"
  name          = "Lawrence Senger"
  secret_id     = "...my_secret_id..."
  workspace_id  = "bd06debd-ebc5-4381-8f44-a1e8a05bc47e"
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

- `access` (String) If you want to filter by access method, use one of desktop, mobile-app or mobile-web. If you are interested in pageviews regardless of access method, use all-access.
- `agent` (String) If you want to filter by agent type, use one of user, automated or spider. If you are interested in pageviews regardless of agent type, use all-agents.
- `article` (String) The title of any article in the specified project. Any spaces should be replaced with underscores. It also should be URI-encoded, so that non-URI-safe characters like %, / or ? are accepted.
- `country` (String) The ISO 3166-1 alpha-2 code of a country for which to retrieve top articles.
- `end` (String) The date of the last day to include, in YYYYMMDD or YYYYMMDDHH format.
- `project` (String) If you want to filter by project, use the domain of any Wikimedia project.
- `start` (String) The date of the first day to include, in YYYYMMDD or YYYYMMDDHH format. Also serves as the date to retrieve data for the top articles.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_wikipedia_pageviews.my_airbyte_source_wikipedia_pageviews ""
```