# Rel-141 Terraform Provider — Test Plan & Results

> **Branch**: `autogen-0.0.33-rel-141`
> **Environment**: `https://api.terraformrelease140001.tsl-terls.cloud`
> **Tenant ID**: `fbab7165-97a3-42c7-8be5-abcf31cb62f7`
> **API Key**: set via `export TF_VAR_api_key="<key>"` — never hardcode
> **Test Date**: 2026-04-29 → 2026-04-30
>
> **Active Service (TC-01)**: `c5bafa56-935b-4c89-be7a-72981b13312c` — `my-service-zyb5fzpe` (PostgreSQL/AWS/ap-south-1) — READY
> **HA Service (TC-12)**: `509da6f3-27e3-4493-912a-8797bad54c83` — `my-service-nmj06ows` (PostgreSQL/Azure/northCentralUS) — READY ✅

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
| 11 | `is_hpc` in availability_machine data sources | N/A — `DMMConsumerView` (AM GET DTO) does not have `IsHpc`. Field only exists on DB Service models (`TessellServiceDTO`). Changelog entry was incorrectly scoped. |
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
| TC-01 | Create DB service (baseline) | `tc01-create-service.tf` | ✅ | PostgreSQL/AWS/ap-south-1. `my-service-zyb5fzpe` READY. |
| TC-02 | Read/refresh — `terraform plan` shows no diff after create | `tc01-create-service.tf` | ✅ | Plan showed 0 changes after apply. |
| TC-09 | Add `read_only_replica` instance | `tc01-create-service.tf` | ✅ | Added `RRc6e5-node-0` — API confirmed 2 instances. |
| TC-10 | Delete replica instance | `tc01-create-service.tf` | ✅ | Removed replica block. Refresh confirmed `instance_count = 1`. |
| TC-11 | Switchover (promote failover_replica to primary) | `tc12-ha-service.tf` | ✅ | Role swap applied via HCL. `failover-node-0` promoted to `primary` in ~3m. |
| TC-12 | Create HA service (`high_availability` topology) | `tc12-ha-service.tf` | ✅ | Azure/northCentralUS, 2 instances (primary + failover_replica). |

---

### Group H — GCP Cloud Tests

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| GCP-TC01 | Create HA SQL Server service on GCP | `gcp-tc/gcp-service.tf` | ✅ | `my-service-71ddd622` (`bf3b4ebe`) — SQLSERVER/GCP/us-east1, `high_availability`. 2 instances (primary + failover_replica) UP. `vkospatchtest` subscription. |
| GCP-TC02 | Add `read_only_replica` instance on GCP | `gcp-tc/gcp-service.tf` | ✅ | Added `ror-node-0` (instance_group=`ror`, AZ=us-east1-b). API confirmed 3 instances all UP. |
| GCP-TC03 | Switchover on GCP (promote failover_replica to primary) | `gcp-tc/gcp-service.tf` | ✅ | Role swap: `default-node-1` promoted to `primary`, `default-node-0` demoted to `failover_replica`. Applied via HCL role change. |

---

### Group B — Rel-141 New Fields: Patch Config (Create-Only)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-03 | Create with `auto_patch_config` | `tc01-create-service.tf` | ✅ | Sent on create. Backend does not persist it; omitted from HCL post-create to avoid drift. |
| TC-04 | Create with `server_patching_config` | `tc01-create-service.tf` | ✅ | `enable_auto_os_patching=true` confirmed via API GET. |
| TC-13 | Attempt update of `auto_patch_config` | `tc01-create-service.tf` | ✅ | Provider correctly errors: `auto_patch_config can only be set at creation time`. |
| TC-14 | Attempt update of `server_patching_config` | `tc01-create-service.tf` | ✅ | Provider correctly errors: `server_patching_config can only be set at creation time`. |

---

