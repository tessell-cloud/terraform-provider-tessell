---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tessell_availability_machine Data Source - terraform-provider-tessell"
subcategory: ""
description: |-
  
---

# tessell_availability_machine (Data Source)

The management of snapshot and data is abstracted as a construct called Availability Machine (AM). The 'snapshots' are made availale under the respective Availability Machine and the life-cycle-management (create, delete, replicate) for snapshots would happen under the Availability Mahchine.

## Example Usage

```terraform
# Get an Availability Machine
data_source "tessell_availability_machine" "example" {
  id = "432bc899-b4b9-48db-a580-517f015188da"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `clones` (List of Object) Clone databases that are created from this Availability Machine (see [below for nested schema](#nestedatt--clones))
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedatt--cloud_availability))
- `daps` (List of Object) (see [below for nested schema](#nestedatt--daps))
- `data_ingestion_status` (String) Availability Machine's data ingestion status
- `date_created` (String)
- `date_modified` (String)
- `engine_type` (String) Database Engine Type
- `id` (String) The ID of this resource.
- `logged_in_user_role` (String) The role of the logged in user for accessing the Availability Machine
- `owner` (String) Availability Machine's owner
- `rpo_sla` (List of Object) This is a definition for Tessell Availability Machine's cloud and region availability details (see [below for nested schema](#nestedatt--rpo_sla))
- `service_name` (String)
- `shared_with` (List of Object) Tessell Entity ACL Sharing Info (see [below for nested schema](#nestedatt--shared_with))
- `subscription` (String) Dmm's subscription name
- `tenant` (String) Dmm's tenancy details
- `tessell_service_id` (String)
- `user_id` (String) Data Management Machine's user

<a id="nestedatt--clones"></a>
### Nested Schema for `clones`

Read-Only:

- `clone_info` (Map of String)
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedobjatt--clones--cloud_availability))
- `date_created` (String)
- `id` (String)
- `name` (String)
- `owner` (String)
- `status` (String)
- `subscription` (String)

<a id="nestedobjatt--clones--cloud_availability"></a>
### Nested Schema for `clones.cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--clones--cloud_availability--regions))

<a id="nestedobjatt--clones--cloud_availability--regions"></a>
### Nested Schema for `clones.cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)




<a id="nestedatt--cloud_availability"></a>
### Nested Schema for `cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--cloud_availability--regions))

<a id="nestedobjatt--cloud_availability--regions"></a>
### Nested Schema for `cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)



<a id="nestedatt--daps"></a>
### Nested Schema for `daps`

Read-Only:

