---
page_title: "awscc_gamelift_game_session_queue Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.
---

# awscc_gamelift_game_session_queue (Resource)

The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.

## Example Usage
### Basic Example
To use awscc_gamelift_game_session_queue to create a GameLift Game Session Queue:
```terraform
resource "awscc_gamelift_game_session_queue" "example" {
  name               = "ExampleQueue"
  timeout_in_seconds = 600

  destinations = [
    {
      destination_arn = "your-fleet-or-fleet-alias-arn" // ARN of your Fleet or Fleet Alias
    },
  ]

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) A descriptive label that is associated with game session queue. Queue names must be unique within each Region.

### Optional

- `custom_event_data` (String) Information that is added to all events that are related to this game session queue.
- `destinations` (Attributes List) A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue. (see [below for nested schema](#nestedatt--destinations))
- `filter_configuration` (Attributes) A list of locations where a queue is allowed to place new game sessions. (see [below for nested schema](#nestedatt--filter_configuration))
- `notification_target` (String) An SNS topic ARN that is set up to receive game session placement notifications.
- `player_latency_policies` (Attributes List) A set of policies that act as a sliding cap on player latency. (see [below for nested schema](#nestedatt--player_latency_policies))
- `priority_configuration` (Attributes) Custom settings to use when prioritizing destinations and locations for game session placements. (see [below for nested schema](#nestedatt--priority_configuration))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `timeout_in_seconds` (Number) The maximum time, in seconds, that a new game session placement request remains in the queue.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--destinations"></a>
### Nested Schema for `destinations`

Optional:

- `destination_arn` (String)


<a id="nestedatt--filter_configuration"></a>
### Nested Schema for `filter_configuration`

Optional:

- `allowed_locations` (List of String) A list of locations to allow game session placement in, in the form of AWS Region codes such as us-west-2.


<a id="nestedatt--player_latency_policies"></a>
### Nested Schema for `player_latency_policies`

Optional:

- `maximum_individual_player_latency_milliseconds` (Number) The maximum latency value that is allowed for any player, in milliseconds. All policies must have a value set for this property.
- `policy_duration_seconds` (Number) The length of time, in seconds, that the policy is enforced while placing a new game session.


<a id="nestedatt--priority_configuration"></a>
### Nested Schema for `priority_configuration`

Optional:

- `location_order` (List of String) The prioritization order to use for fleet locations, when the PriorityOrder property includes LOCATION.
- `priority_order` (List of String) The recommended sequence to use when prioritizing where to place new game sessions.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_gamelift_game_session_queue.example
  id = "name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_gamelift_game_session_queue.example "name"
```
