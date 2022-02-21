output "users-db-address" {
  description = "Users service database address"
  value       = module.users_db.db_address
}

output "connectors-db-address" {
  description = "Connectors service database address"
  value       = module.connectors_db.db_address
}

output "airbyte-db-address" {
  description = "Airbyte service database address"
  value       = module.airbyte_db.db_address
}

output "cert_manager_irsa_role_arn" {
  value = module.eks_cluster.cert_manager_irsa_role_arn
}