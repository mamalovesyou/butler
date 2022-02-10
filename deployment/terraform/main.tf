locals {
  cluster_name = "${var.prefix}-cluster-${terraform.workspace}"
}

# Secrets
resource "aws_secretsmanager_secret" "env_variables" {
  name = "${terraform.workspace}-${var.prefix}-env"
}

# Versions
resource "aws_secretsmanager_secret_version" "env_variables_value" {
  secret_id     = aws_secretsmanager_secret.env_variables.id
  secret_string = jsonencode({})
}

module "network" {
  source       = "./modules/network"
  prefix       = var.prefix
  cluster_name = local.cluster_name
}

module "eks_cluster" {
  source = "./modules/eks"

  cluster_name    = local.cluster_name
  prefix          = var.prefix
  vpc_id          = module.network.vpc_id
  private_subnets = module.network.private_subnets
}


module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4"

  name        = "${var.prefix}-rds-security-group-${terraform.workspace}"
  description = "Complete PostgreSQL example security group"
  vpc_id      = module.network.vpc_id

  # ingress
  ingress_with_cidr_blocks = [
    {
      from_port   = 5432
      to_port     = 5432
      protocol    = "tcp"
      description = "PostgreSQL access from within VPC"
      cidr_blocks = module.network.vpc_cidr_block
    },
  ]

  tags = {
    Environment = terraform.workspace
  }
}


module "users_db" {
  source = "./modules/postgres"

  prefix             = var.prefix
  vpc_id             = module.network.vpc_id
  database_subnets   = module.network.database_subnets
  vpc_cidr_block     = module.network.vpc_cidr_block
  db_name            = var.users_db_name
  db_user            = var.users_db_user
  db_passowrd        = var.users_db_password
  security_group_ids = [module.security_group.security_group_id]
}

module "airbyte_db" {
  source = "./modules/postgres"

  prefix             = var.prefix
  vpc_id             = module.network.vpc_id
  database_subnets   = module.network.database_subnets
  vpc_cidr_block     = module.network.vpc_cidr_block
  db_name            = var.airbyte_db_name
  db_user            = var.airbyte_db_user
  db_passowrd        = var.airbyte_db_password
  security_group_ids = [module.security_group.security_group_id]
}