terraform {
    required_version = ">= 0.13.0"
}

provider "aws" {
  region = var.region
}

module "networking" {
  source = "./modules/network"

  prefix = var.prefix
  eks_cluster_name = module.k8s.eks_cluster_name
}

module "k8s" {
  source = "./modules/k8s"

  prefix = var.prefix
  vpc_id = module.networking.vpc_id
  private_subnets = module.networking.private_subnets
}

module "users_db" {
  source = "./modules/postgres"

  prefix = var.prefix
  vpc_id = module.networking.vpc_id
  database_subnets = module.networking.database_subnets
  vpc_cidr_block = module.networking.vpc_cidr_block
  db_name = var.users_db_name
  db_user = var.users_db_user
  db_password = var.users_db_password
}