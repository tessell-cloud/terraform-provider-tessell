# Create a Tessell DB snapshot from an AvailabilityMachine
resource "tessell_db_snapshot" "example" {
  name                    = "my-test-db-snapshot"
  availability_machine_id = "6ca375ef-0514-43df-bee6-9167d0c49a9a"
}