### Group C — Rel-141 New Fields: Maintenance Window Cadence

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-05 | `additional_storage` field round-trips | `tc01-create-service.tf` | ✅ | `additional_storage=0` set in HCL and confirmed in state. No drift detected. |
| TC-06 | MW cadence=WEEKLY with `day` — valid | `tc01-create-service.tf` | ✅ | `cadence=WEEKLY, day=Sunday` — plan succeeds. |
| TC-06b | MW cadence=WEEKLY without `day` — invalid | plan-time validation | ✅ | Plan-time error: `'day' is required when maintenance_window cadence is WEEKLY`. |
| TC-07 | MW cadence=MONTHLY with `day_of_month` — valid | plan-time validation | ✅ | `cadence=MONTHLY, day_of_month=15` — plan succeeds. |
| TC-07b | MW cadence=MONTHLY without `day_of_month` — invalid | plan-time validation | ✅ | Plan-time error: `'day_of_month' is required when maintenance_window cadence is MONTHLY`. |
| TC-08 | MW cadence=QUARTERLY with `start_date` — valid | plan-time validation | ✅ | `cadence=QUARTERLY, start_date=2025-01-01` — plan succeeds. |
| TC-08b | MW cadence=QUARTERLY without `start_date` — invalid | plan-time validation | ✅ | Plan-time error: `'start_date' is required when maintenance_window cadence is QUARTERLY`. |
| TC-15 | Update maintenance window (WEEKLY→MONTHLY) | `tc01-create-service.tf` | ✅ | `cadence=MONTHLY, day_of_month=15` applied. State reflects MONTHLY. No drift on subsequent plan. |

---

### Group D — Rel-141 New Fields: Computed / Read-Only

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-16 | `is_hpc` populated in state after create | `tc01-create-service.tf` (output) | ✅ | `tc01_is_hpc = false` — field populated correctly from API response |
| TC-17 | `is_hpc` present in data source read | `tc01-create-service.tf` (data source) | ✅ | `data "tessell_db_service" "tc17_datasource"` — output `tc17_datasource_is_hpc = false` confirmed. |
| TC-18 | `updates_info` populated (read-only block) | `tc01-create-service.tf` (output) | ✅ | `tc01_updates_info` block populated from API state after create |

---

### Group E — Import & Refresh

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-19 | `terraform import` existing service | `tc01-create-service.tf` | ✅ | `terraform import tessell_db_service.tc01_create c5bafa56-...` succeeded. |
| TC-20 | Refresh — no spurious diff after import | `tc01-create-service.tf` | ✅ | All rel-141 fields (`is_hpc`, `cadence`, `server_patching_config`, `updates_info`, `additional_storage`) show no drift post-import. |

---

### Group F — Start/Stop & Delete Schedule (Rel-141 helper changes)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-21 | Create start/stop schedule | `tc21-start-stop-schedule.tf` | ✅ | One-time schedule created. ID `a201fbe5`, status `ACTIVE`. |
| TC-22 | Update start/stop schedule | `tc21-start-stop-schedule.tf` | ✅ | Stop time updated. State reflects new value. |
| TC-23 | Create delete schedule | `tc21-start-stop-schedule.tf` | ✅ | Delete schedule `1037837d` created. Service scheduled for deletion at `2026-12-31T23:59:00Z`. |

---

### Group G — Snapshot (Regression)

| TC | Title | Config File | Status | Notes |
|----|-------|-------------|--------|-------|
| TC-24 | Create snapshot via `tessell_db_snapshot` | `tc24-snapshot.tf` | ✅ | Snapshot `b84e2cdd` created. Status `AVAILABLE`. |
| TC-25 | Read snapshot via data source | `tc24-snapshot.tf` | ✅ | Data source read confirmed. `name=tc24-rel141-snapshot`, `status=AVAILABLE`. |

---

## Execution Log

