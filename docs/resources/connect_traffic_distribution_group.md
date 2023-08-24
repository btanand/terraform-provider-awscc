---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_traffic_distribution_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Connect::TrafficDistributionGroup
---

# awscc_connect_traffic_distribution_group (Resource)

Resource Type definition for AWS::Connect::TrafficDistributionGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_arn` (String) The identifier of the Amazon Connect instance that has been replicated.
- `name` (String) The name for the traffic distribution group.

### Optional

- `description` (String) A description for the traffic distribution group.
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `is_default` (Boolean) If this is the default traffic distribution group.
- `status` (String) The status of the traffic distribution group.
- `traffic_distribution_group_arn` (String) The identifier of the traffic distribution group.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_connect_traffic_distribution_group.example <resource ID>
```