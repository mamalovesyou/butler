output "users_db_address" {
  description = "Users service database address"
  value       = module.users_db.db_address
}

output "cert_manager_irsa_role_arn" {
  value = module.k8s.cert_manager_irsa_role_arn
}

output "kubeconfig_base64" {
  value = base64encode(module.k8s.kubeconfig)
}