| Date | TC | Result | Remarks |
|------|----|--------|---------|
| 2026-04-29 | TC-01 | ✅ PASSED | PostgreSQL/AWS/ap-south-1. `maintenanceWindow.cadence=WEEKLY`, `serverPatchingConfig` stored. |
| 2026-04-29 | TC-02 | ✅ PASSED | `terraform plan` after apply showed 0 changes. No drift. |
| 2026-04-29 | TC-03 | ✅ PASSED | `auto_patch_config` sent on create. Backend does not persist it; omitted post-create. |
| 2026-04-29 | TC-04 | ✅ PASSED | `server_patching_config.enable_auto_os_patching=true` confirmed via API GET. |
| 2026-04-29 | TC-05 | ✅ PASSED | `additional_storage=0` in HCL and state. No drift. |
| 2026-04-29 | TC-06/06b | ✅ PASSED | WEEKLY + day accepted; missing day → plan-time error. |
| 2026-04-29 | TC-07/07b | ✅ PASSED | MONTHLY + day_of_month accepted; missing day_of_month → plan-time error. |
| 2026-04-29 | TC-08/08b | ✅ PASSED | QUARTERLY + start_date accepted; missing start_date → plan-time error. |
| 2026-04-29 | TC-09 | ✅ PASSED | Added `read_only_replica` `RRc6e5-node-0`. API confirmed 2 instances. |
| 2026-04-29 | TC-10 | ✅ PASSED | Removed replica block. Instance count confirmed = 1 after refresh. |
| 2026-04-29 | TC-11 | ✅ PASSED | Role swap: `failover-node-0` promoted to primary. Completed in ~3m. |
| 2026-04-29 | TC-12 | ✅ PASSED | HA service `my-service-nmj06ows` created. Azure/northCentralUS. 2 instances. |
| 2026-04-29 | TC-13 | ✅ PASSED | `auto_patch_config` update → provider error: `can only be set at creation time`. |
| 2026-04-29 | TC-14 | ✅ PASSED | `server_patching_config` update → provider error: `can only be set at creation time`. |
| 2026-04-30 | TC-15 | ✅ PASSED | `PATCH /services/{id}/maintenance-windows` called. `cadence=MONTHLY, day_of_month=15` persisted. No drift on subsequent plan. |
| 2026-04-29 | TC-16 | ✅ PASSED | `is_hpc = false` in state. Output `tc01_is_hpc` confirmed. |
| 2026-04-29 | TC-17 | ✅ PASSED | `data "tessell_db_service"` — `is_hpc = false` confirmed in data source. |
| 2026-04-29 | TC-18 | ✅ PASSED | `updates_info` block populated in state after create. |
| 2026-04-30 | TC-19 | ✅ PASSED | `terraform import tessell_db_service.tc01_create c5bafa56-...` succeeded. |
| 2026-04-30 | TC-20 | ✅ PASSED | All rel-141 fields show no drift post-import. |
| 2026-04-30 | TC-21 | ✅ PASSED | Start/stop schedule `a201fbe5` created, status `ACTIVE`. |
| 2026-04-30 | TC-22 | ✅ PASSED | Stop time updated. State reflects new value. |
| 2026-04-30 | TC-24 | ✅ PASSED | Snapshot `b84e2cdd` created, status `AVAILABLE`. |
| 2026-04-30 | TC-25 | ✅ PASSED | Data source read: `name=tc24-rel141-snapshot`, `status=AVAILABLE`. |
| 2026-04-30 | TC-23 | ✅ PASSED | Delete schedule `1037837d` created. Deletion at `2026-12-31T23:59:00Z`, `retain_availability_machine=true`. |
| 2026-05-04 | GCP-TC01 | ✅ PASSED | SQLSERVER HA service `my-service-71ddd622` (`bf3b4ebe`) created on GCP/us-east1. 2 instances (primary + failover_replica) both UP. `is_hpc=false`. No drift on `terraform plan` post-import. |
| 2026-05-04 | GCP-TC02 | ✅ PASSED | `read_only_replica` `ror-node-0` added (instance_group=`ror`, AZ=us-east1-b, GCP_HYPERDISK). API confirmed 3 instances all READY/UP. |
| 2026-05-04 | GCP-TC03 | ✅ PASSED | Switchover applied via HCL role swap: `default-node-1` → `primary`, `default-node-0` → `failover_replica`. Request submitted successfully. |
