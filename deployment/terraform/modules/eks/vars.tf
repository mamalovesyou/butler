variable "prefix" {
  default     = "butler"
  description = "Prefix used to name the eks cluster"
}

variable "cluster_name" {
  description = "Name of the EKS cluster. Also used as a prefix in names of related resources"
  type        = string
}

variable "private_subnets" {
  type        = list(string)
  description = "Private subnets list"
}

variable "vpc_id" {
  type        = string
  description = "VPC ID"
}

variable "assume_developer_role" {
  description = "A list of ARN's of users/roles that can assume the cluster_admin role"
  type        = list(string)
}
