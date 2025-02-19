---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_jira Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Jira connector.
---

# harness_platform_connector_jira (Resource)

Resource for creating a Jira connector.

## Example Usage

```terraform
resource "harness_platform_connector_jira" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://jira.com"
  delegate_selectors = ["harness-delegate"]
  username           = "admin"
  password_ref       = "account.secret_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `password_ref` (String) Reference to a secret containing the password to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `url` (String) URL of the Jira server.

### Optional

- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.
- `username` (String) Username to use for authentication.
- `username_ref` (String) Reference to a secret containing the username to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import account level jira connector 
terraform import harness_platform_connector_jira.example <connector_id>

# Import org level jira connector 
terraform import harness_platform_connector_jira.example <ord_id>/<connector_id>

# Import project level jira connector 
terraform import harness_platform_connector_jira.example <org_id>/<project_id>/<connector_id>
```
