
output "vpc_id" {
  description = "AWS vpc id"
  value       = module.network.vpc_id
}

output "private_subnets" {
  description = "AWS vpc private subnets"
  value       = module.network.private_subnets
}

output "database_subnets" {
  description = "AWS vpc database subnets"
  value       = module.network.database_subnets
}

output "vpc_cidr_block" {
  description = "AWS vpc cidr block"
  value       = module.network.vpc_cidr_block
}
