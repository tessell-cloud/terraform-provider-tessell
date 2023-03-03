---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tessell_availability_machines Data Source - terraform-provider-tessell"
subcategory: ""
description: |-
  
---

# tessell_availability_machines (Data Source)



## Example Usage

```terraform
# Get all existing Availability Machines
data_source "tessell_availability_machines" "example" {}

# Get all Oracle Availability Machines
data_source "tessell_availability_machines" "example" {
  engine_type = "ORACLE"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `engine_type` (String) Availaility Machine's engine-types
- `load_acls` (Boolean) Load ACL information
- `name` (String) Name of the Availability Machine
- `owners` (List of String) List of Email Addresses for entity or resource owners
- `status` (String) status

### Read-Only

- `availability_machines` (List of Object) (see [below for nested schema](#nestedatt--availability_machines))
- `id` (String) The ID of this resource.

<a id="nestedatt--availability_machines"></a>
### Nested Schema for `availability_machines`

Read-Only:

- `clones` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--clones))
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--cloud_availability))
- `daps` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps))
- `data_ingestion_status` (String)
- `date_created` (String)
- `date_modified` (String)
- `engine_type` (String)
- `id` (String)
- `logged_in_user_role` (String)
- `owner` (String)
- `rpo_sla` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--rpo_sla))
- `service_name` (String)
- `shared_with` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--shared_with))
- `subscription` (String)
- `tenant` (String)
- `tessell_service_id` (String)
- `user_id` (String)

<a id="nestedobjatt--availability_machines--clones"></a>
### Nested Schema for `availability_machines.clones`

Read-Only:

- `clone_info` (Map of String)
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--clones--cloud_availability))
- `date_created` (String)
- `id` (String)
- `name` (String)
- `owner` (String)
- `status` (String)
- `subscription` (String)

<a id="nestedobjatt--availability_machines--clones--cloud_availability"></a>
### Nested Schema for `availability_machines.clones.cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--clones--cloud_availability--regions))

<a id="nestedobjatt--availability_machines--clones--cloud_availability--regions"></a>
### Nested Schema for `availability_machines.clones.cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)




<a id="nestedobjatt--availability_machines--cloud_availability"></a>
### Nested Schema for `availability_machines.cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--cloud_availability--regions))

<a id="nestedobjatt--availability_machines--cloud_availability--regions"></a>
### Nested Schema for `availability_machines.cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)



<a id="nestedobjatt--availability_machines--daps"></a>
### Nested Schema for `availability_machines.daps`

Read-Only:

- `availability_machine_id` (String)
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--cloud_availability))
- `content_info` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--content_info))
- `content_type` (String)
- `data_access_config` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--data_access_config))
- `date_created` (String)
- `date_modified` (String)
- `engine_type` (String)
- `id` (String)
- `logged_in_user_role` (String)
- `name` (String)
- `owner` (String)
- `service_name` (String)
- `shared_with` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--shared_with))
- `status` (String)
- `tessell_service_id` (String)

<a id="nestedobjatt--availability_machines--daps--cloud_availability"></a>
### Nested Schema for `availability_machines.daps.cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--cloud_availability--regions))

<a id="nestedobjatt--availability_machines--daps--cloud_availability--regions"></a>
### Nested Schema for `availability_machines.daps.cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)



<a id="nestedobjatt--availability_machines--daps--content_info"></a>
### Nested Schema for `availability_machines.daps.content_info`

Read-Only:

- `as_is_content` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--content_info--as_is_content))
- `sanitized_content` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--content_info--sanitized_content))

<a id="nestedobjatt--availability_machines--daps--content_info--as_is_content"></a>
### Nested Schema for `availability_machines.daps.content_info.sanitized_content`

Read-Only:

- `automated` (Boolean)
- `manual` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--content_info--sanitized_content--manual))

<a id="nestedobjatt--availability_machines--daps--content_info--sanitized_content--manual"></a>
### Nested Schema for `availability_machines.daps.content_info.sanitized_content.manual`

