locals {
  cluster_name = "${var.prefix}-cluster-${terraform.workspace}"
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

module "users_db" {
  source = "./modules/postgres"

  prefix           = var.prefix
  vpc_id           = module.network.vpc_id
  database_subnets = module.network.database_subnets
  vpc_cidr_block   = module.network.vpc_cidr_block
  db_name          = var.users_db_name
  db_user          = var.users_db_user
  db_password      = var.users_db_password
}