output "users_db_address" {
  description = "Users service database address"
  value       = module.users_db[0].db_address
}

output "cert_manager_irsa_role_arn" {
  value = module.eks_cluster[0].cert_manager_irsa_role_arn
}

output "kubeconfig_base64" {
  value = base64encode(module.eks_cluster[0].kubeconfig)
}