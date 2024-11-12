---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tessell_db_service_start_stop_schedule Data Source - terraform-provider-tessell"
subcategory: ""
description: |-
  
---

# tessell_db_service_start_stop_schedule (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The ID of the schedule

### Read-Only

- `date_created` (String) Timestamp when the schedule was created
- `date_modified` (String) Timestamp when the schedule was last modified
- `description` (String)
- `last_run` (String) The date-time at which this schedule was last executed
- `metadata` (List of Object) (see [below for nested schema](#nestedatt--metadata))
- `name` (String) Name of the schedule
- `schedule_info` (List of Object) Describes the start/stop schedule of tessell service (see [below for nested schema](#nestedatt--schedule_info))
- `service_id` (String) The ID of the DB Service
- `status` (String) StartStopScheduleStatus

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Read-Only:

- `schedule_counter` (Number)


<a id="nestedatt--schedule_info"></a>
### Nested Schema for `schedule_info`

Read-Only:

- `one_time` (List of Object) (see [below for nested schema](#nestedobjatt--schedule_info--one_time))
- `recurring` (List of Object) (see [below for nested schema](#nestedobjatt--schedule_info--recurring))

<a id="nestedobjatt--schedule_info--one_time"></a>
### Nested Schema for `schedule_info.one_time`

Read-Only:

- `db_service_start_at` (String)
- `db_service_stop_at` (String)


<a id="nestedobjatt--schedule_info--recurring"></a>
### Nested Schema for `schedule_info.recurring`

Read-Only:

- `daily_schedule` (Boolean)
- `db_service_start_at` (String)
- `db_service_stop_at` (String)
- `schedule_expiry` (List of Object) (see [below for nested schema](#nestedobjatt--schedule_info--recurring--schedule_expiry))
- `schedule_start_date` (String)
- `weekly_schedule` (List of Object) (see [below for nested schema](#nestedobjatt--schedule_info--recurring--weekly_schedule))

<a id="nestedobjatt--schedule_info--recurring--schedule_expiry"></a>
### Nested Schema for `schedule_info.recurring.schedule_expiry`

Read-Only:

- `after_occurrences` (Number)
- `never` (Boolean)
- `on` (String)


<a id="nestedobjatt--schedule_info--recurring--weekly_schedule"></a>
### Nested Schema for `schedule_info.recurring.weekly_schedule`

Read-Only:

- `days` (List of String)

