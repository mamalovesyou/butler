output "users_db_address" {
  description = "Users service database address"
  value       = module.users_db.db_address
}

output "cert_manager_irsa_role_arn" {
  value = module.eks_cluster.cert_manager_irsa_role_arn
}