Read-Only:

- `shared_at` (String)
- `snapshot_id` (String)
- `snapshot_name` (String)
- `snapshot_time` (String)



<a id="nestedobjatt--availability_machines--daps--content_info--sanitized_content"></a>
### Nested Schema for `availability_machines.daps.content_info.sanitized_content`

Read-Only:

- `automated` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--content_info--sanitized_content--automated))
- `manual` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--content_info--sanitized_content--manual))

<a id="nestedobjatt--availability_machines--daps--content_info--sanitized_content--automated"></a>
### Nested Schema for `availability_machines.daps.content_info.sanitized_content.automated`

Read-Only:

- `sanitization_schedule_id` (String)


<a id="nestedobjatt--availability_machines--daps--content_info--sanitized_content--manual"></a>
### Nested Schema for `availability_machines.daps.content_info.sanitized_content.manual`

Read-Only:

- `shared_at` (String)
- `snapshot_id` (String)
- `snapshot_name` (String)
- `snapshot_time` (String)




<a id="nestedobjatt--availability_machines--daps--data_access_config"></a>
### Nested Schema for `availability_machines.daps.data_access_config`

Read-Only:

- `daily_backups` (Number)


<a id="nestedobjatt--availability_machines--daps--shared_with"></a>
### Nested Schema for `availability_machines.daps.shared_with`

Read-Only:

- `users` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--daps--shared_with--users))

<a id="nestedobjatt--availability_machines--daps--shared_with--users"></a>
### Nested Schema for `availability_machines.daps.shared_with.users`

Read-Only:

- `email_id` (String)
- `role` (String)




<a id="nestedobjatt--availability_machines--rpo_sla"></a>
### Nested Schema for `availability_machines.rpo_sla`

Read-Only:

- `availability_machine` (String)
- `availability_machine_id` (String)
- `rpo_sla_status` (String)
- `schedule` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--rpo_sla--schedule))
- `sla` (String)
- `topology` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--rpo_sla--topology))

<a id="nestedobjatt--availability_machines--rpo_sla--schedule"></a>
### Nested Schema for `availability_machines.rpo_sla.schedule`

Read-Only:

- `backup_start_time` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--rpo_sla--schedule--backup_start_time))
- `daily_schedule` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--rpo_sla--schedule--daily_schedule))

<a id="nestedobjatt--availability_machines--rpo_sla--schedule--backup_start_time"></a>
### Nested Schema for `availability_machines.rpo_sla.schedule.daily_schedule`

Read-Only:

- `hour` (Number)
- `minute` (Number)


<a id="nestedobjatt--availability_machines--rpo_sla--schedule--daily_schedule"></a>
### Nested Schema for `availability_machines.rpo_sla.schedule.daily_schedule`

Read-Only:

- `backup_start_times` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--rpo_sla--schedule--daily_schedule--backup_start_times))
- `backups_per_day` (Number)

<a id="nestedobjatt--availability_machines--rpo_sla--schedule--daily_schedule--backup_start_times"></a>
### Nested Schema for `availability_machines.rpo_sla.schedule.daily_schedule.backup_start_times`

Read-Only:

- `hour` (Number)
- `minute` (Number)




<a id="nestedobjatt--availability_machines--rpo_sla--topology"></a>
### Nested Schema for `availability_machines.rpo_sla.topology`

Read-Only:

- `availability_zones` (List of String)
- `cloud_type` (String)
- `region` (String)
- `type` (String)



<a id="nestedobjatt--availability_machines--shared_with"></a>
### Nested Schema for `availability_machines.shared_with`

Read-Only:

- `users` (List of Object) (see [below for nested schema](#nestedobjatt--availability_machines--shared_with--users))

<a id="nestedobjatt--availability_machines--shared_with--users"></a>
### Nested Schema for `availability_machines.shared_with.users`

Read-Only:

- `email_id` (String)
- `role` (String)

