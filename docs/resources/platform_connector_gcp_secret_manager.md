---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_gcp_secret_manager Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a GCP Secret Manager connector.
---

# harness_platform_connector_gcp_secret_manager (Resource)

Resource for creating a GCP Secret Manager connector.

## Example Usage

```terraform
resource "harness_platform_connector_gcp_secret_manager" "gcp_sm" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  credentials_ref    = "account.${harness_platform_secret_text.test.id}"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credentials_ref` (String) Reference to the secret containing credentials of IAM service account for Google Secret Manager. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `is_default` (Boolean) Indicative if this is default Secret manager for secrets.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import account level gcp secret manager connector 
terraform import harness_platform_connector_gcp_secret_manager.example <connector_id>

# Import org level gcp secret manager connector 
terraform import harness_platform_connector_gcp_secret_manager.example <ord_id>/<connector_id>

# Import project level gcp secret manager connector 
terraform import harness_platform_connector_gcp_secret_manager.example <org_id>/<project_id>/<connector_id>
```
