# Provision Postgresql Single Instance DB service
resource "tessell_db_service" "example" {
  name                       = "finance-db"
  description                = "Database for finance department"
  subscription               = "default"
  engine_type                = "POSTGRESQL"
  topology                   = "single_instance"
  software_image             = "PostgreSQL 13"
  software_image_version     = "PostgreSQL 13.3"
  auto_minor_version_update  = true
  enable_deletion_protection = false

  infrastructure {
    cloud              = "aws"
    region             = "ap-south-1"
    availability_zone  = "ap-south-1a"
    vpc                = "default-vpc-1234"
    compute_type       = "t2.small"
    additional_storage = 0
    encryption_key     = "finance-db-encyption-key-with-salt"
  }

  service_connectivity {
    service_port         = "5432"
    enable_public_access = true
    allowed_ip_addresses = [
      "11.22.33.44"
    ]
  }

  creds {
    master_user     = "master"
    master_password = "MyPassword@123"
  }

  maintenance_window {
    day      = "Sunday"
    time     = "02:00"
    duration = 30
  }

  snapshot_configuration {
    auto_snapshot = true
    sla           = "2-days-pitr"
    snapshot_window {
      time     = "01:00"
      duration = 30
    }
  }

  engine_configuration {
    postgresql_config {
      parameter_profile = "PostgreSQL 13 Profile"
    }
  }

  databases {
    database_name = "db1"
    database_configuration {
      postgresql_config {
        parameter_profile = "PostgreSQL 13 Profile"
      }
    }
  }

  tags {
    name  = "department"
    value = "finance"
  }
  tags {
    name  = "finance"
    value = ""
  }
}

# Create a clone of a Oracle DB service
resource "tessell_db_service" "example" {
  name                       = "finance-db-clone"
  description                = "Clone of finance database for QA use"
  subscription               = "default"
  engine_type                = "ORACLE"
  topology                   = "single_instance"
  software_image             = "Oracle 12c"
  software_image_version     = "12.1.0.2.0"
  auto_minor_version_update  = true
  enable_deletion_protection = false

  infrastructure {
    cloud              = "aws"
    region             = "ap-south-1"
    availability_zone  = "ap-south-1a"
    vpc                = "default-vpc-1234"
    compute_type       = "m5.large"
    additional_storage = 0
  }

  service_connectivity {
    service_port         = "1521"
    enable_public_access = true
    allowed_ip_addresses = [
      "11.22.33.44"
    ]
  }

  creds {
    master_user     = "master"
    master_password = "MyPassword@123"
  }

  snapshot_configuration {
    auto_snapshot = true
    sla           = "2-days-pitr"
    snapshot_window {
      time     = "01:00"
      duration = 30
    }
  }

  engine_configuration {
    oracle_config {
      parameter_profile      = "Oracle Parameter Profile"
      options_profile        = "Oracle 12.1.0.2.0 Options Profile"
      character_set          = "AL32UTF8"
      national_character_set = "AL16UTF16"
    }
  }

  databases {
    source_database_id = "14107367-fbc3-4f3d-8cc3-b5a7e5910e08"
    database_name      = "orcl"
    database_configuration {
      oracle_config {
        parameter_profile = "Oracle Parameter Profile"
        options_profile   = "Oracle 12.1.0.2.0 Options Profile"
      }
    }
  }

  tags {
    name  = "department"
    value = "finance"
  }
  tags {
    name  = "finance-qa"
    value = ""
  }
  snapshot_id = "4592803a-a8a3-4a5e-be4c-ddf1df654242"
}