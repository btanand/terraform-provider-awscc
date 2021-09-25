---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_globalaccelerator_endpoint_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::GlobalAccelerator::EndpointGroup
---

# awscc_globalaccelerator_endpoint_group (Data Source)

Data Source schema for AWS::GlobalAccelerator::EndpointGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **endpoint_configurations** (Attributes List) The list of endpoint objects. (see [below for nested schema](#nestedatt--endpoint_configurations))
- **endpoint_group_arn** (String) The Amazon Resource Name (ARN) of the endpoint group
- **endpoint_group_region** (String) The name of the AWS Region where the endpoint group is located
- **health_check_interval_seconds** (Number) The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
- **health_check_path** (String)
- **health_check_port** (Number) The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
- **health_check_protocol** (String) The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
- **listener_arn** (String) The Amazon Resource Name (ARN) of the listener
- **port_overrides** (Attributes List) (see [below for nested schema](#nestedatt--port_overrides))
- **threshold_count** (Number) The number of consecutive health checks required to set the state of the endpoint to unhealthy.
- **traffic_dial_percentage** (Number) The percentage of traffic to sent to an AWS Region

<a id="nestedatt--endpoint_configurations"></a>
### Nested Schema for `endpoint_configurations`

Read-Only:

- **client_ip_preservation_enabled** (Boolean) true if client ip should be preserved
- **endpoint_id** (String) Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID
- **weight** (Number) The weight for the endpoint.


<a id="nestedatt--port_overrides"></a>
### Nested Schema for `port_overrides`

Read-Only:

- **endpoint_port** (Number) A network port number
- **listener_port** (Number) A network port number

