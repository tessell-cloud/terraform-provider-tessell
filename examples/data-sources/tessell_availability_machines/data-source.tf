# Get all existing Availability Machines
data_source "tessell_availability_machines" "example" {}

# Get all Oracle Availability Machines
data_source "tessell_availability_machines" "example" {
  engine_type = "ORACLE"
}