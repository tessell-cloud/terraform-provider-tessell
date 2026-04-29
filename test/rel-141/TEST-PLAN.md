# Rel-141 Terraform Provider — Test Plan

> **Branch**: `autogen-0.0.32-rel-141`
> **Environment**: `https://api.terraformrelease140001.tsl-terls.cloud`
> **Tenant ID**: `fbab7165-97a3-42c7-8be5-abcf31cb62f7`
> **API Key**: set via `export TF_VAR_api_key="<key>"` — never hardcode

---

## How to run

```bash
cd test/rel-141
export TF_VAR_api_key="<key>"
terraform init
terraform plan          # dry-run
terraform apply -auto-approve
terraform destroy -auto-approve
```

---

## Rel-141 Changes Implemented

| # | Change | File(s) |
|---|--------|---------|
| 1 | `maintenance_window.cadence` (WEEKLY/MONTHLY/QUARTERLY) | `resource_db_service.go` |
| 2 | `maintenance_window.day_of_month` (MONTHLY) | `resource_db_service.go` |
| 3 | `maintenance_window.start_date` (QUARTERLY) | `resource_db_service.go` |
| 4 | Plan-time validation: cadence ↔ required sub-field | `resource_db_service.go` |
| 5 | `auto_patch_config` block (create-only, error on update) | `resource_db_service.go` |
| 6 | `server_patching_config` block (create-only, error on update) | `resource_db_service.go` |
| 7 | `is_hpc` Computed field at service level & `cloned_from_info` | `resource_db_service.go`, `helpers.go` |
| 8 | `gcp_infra_config` in instances schema | `resource_db_service.go` |
| 9 | `updates_info.upcoming_maintenance_window` (Computed) | `resource_db_service.go` |
| 10 | `is_hpc` in data sources (db_service, db_services) | `data_source_db_service.go`, `data_source_db_services.go` |
| 11 | `is_hpc` in availability_machine data sources | `data_source_availability_machine*.go`, `helpers.go` |
| 12 | Restored `docs/resources/db_snapshot.md` | `docs/resources/db_snapshot.md` |

---

## Test Cases

### Legend
- ⬜ NOT STARTED
- 🔄 IN PROGRESS
- ✅ PASSED
- ❌ FAILED
- ⏭️ SKIPPED

---

### Group A — Core Lifecycle (Must Pass)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-01 | Create DB service (baseline) | `tc01-create-service.tf` | ⬜ | ORACLE, Azure, single_instance |
| TC-02 | Read/refresh — `terraform plan` shows no diff after create | `tc01-create-service.tf` | ⬜ | Run plan immediately after apply; expect 0 changes |
| TC-09 | Add replica instance | `tc09-add-instance.tf` | ⬜ | Adds replica-node-1 |
| TC-10 | Delete replica instance | (modify tc09, remove replica block) | ⬜ | Remove replica, apply |
| TC-11 | Switchover (promote replica to primary) | `tc11-switchover.tf` | ⬜ | Swap roles in instances block |
| TC-12 | Destroy service | `terraform destroy` | ⬜ | Clean teardown after TC-11 |

---

### Group B — Rel-141 New Fields: Patch Config (Create-Only)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-03 | Create with `auto_patch_config` | `tc01-create-service.tf` | ⬜ | Included in baseline TC-01 config |
| TC-04 | Create with `server_patching_config` | `tc01-create-service.tf` | ⬜ | Included in baseline TC-01 config |
| TC-13 | Attempt update of `auto_patch_config` | manual edit after TC-01 | ⬜ | Expect `diag.Errorf` — plan must error |
| TC-14 | Attempt update of `server_patching_config` | manual edit after TC-01 | ⬜ | Expect `diag.Errorf` — plan must error |

---

### Group C — Rel-141 New Fields: Maintenance Window Cadence

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-05 | `additional_storage` field round-trips | `tc01-create-service.tf` | ⬜ | Value 53687091200 must persist in state |
| TC-06 | MW cadence=WEEKLY with `day` — valid | `tc06-mw-weekly-valid.tf` | ⬜ | Plan succeeds |
| TC-06b | MW cadence=WEEKLY without `day` — invalid | `tc06-mw-weekly-valid.tf` (comment swap) | ⬜ | Plan-time error expected |
| TC-07 | MW cadence=MONTHLY with `day_of_month` — valid | `tc06-mw-weekly-valid.tf` (comment swap) | ⬜ | Plan succeeds |
| TC-07b | MW cadence=MONTHLY without `day_of_month` — invalid | `tc06-mw-weekly-valid.tf` (comment swap) | ⬜ | Plan-time error expected |
| TC-08 | MW cadence=QUARTERLY with `start_date` — valid | `tc06-mw-weekly-valid.tf` (comment swap) | ⬜ | Plan succeeds |
| TC-08b | MW cadence=QUARTERLY without `start_date` — invalid | `tc06-mw-weekly-valid.tf` (comment swap) | ⬜ | Plan-time error expected |
| TC-15 | Update maintenance window (WEEKLY→MONTHLY) | modify tc01 state | ⬜ | Update should succeed via API |

---

### Group D — Rel-141 New Fields: Computed / Read-Only

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-16 | `is_hpc` populated in state after create | `tc01-create-service.tf` (output) | ⬜ | Check `tc01_is_hpc` output |
| TC-17 | `is_hpc` present in data source read | manual data source query | ⬜ | Use `data "tessell_db_service"` |
| TC-18 | `updates_info` populated (read-only block) | `tc01-create-service.tf` (output) | ⬜ | Check `tc01_updates_info` output |

---

### Group E — Import & Refresh

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-19 | `terraform import` existing service | `tc01-create-service.tf` | ⬜ | `terraform import tessell_db_service.tc01_create <id>` |
| TC-20 | Refresh — no spurious diff after import | after TC-19 | ⬜ | `terraform plan` must show 0 changes |

---

### Group F — Start/Stop & Delete Schedule (Rel-141 helper changes)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-21 | Create start/stop schedule | separate `tessell_db_service_start_stop_schedule` resource | ⬜ | Confirm new fields in helpers |
| TC-22 | Update start/stop schedule | modify schedule | ⬜ |  |
| TC-23 | Create delete schedule | `tessell_db_service_delete_schedule` resource | ⬜ | Confirm new fields in helpers |

---

### Group G — Snapshot (Regression)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-24 | Create snapshot via `tessell_db_snapshot` | separate snapshot resource | ⬜ | Confirm doc restored, API works |
| TC-25 | Read snapshot via data source | `data "tessell_db_snapshot"` | ⬜ |  |

---

### Group H — Availability Machine Data Source

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-26 | `is_hpc` present in AM data source | `data "tessell_availability_machine"` | ⬜ | Verify rel-141 AM changes |
| TC-27 | `is_hpc` present in AMs list data source | `data "tessell_availability_machines"` | ⬜ |  |

---

## Execution Log

| Date | TC | Result | Tester | Remarks |
|------|----|--------|--------|---------|
|      |    |        |        |         |
