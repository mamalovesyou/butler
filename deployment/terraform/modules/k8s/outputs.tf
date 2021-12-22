
output "eks_cluster_name" {
  description = "EKS cluster name"
  value       = local.cluster_name
}

output "cert_manager_irsa_role_arn" {
  value = module.cert_manager_irsa.this_iam_role_arn
}

output "kubeconfig" {
  value = module.eks.kubeconfig
}