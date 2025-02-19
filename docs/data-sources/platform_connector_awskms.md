---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_awskms Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Datasource for looking up an AWS KMS connector.
---

# harness_platform_connector_awskms (Data Source)

Datasource for looking up an AWS KMS connector.

## Example Usage

```terraform
data "harness_platform_connector_awskms" "example" {
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

- `arn_ref` (String) A reference to the Harness secret containing the ARN of the AWS KMS. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `credentials` (List of Object) Credentials to connect to AWS. (see [below for nested schema](#nestedatt--credentials))
- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `region` (String) The AWS region where the AWS Secret Manager is.
- `tags` (Set of String) Tags to associate with the resource.

<a id="nestedatt--credentials"></a>
### Nested Schema for `credentials`

Read-Only:

- `assume_role` (List of Object) (see [below for nested schema](#nestedobjatt--credentials--assume_role))
- `inherit_from_delegate` (Boolean)
- `manual` (List of Object) (see [below for nested schema](#nestedobjatt--credentials--manual))

<a id="nestedobjatt--credentials--assume_role"></a>
### Nested Schema for `credentials.assume_role`

Read-Only:

- `duration` (Number)
- `external_id` (String)
- `role_arn` (String)


<a id="nestedobjatt--credentials--manual"></a>
### Nested Schema for `credentials.manual`

Read-Only:

- `access_key_ref` (String)
- `secret_key_ref` (String)


