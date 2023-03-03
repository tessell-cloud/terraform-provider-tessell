---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tessell_db_services Data Source - terraform-provider-tessell"
subcategory: ""
description: |-
  
---

# tessell_db_services (Data Source)



## Example Usage

```terraform
# List all tessell_db_service
data_source "tessell_db_services" "example" {}

# Get a tessell_db_service using the service name
data_source "tessell_db_services" "example" {
  name = "my-test-db"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cloned_from_availability_machine_id` (String) The id of the Availability Machine from which the services are cloned
- `cloned_from_service_id` (String) The id of the DB Service from which the services are cloned
- `engine_types` (List of String) DB Service's engine-types
- `load_acls` (Boolean) Load ACL information
- `load_databases` (Boolean) Load the databases that are part of the DB Service
- `load_instances` (Boolean) Load the instances that are part of the DB Service
- `name` (String) Name of the DB Service
- `owners` (List of String) List of Email Addresses for entity or resource owners
- `statuses` (List of String) statuses

### Read-Only

- `db_services` (List of Object) (see [below for nested schema](#nestedatt--db_services))
- `id` (String) The ID of this resource.

<a id="nestedatt--db_services"></a>
### Nested Schema for `db_services`

Read-Only:

- `auto_minor_version_update` (Boolean)
- `availability_machine_id` (String)
- `cloned_from_info` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--cloned_from_info))
- `databases` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases))
- `date_created` (String)
- `deletion_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--deletion_config))
- `deletion_schedule` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--deletion_schedule))
- `description` (String)
- `enable_deletion_protection` (Boolean)
- `engine_configuration` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration))
- `engine_type` (String)
- `id` (String)
- `infrastructure` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--infrastructure))
- `instances` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--instances))
- `integrations_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--integrations_config))
- `license_type` (String)
- `logged_in_user_role` (String)
- `maintenance_window` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--maintenance_window))
- `name` (String)
- `num_of_instances` (Number)
- `owner` (String)
- `service_connectivity` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--service_connectivity))
- `shared_with` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--shared_with))
- `software_image` (String)
- `software_image_version` (String)
- `started_at` (String)
- `status` (String)
- `stopped_at` (String)
- `subscription` (String)
- `tags` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--tags))
- `tenant_id` (String)
- `tessell_genie_status` (String)
- `topology` (String)
- `upcoming_scheduled_actions` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--upcoming_scheduled_actions))
- `updates_in_progress` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--updates_in_progress))

<a id="nestedobjatt--db_services--cloned_from_info"></a>
### Nested Schema for `db_services.cloned_from_info`

Read-Only:

- `availability_machine` (String)
- `availability_machine_id` (String)
- `maximum_recoverability` (Boolean)
- `pitr_time` (String)
- `snapshot_id` (String)
- `snapshot_name` (String)
- `tessell_service` (String)
- `tessell_service_id` (String)


<a id="nestedobjatt--db_services--databases"></a>
### Nested Schema for `db_services.databases`

Read-Only:

- `cloned_from_info` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases--cloned_from_info))
- `database_configuration` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases--database_configuration))
- `database_name` (String)
- `date_created` (String)
- `description` (String)
- `engine_type` (String)
- `id` (String)
- `source_database_id` (String)
- `status` (String)
- `tessell_service_id` (String)

<a id="nestedobjatt--db_services--databases--cloned_from_info"></a>
### Nested Schema for `db_services.databases.cloned_from_info`

Read-Only:

- `database_id` (String)


<a id="nestedobjatt--db_services--databases--database_configuration"></a>
### Nested Schema for `db_services.databases.database_configuration`

Read-Only:

- `mysql_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases--database_configuration--mysql_config))
- `oracle_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases--database_configuration--oracle_config))
- `postgresql_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases--database_configuration--postgresql_config))
- `sql_server_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--databases--database_configuration--sql_server_config))

<a id="nestedobjatt--db_services--databases--database_configuration--mysql_config"></a>
### Nested Schema for `db_services.databases.database_configuration.sql_server_config`

Read-Only:

- `parameter_profile` (String)


<a id="nestedobjatt--db_services--databases--database_configuration--oracle_config"></a>
### Nested Schema for `db_services.databases.database_configuration.sql_server_config`

Read-Only:

- `options_profile` (String)
- `parameter_profile` (String)


