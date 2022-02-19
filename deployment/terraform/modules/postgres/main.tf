locals {
  identifier = "${var.prefix}-${replace(var.db_name, "_", "-")}-${terraform.workspace}"
}

module "postgres_db" {

  source  = "terraform-aws-modules/rds/aws"
  version = "~> 3.0"

  #  option_group_name = "${terraform.workspace}-${var.db_name}"
  #  option_group_use_name_prefix = true
  #  create_db_option_group    = false

  identifier = local.identifier

  create_db_option_group    = false
  create_db_parameter_group = false

  engine            = "postgres"
  engine_version    = "13.3"
  instance_class    = var.instance_type
  allocated_storage = 5

  name     = var.db_name
  username = var.db_user
  password = var.db_passowrd
  port     = "5432"

  multi_az               = true
  subnet_ids             = var.database_subnets
  vpc_security_group_ids = var.security_group_ids

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  apply_immediately = true

  backup_retention_period = 10
  skip_final_snapshot     = true
  deletion_protection     = false

  performance_insights_enabled          = true
  performance_insights_retention_period = 7
  monitoring_role_name                  = "${terraform.workspace}-${var.db_name}-monitoring-role"
  create_monitoring_role                = true
  monitoring_interval                   = 60

  parameters = [
    {
      name  = "autovacuum"
      value = 1
    },
    {
      name  = "client_encoding"
      value = "utf8"
    }
  ]

  tags = {
    Environment = terraform.workspace
  }
}