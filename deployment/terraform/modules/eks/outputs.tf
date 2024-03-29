output "cert_manager_irsa_role_arn" {
  value = module.cert_manager_irsa.this_iam_role_arn
}

output "kubeconfig" { value = module.eks.kubeconfig }
output "cluster_endpoint" { value = module.eks.cluster_endpoint }