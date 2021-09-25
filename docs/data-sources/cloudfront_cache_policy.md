---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_cache_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFront::CachePolicy
---

# awscc_cloudfront_cache_policy (Data Source)

Data Source schema for AWS::CloudFront::CachePolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **cache_policy_config** (Attributes) (see [below for nested schema](#nestedatt--cache_policy_config))
- **last_modified_time** (String)

<a id="nestedatt--cache_policy_config"></a>
### Nested Schema for `cache_policy_config`

Read-Only:

- **comment** (String)
- **default_ttl** (Number)
- **max_ttl** (Number)
- **min_ttl** (Number)
- **name** (String)
- **parameters_in_cache_key_and_forwarded_to_origin** (Attributes) (see [below for nested schema](#nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin))

<a id="nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin"></a>
### Nested Schema for `cache_policy_config.parameters_in_cache_key_and_forwarded_to_origin`

Read-Only:

- **cookies_config** (Attributes) (see [below for nested schema](#nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin--cookies_config))
- **enable_accept_encoding_brotli** (Boolean)
- **enable_accept_encoding_gzip** (Boolean)
- **headers_config** (Attributes) (see [below for nested schema](#nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin--headers_config))
- **query_strings_config** (Attributes) (see [below for nested schema](#nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin--query_strings_config))

<a id="nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin--cookies_config"></a>
### Nested Schema for `cache_policy_config.parameters_in_cache_key_and_forwarded_to_origin.cookies_config`

Read-Only:

- **cookie_behavior** (String)
- **cookies** (List of String)


<a id="nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin--headers_config"></a>
### Nested Schema for `cache_policy_config.parameters_in_cache_key_and_forwarded_to_origin.headers_config`

Read-Only:

- **header_behavior** (String)
- **headers** (List of String)


<a id="nestedatt--cache_policy_config--parameters_in_cache_key_and_forwarded_to_origin--query_strings_config"></a>
### Nested Schema for `cache_policy_config.parameters_in_cache_key_and_forwarded_to_origin.query_strings_config`

Read-Only:

- **query_string_behavior** (String)
- **query_strings** (List of String)