- `availability_machine_id` (String)
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedobjatt--daps--cloud_availability))
- `content_info` (List of Object) (see [below for nested schema](#nestedobjatt--daps--content_info))
- `content_type` (String)
- `data_access_config` (List of Object) (see [below for nested schema](#nestedobjatt--daps--data_access_config))
- `date_created` (String)
- `date_modified` (String)
- `engine_type` (String)
- `id` (String)
- `logged_in_user_role` (String)
- `name` (String)
- `owner` (String)
- `service_name` (String)
- `shared_with` (List of Object) (see [below for nested schema](#nestedobjatt--daps--shared_with))
- `status` (String)
- `tessell_service_id` (String)

<a id="nestedobjatt--daps--cloud_availability"></a>
### Nested Schema for `daps.cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--daps--cloud_availability--regions))

<a id="nestedobjatt--daps--cloud_availability--regions"></a>
### Nested Schema for `daps.cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)



<a id="nestedobjatt--daps--content_info"></a>
### Nested Schema for `daps.content_info`

Read-Only:

- `as_is_content` (List of Object) (see [below for nested schema](#nestedobjatt--daps--content_info--as_is_content))
- `sanitized_content` (List of Object) (see [below for nested schema](#nestedobjatt--daps--content_info--sanitized_content))

<a id="nestedobjatt--daps--content_info--as_is_content"></a>
### Nested Schema for `daps.content_info.as_is_content`

Read-Only:

- `automated` (Boolean)
- `manual` (List of Object) (see [below for nested schema](#nestedobjatt--daps--content_info--as_is_content--manual))

<a id="nestedobjatt--daps--content_info--as_is_content--manual"></a>
### Nested Schema for `daps.content_info.as_is_content.manual`

Read-Only:

- `shared_at` (String)
- `snapshot_id` (String)
- `snapshot_name` (String)
- `snapshot_time` (String)



<a id="nestedobjatt--daps--content_info--sanitized_content"></a>
### Nested Schema for `daps.content_info.sanitized_content`

Read-Only:

- `automated` (List of Object) (see [below for nested schema](#nestedobjatt--daps--content_info--sanitized_content--automated))
- `manual` (List of Object) (see [below for nested schema](#nestedobjatt--daps--content_info--sanitized_content--manual))

<a id="nestedobjatt--daps--content_info--sanitized_content--automated"></a>
### Nested Schema for `daps.content_info.sanitized_content.manual`

Read-Only:

- `sanitization_schedule_id` (String)


<a id="nestedobjatt--daps--content_info--sanitized_content--manual"></a>
### Nested Schema for `daps.content_info.sanitized_content.manual`

Read-Only:

- `shared_at` (String)
- `snapshot_id` (String)
- `snapshot_name` (String)
- `snapshot_time` (String)




<a id="nestedobjatt--daps--data_access_config"></a>
### Nested Schema for `daps.data_access_config`

Read-Only:

- `daily_backups` (Number)


<a id="nestedobjatt--daps--shared_with"></a>
### Nested Schema for `daps.shared_with`

Read-Only:

- `users` (List of Object) (see [below for nested schema](#nestedobjatt--daps--shared_with--users))

<a id="nestedobjatt--daps--shared_with--users"></a>
### Nested Schema for `daps.shared_with.users`

Read-Only:

- `email_id` (String)
- `role` (String)




<a id="nestedatt--rpo_sla"></a>
### Nested Schema for `rpo_sla`

Read-Only:

- `availability_machine` (String)
- `availability_machine_id` (String)
- `rpo_sla_status` (String)
- `schedule` (List of Object) (see [below for nested schema](#nestedobjatt--rpo_sla--schedule))
- `sla` (String)
- `topology` (List of Object) (see [below for nested schema](#nestedobjatt--rpo_sla--topology))

<a id="nestedobjatt--rpo_sla--schedule"></a>
### Nested Schema for `rpo_sla.schedule`

Read-Only:

- `backup_start_time` (List of Object) (see [below for nested schema](#nestedobjatt--rpo_sla--schedule--backup_start_time))
- `daily_schedule` (List of Object) (see [below for nested schema](#nestedobjatt--rpo_sla--schedule--daily_schedule))

<a id="nestedobjatt--rpo_sla--schedule--backup_start_time"></a>
### Nested Schema for `rpo_sla.schedule.backup_start_time`

Read-Only:

- `hour` (Number)
- `minute` (Number)


<a id="nestedobjatt--rpo_sla--schedule--daily_schedule"></a>
### Nested Schema for `rpo_sla.schedule.daily_schedule`

Read-Only:

- `backup_start_times` (List of Object) (see [below for nested schema](#nestedobjatt--rpo_sla--schedule--daily_schedule--backup_start_times))
- `backups_per_day` (Number)

<a id="nestedobjatt--rpo_sla--schedule--daily_schedule--backup_start_times"></a>
### Nested Schema for `rpo_sla.schedule.daily_schedule.backups_per_day`

Read-Only:

- `hour` (Number)
- `minute` (Number)




<a id="nestedobjatt--rpo_sla--topology"></a>
### Nested Schema for `rpo_sla.topology`

Read-Only:

- `availability_zones` (List of String)
- `cloud_type` (String)
- `region` (String)
- `type` (String)



<a id="nestedatt--shared_with"></a>
### Nested Schema for `shared_with`

Read-Only:

- `users` (List of Object) (see [below for nested schema](#nestedobjatt--shared_with--users))

<a id="nestedobjatt--shared_with--users"></a>
### Nested Schema for `shared_with.users`

Read-Only:

- `email_id` (String)
- `role` (String)

