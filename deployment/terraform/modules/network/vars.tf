variable "prefix" {
  default     = "butler"
  description = "Prefix the vpc name"
}

variable "eks_cluster_name" {
  type = string
  description = "The eks cluster name"
}

variable "cidr" {
  type = string
  default = "10.0.0.0/16"
  description = "VPC cidr"
}

variable "private_subnets" {
  type = list(string)
  default = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  description = "VPC private subnets"
}

variable "public_subnets" {
  type = list(string)
  default = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]
  description = "VPC public subnets"
}

variable "database_subnets" {
  type = list(string)
  default = ["10.0.7.0/24", "10.0.8.0/24", "10.0.9.0/24"]
  description = "VPC database subnets"
}

