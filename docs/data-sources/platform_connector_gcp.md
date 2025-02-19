---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_gcp Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Datasource for looking up a Gcp connector.
---

# harness_platform_connector_gcp (Data Source)

Datasource for looking up a Gcp connector.

## Example Usage

```terraform
data "harness_platform_connector_gcp" "example" {
  identifier = "identifier"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Read-Only

- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `inherit_from_delegate` (List of Object) Inherit configuration from delegate. (see [below for nested schema](#nestedatt--inherit_from_delegate))
- `manual` (List of Object) Manual credential configuration. (see [below for nested schema](#nestedatt--manual))
- `tags` (Set of String) Tags to associate with the resource.

<a id="nestedatt--inherit_from_delegate"></a>
### Nested Schema for `inherit_from_delegate`

Read-Only:

- `delegate_selectors` (Set of String)


<a id="nestedatt--manual"></a>
### Nested Schema for `manual`

Read-Only:

- `delegate_selectors` (Set of String)
- `secret_key_ref` (String)


