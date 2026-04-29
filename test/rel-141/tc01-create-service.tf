# TC-01 + TC-03 + TC-04 + TC-05 + TC-06
# Create DB service with all rel-141 fields:
#   - maintenance_window.cadence (WEEKLY) [rel-141 new field]
#   - auto_patch_config                   [rel-141 new block]
#   - server_patching_config              [rel-141 new block]
#   - infrastructure.additional_storage   [rel-141, was already in schema]
#
# Run:
#   export TF_VAR_api_key="KD5gqopVYvu7m7oNGgVB6JLpAWB3wpKurzA6VQ"
#   terraform init && terraform apply -auto-approve

resource "tessell_db_service" "tc01_create" {
  name                      = "tc01-rel141-oracle"
  description               = "TC-01: baseline create with all rel-141 fields"
  subscription              = "as-pre-01"
  edition                   = "ENTERPRISE"
  engine_type               = "ORACLE"
  topology                  = "single_instance"
  software_image            = "Oracle 21c"
  software_image_version    = "21.0.0.0.0"
  auto_minor_version_update = true
  enable_deletion_protection = false
  enable_stop_protection    = false

  infrastructure {
    cloud                 = "azure"
    region                = "northCentralUS"
    vpc                   = "tessell-virtual-network-82aoz"
    private_subnet        = "tessell-private-subnet"
    compute_type          = "tesl_4_b-2"
    enable_encryption     = false
    additional_storage    = 53687091200 # TC-05: additional_storage field
    iops                  = 3000
    throughput            = 125
    timezone              = "America/Chicago"
    compute_name_prefix   = "server1"
    enable_compute_sharing = true
  }

  service_connectivity {
    service_port        = "1521"
    enable_public_access = false
    enable_ssl          = false
  }

  creds {
    master_user     = "master"
    master_password = "Tessell123ZX#"
  }

  engine_configuration {
    oracle_config {
      multi_tenant           = true
      parameter_profile_id   = "5390c715-fb24-464b-9167-b3ee6fe988dd"
      option_profile_id      = "b52ce979-fbbd-46eb-af40-e670910ac51e"
      character_set          = "AL32UTF8"
      national_character_set = "AL16UTF16"
      enable_archive_mode    = true
    }
  }

  databases {
    database_name = "orcle98"
    database_configuration {
      oracle_config {
        parameter_profile_id = "5390c715-fb24-464b-9167-b3ee6fe988dd"
        option_profile_id    = "b52ce979-fbbd-46eb-af40-e670910ac51e"
      }
    }
  }

  rpo_policy_config {
    enable_auto_snapshot    = true
    enable_auto_backup      = false
    include_transaction_logs = true
    standard_policy {
      retention_days          = 7
      include_transaction_logs = true
      snapshot_start_time {
        hour   = 1
        minute = 0
      }
    }
  }

  # TC-06: WEEKLY cadence — rel-141 new required field
  maintenance_window {
    cadence  = "WEEKLY"
    time     = "02:00"
    duration = 30
    day      = "Sunday"
  }

  # TC-03: auto_patch_config — rel-141 new block, create-only
  auto_patch_config {
    os_auto_patch_enabled = true
    db_auto_patch_enabled = true
    patch_strategy        = "LATEST_CERTIFIED"
  }

  # TC-04: server_patching_config — rel-141 new block, create-only
  server_patching_config {
    enable_auto_os_patching = true
  }

  instances {
    name               = "default-node-0"
    instance_group_name = "default"
    role               = "primary"
    region             = "northCentralUS"
    compute_id         = "25ccc4b2-7a0a-49ed-a9bd-bc224b3f593e"
    compute_type       = "tesl_4_b-2"
    vpc                = "tessell-virtual-network-82aoz"
    enable_perf_insights = false
    storage_config {
      provider = "AZURE_MANAGED_DISK"
    }
    security_config {
      security_profile_id = null
    }
  }
}

output "tc01_service_id" {
  value = tessell_db_service.tc01_create.id
}

output "tc01_is_hpc" {
  description = "TC-15: is_hpc should be populated as Computed field"
  value       = tessell_db_service.tc01_create.is_hpc
}

output "tc01_updates_info" {
  description = "TC-18: updates_info should be populated (read-only)"
  value       = tessell_db_service.tc01_create.updates_info
}
