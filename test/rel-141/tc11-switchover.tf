# TC-11: Switchover — promote replica to primary
# Prerequisites: TC-09 applied, both instances in READY state
#
# Change from TC-09:
#   - default-node-0: role "primary" -> "replica"
#   - replica-node-1: role "replica" -> "primary"
#
# Apply with: terraform apply -auto-approve
# Verify: terraform show — check which instance is now primary

resource "tessell_db_service" "tc11_switchover" {
  name                       = "tc01-rel141-oracle"
  description                = "TC-11: switchover"
  subscription               = "as-pre-01"
  edition                    = "ENTERPRISE"
  engine_type                = "ORACLE"
  topology                   = "single_instance"
  software_image             = "Oracle 21c"
  software_image_version     = "21.0.0.0.0"
  auto_minor_version_update  = true
  enable_deletion_protection = false
  enable_stop_protection     = false

  infrastructure {
    cloud                  = "azure"
    region                 = "northCentralUS"
    vpc                    = "tessell-virtual-network-82aoz"
    private_subnet         = "tessell-private-subnet"
    compute_type           = "tesl_4_b-2"
    enable_encryption      = false
    additional_storage     = 53687091200
    iops                   = 3000
    throughput             = 125
    timezone               = "America/Chicago"
    compute_name_prefix    = "server1"
    enable_compute_sharing = true
  }

  service_connectivity {
    service_port         = "1521"
    enable_public_access = false
    enable_ssl           = false
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
    enable_auto_snapshot     = true
    enable_auto_backup       = false
    include_transaction_logs = true
    standard_policy {
      retention_days           = 7
      include_transaction_logs = true
      snapshot_start_time {
        hour   = 1
        minute = 0
      }
    }
  }

  maintenance_window {
    cadence  = "WEEKLY"
    time     = "02:00"
    duration = 30
    day      = "Sunday"
  }

  auto_patch_config {
    os_auto_patch_enabled = true
    db_auto_patch_enabled = true
    patch_strategy        = "LATEST_CERTIFIED"
  }

  server_patching_config {
    enable_auto_os_patching = true
  }

  # TC-11: role swapped — replica-node-1 is now primary
  instances {
    name                = "default-node-0"
    instance_group_name = "default"
    role                = "replica"   # was "primary" in TC-09
    region              = "northCentralUS"
    compute_id          = "25ccc4b2-7a0a-49ed-a9bd-bc224b3f593e"
    compute_type        = "tesl_4_b-2"
    vpc                 = "tessell-virtual-network-82aoz"
    enable_perf_insights = false
    storage_config {
      provider = "AZURE_MANAGED_DISK"
    }
    security_config {
      security_profile_id = null
    }
  }

  instances {
    name                = "replica-node-1"
    instance_group_name = "replica-group"
    role                = "primary"   # was "replica" in TC-09
    region              = "northCentralUS"
    compute_type        = "tesl_4_b-2"
    vpc                 = "tessell-virtual-network-82aoz"
    enable_perf_insights = false
    storage_config {
      provider = "AZURE_MANAGED_DISK"
    }
    security_config {
      security_profile_id = null
    }
  }
}

output "tc11_primary_instance" {
  description = "After switchover replica-node-1 should be primary"
  value = [
    for i in tessell_db_service.tc11_switchover.instances :
    i.name if i.role == "primary"
  ]
}
