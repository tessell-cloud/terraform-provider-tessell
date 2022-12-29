# List all tessell_db_service
data_source "tessell_db_services" "example" {}

# Get a tessell_db_service using the service name
data_source "tessell_db_services" "example" {
  name = "my-test-db"
}