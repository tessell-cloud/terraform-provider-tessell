# Create a PostgreSQL parameter profile with custom parameter values
resource "tessell_parameter_profile" "example" {
  name        = "my-postgres-profile"
  description = "Custom PostgreSQL parameter profile"
  engine_type = "POSTGRESQL"
  db_version  = "16"

  engine_info {
    edition = "COMMUNITY"
  }

  parameters {
    name  = "max_connections"
    value = "600"
  }

  parameters {
    name  = "work_mem"
    value = "8192"
  }

  propagation_policy {
    strategy = "IMMEDIATELY"
  }
}
