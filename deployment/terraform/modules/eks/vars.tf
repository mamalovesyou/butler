variable "prefix" {
  default = "butler"
  description = "Prefix used to name the eks cluster"
}

variable "cluster_name" {
  description = "EKS Cluster name"
}

variable "private_subnets" {
  type = list(string)
  description = "Private subnets list"
}

variable "vpc_id" {
  type = string
  description = "VPC ID"
}
