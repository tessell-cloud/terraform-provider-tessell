# TC-06 / TC-07 / TC-08: Maintenance Window cadence validation
# Tests that plan-time validation fires correctly for each cadence type.
#
# VALID config for WEEKLY — plan should succeed:
#   cadence = "WEEKLY", day = "Sunday"
#
# INVALID configs to test manually (uncomment one at a time and run `terraform plan`):
#   A) cadence = "WEEKLY" but no `day`       → error: 'day' is required
#   B) cadence = "MONTHLY" but no day_of_month → error: 'day_of_month' is required
#   C) cadence = "QUARTERLY" but no start_date → error: 'start_date' is required

locals {
  # Switch this to test each invalid cadence scenario
  # Options: "WEEKLY_VALID", "WEEKLY_MISSING_DAY", "MONTHLY_VALID",
  #          "MONTHLY_MISSING_DOM", "QUARTERLY_VALID", "QUARTERLY_MISSING_SD"
  scenario = "WEEKLY_VALID"
}

resource "tessell_db_service" "tc06_mw_cadence" {
  name                       = "tc06-mw-cadence-test"
  description                = "TC-06/07/08: maintenance_window cadence validation"
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

  # ── TC-06 VALID: WEEKLY with day ─────────────────────────────────────────
  maintenance_window {
    cadence  = "WEEKLY"
    time     = "02:00"
    duration = 30
    day      = "Sunday"
  }

  # ── Uncomment to test TC-06 INVALID: WEEKLY without day ──────────────────
  # maintenance_window {
  #   cadence  = "WEEKLY"
  #   time     = "02:00"
  #   duration = 30
  #   # day omitted → plan-time error expected
  # }

  # ── Uncomment to test TC-07 VALID: MONTHLY ───────────────────────────────
  # maintenance_window {
  #   cadence      = "MONTHLY"
  #   time         = "02:00"
  #   duration     = 30
  #   day_of_month = 15
  # }

  # ── Uncomment to test TC-07 INVALID: MONTHLY without day_of_month ────────
  # maintenance_window {
  #   cadence  = "MONTHLY"
  #   time     = "02:00"
  #   duration = 30
  #   # day_of_month omitted → plan-time error expected
  # }

  # ── Uncomment to test TC-08 VALID: QUARTERLY ─────────────────────────────
  # maintenance_window {
  #   cadence    = "QUARTERLY"
  #   time       = "02:00"
  #   duration   = 30
  #   start_date = "2026-07-01"
  # }

  # ── Uncomment to test TC-08 INVALID: QUARTERLY without start_date ────────
  # maintenance_window {
  #   cadence  = "QUARTERLY"
  #   time     = "02:00"
  #   duration = 30
  #   # start_date omitted → plan-time error expected
  # }

  instances {
    name                = "default-node-0"
    instance_group_name = "default"
    role                = "primary"
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
}