<a id="nestedobjatt--db_services--databases--database_configuration--postgresql_config"></a>
### Nested Schema for `db_services.databases.database_configuration.sql_server_config`

Read-Only:

- `parameter_profile` (String)


<a id="nestedobjatt--db_services--databases--database_configuration--sql_server_config"></a>
### Nested Schema for `db_services.databases.database_configuration.sql_server_config`

Read-Only:

- `parameter_profile` (String)




<a id="nestedobjatt--db_services--deletion_config"></a>
### Nested Schema for `db_services.deletion_config`

Read-Only:

- `retain_availability_machine` (Boolean)


<a id="nestedobjatt--db_services--deletion_schedule"></a>
### Nested Schema for `db_services.deletion_schedule`

Read-Only:

- `delete_at` (String)
- `deletion_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--deletion_schedule--deletion_config))

<a id="nestedobjatt--db_services--deletion_schedule--deletion_config"></a>
### Nested Schema for `db_services.deletion_schedule.deletion_config`

Read-Only:

- `retain_availability_machine` (Boolean)



<a id="nestedobjatt--db_services--engine_configuration"></a>
### Nested Schema for `db_services.engine_configuration`

Read-Only:

- `apache_kafka_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--apache_kafka_config))
- `mysql_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--mysql_config))
- `oracle_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--oracle_config))
- `post_script_info` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--post_script_info))
- `postgresql_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--postgresql_config))
- `pre_script_info` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--pre_script_info))
- `sql_server_config` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--engine_configuration--sql_server_config))

<a id="nestedobjatt--db_services--engine_configuration--apache_kafka_config"></a>
### Nested Schema for `db_services.engine_configuration.apache_kafka_config`

Read-Only:

- `parameter_profile` (String)


<a id="nestedobjatt--db_services--engine_configuration--mysql_config"></a>
### Nested Schema for `db_services.engine_configuration.mysql_config`

Read-Only:

- `parameter_profile` (String)


<a id="nestedobjatt--db_services--engine_configuration--oracle_config"></a>
### Nested Schema for `db_services.engine_configuration.oracle_config`

Read-Only:

- `character_set` (String)
- `multi_tenant` (Boolean)
- `national_character_set` (String)
- `options_profile` (String)
- `parameter_profile` (String)


<a id="nestedobjatt--db_services--engine_configuration--post_script_info"></a>
### Nested Schema for `db_services.engine_configuration.post_script_info`

Read-Only:

- `script_id` (String)
- `script_version` (String)


<a id="nestedobjatt--db_services--engine_configuration--postgresql_config"></a>
### Nested Schema for `db_services.engine_configuration.postgresql_config`

Read-Only:

- `parameter_profile` (String)


<a id="nestedobjatt--db_services--engine_configuration--pre_script_info"></a>
### Nested Schema for `db_services.engine_configuration.pre_script_info`

Read-Only:

- `script_id` (String)
- `script_version` (String)


<a id="nestedobjatt--db_services--engine_configuration--sql_server_config"></a>
### Nested Schema for `db_services.engine_configuration.sql_server_config`

Read-Only:

- `parameter_profile` (String)



<a id="nestedobjatt--db_services--infrastructure"></a>
### Nested Schema for `db_services.infrastructure`

Read-Only:

- `additional_storage` (Number)
- `availability_zone` (String)
- `cloud` (String)
- `cloud_availability` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--infrastructure--cloud_availability))
- `compute_type` (String)
- `enable_encryption` (Boolean)
- `encryption_key` (String)
- `region` (String)
- `storage` (Number)
- `vpc` (String)

<a id="nestedobjatt--db_services--infrastructure--cloud_availability"></a>
### Nested Schema for `db_services.infrastructure.cloud_availability`

Read-Only:

- `cloud` (String)
- `regions` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--infrastructure--cloud_availability--regions))

<a id="nestedobjatt--db_services--infrastructure--cloud_availability--regions"></a>
### Nested Schema for `db_services.infrastructure.cloud_availability.regions`

Read-Only:

- `availability_zones` (List of String)
- `region` (String)




<a id="nestedobjatt--db_services--instances"></a>
### Nested Schema for `db_services.instances`

Read-Only:

- `availability_zone` (String)
- `cloud` (String)
- `compute_type` (String)
- `connect_string` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--instances--connect_string))
- `date_created` (String)
- `encryption_key` (String)
- `id` (String)
- `instance_group_id` (String)
- `last_started_at` (String)
- `last_stopped_at` (String)
- `name` (String)
- `region` (String)
- `role` (String)
- `software_image` (String)
- `software_image_version` (String)
- `status` (String)
- `tessell_service_id` (String)
- `type` (String)
- `updates_in_progress` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--instances--updates_in_progress))
- `vpc` (String)

<a id="nestedobjatt--db_services--instances--connect_string"></a>
### Nested Schema for `db_services.instances.connect_string`

Read-Only:

- `connect_descriptor` (String)
- `endpoint` (String)
- `master_user` (String)
- `service_port` (String)


<a id="nestedobjatt--db_services--instances--updates_in_progress"></a>
### Nested Schema for `db_services.instances.updates_in_progress`

Read-Only:

- `reference_id` (String)
- `submitted_at` (String)
- `update_info` (Map of String)
- `update_type` (String)



<a id="nestedobjatt--db_services--integrations_config"></a>
### Nested Schema for `db_services.integrations_config`

Read-Only:

- `integrations` (List of String)


<a id="nestedobjatt--db_services--maintenance_window"></a>
### Nested Schema for `db_services.maintenance_window`

Read-Only:

- `day` (String)
- `duration` (Number)
- `time` (String)


<a id="nestedobjatt--db_services--service_connectivity"></a>
### Nested Schema for `db_services.service_connectivity`

Read-Only:

- `allowed_ip_addresses` (List of String)
- `connect_strings` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--service_connectivity--connect_strings))
- `dns_prefix` (String)
- `enable_public_access` (Boolean)
- `private_link` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--service_connectivity--private_link))
- `service_port` (Number)
- `update_in_progress_info` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--service_connectivity--update_in_progress_info))

<a id="nestedobjatt--db_services--service_connectivity--connect_strings"></a>
### Nested Schema for `db_services.service_connectivity.connect_strings`

Read-Only:

- `connect_descriptor` (String)
- `endpoint` (String)
- `master_user` (String)
- `service_port` (Number)
- `type` (String)
- `usage_type` (String)


<a id="nestedobjatt--db_services--service_connectivity--private_link"></a>
### Nested Schema for `db_services.service_connectivity.private_link`

Read-Only:

- `endpoint_service_name` (String)
- `service_principals` (List of String)


<a id="nestedobjatt--db_services--service_connectivity--update_in_progress_info"></a>
### Nested Schema for `db_services.service_connectivity.update_in_progress_info`

Read-Only:

- `allowed_ip_addresses` (List of String)
- `dns_prefix` (String)
- `enable_public_access` (Boolean)
- `private_link` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--service_connectivity--update_in_progress_info--private_link))

<a id="nestedobjatt--db_services--service_connectivity--update_in_progress_info--private_link"></a>
### Nested Schema for `db_services.service_connectivity.update_in_progress_info.private_link`

Read-Only:

- `service_principals` (List of String)




<a id="nestedobjatt--db_services--shared_with"></a>
### Nested Schema for `db_services.shared_with`

Read-Only:

- `users` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--shared_with--users))

<a id="nestedobjatt--db_services--shared_with--users"></a>
### Nested Schema for `db_services.shared_with.users`

Read-Only:

- `email_id` (String)
- `role` (String)



<a id="nestedobjatt--db_services--tags"></a>
### Nested Schema for `db_services.tags`

Read-Only:

- `name` (String)
- `value` (String)


<a id="nestedobjatt--db_services--upcoming_scheduled_actions"></a>
### Nested Schema for `db_services.upcoming_scheduled_actions`

Read-Only:

- `delete` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--upcoming_scheduled_actions--delete))
- `start_stop` (List of Object) (see [below for nested schema](#nestedobjatt--db_services--upcoming_scheduled_actions--start_stop))

<a id="nestedobjatt--db_services--upcoming_scheduled_actions--delete"></a>
### Nested Schema for `db_services.upcoming_scheduled_actions.delete`

Read-Only:

- `at` (String)


<a id="nestedobjatt--db_services--upcoming_scheduled_actions--start_stop"></a>
### Nested Schema for `db_services.upcoming_scheduled_actions.start_stop`

Read-Only:

- `action` (String)
- `at` (String)



<a id="nestedobjatt--db_services--updates_in_progress"></a>
### Nested Schema for `db_services.updates_in_progress`

Read-Only:

- `reference_id` (String)
- `submitted_at` (String)
- `update_info` (Map of String)
- `update_type` (String)

