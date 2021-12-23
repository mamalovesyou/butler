locals {
  create_cluster = {
    "global" = 0
    "prod"   = 1
  }

  cluster_name = "${var.prefix}-cluster-${terraform.workspace}"
}

module "network" {
  source = "./modules/network"

  count = lookup(local.create_cluster, terraform.workspace)

  prefix       = var.prefix
  cluster_name = local.cluster_name
}

module "eks_cluster" {
  source = "./modules/eks"

  count           = lookup(local.create_cluster, terraform.workspace)
  cluster_name    = local.cluster_name
  prefix          = var.prefix
  vpc_id          = module.network[0].vpc_id
  private_subnets = module.network[0].private_subnets
}

module "users_db" {
  source = "./modules/postgres"

  count = lookup(local.create_cluster, terraform.workspace)

  prefix           = var.prefix
  vpc_id           = module.network[0].vpc_id
  database_subnets = module.network[0].database_subnets
  vpc_cidr_block   = module.network[0].vpc_cidr_block
  db_name          = var.users_db_name
  db_user          = var.users_db_user
  db_password      = var.users_db_password
}