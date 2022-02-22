variable "prefix" {
  default     = "butler"
  description = "Prefix used to name the eks cluster"
}

variable "cluster_name" {
  description = "EKS Cluster name"
}

variable "private_subnets" {
  type        = list(string)
  description = "Private subnets list"
}

variable "vpc_id" {
  type        = string
  description = "VPC ID"
}

variable "map_users" {
  description = "Additional IAM users to add to the aws-auth configmap."
  type = list(object({
    userarn  = string
    username = string
    groups   = list(string)
  }))

  default = [
    {
      userarn  = "arn:aws:iam::135314574862:user/github-ci"
      username = "github-ci"
      groups   = ["system:masters"]
    },
    {
      userarn  = "arn:aws:iam::135314574862:user/matthieu"
      username = "matthieu"
      groups   = ["system:masters"]
    },
  ]
}