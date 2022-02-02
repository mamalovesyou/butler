variable "region" {
  default     = "us-east-2"
  description = "AWS region"
}

variable "prefix" {
  default     = "butler"
  description = "Prefix used to name aws entities"
}

variable "assume_developer_role" {
  description = "A list of ARN's of users/roles that can assume the cluster_admin role"
  type        = list(string)
  default = [
    "arn:aws:iam::135314574862:user/github-ci",
    "arn:aws:iam::135314574862:user/matthieu"
  ]
}

variable "users_db_name" {
  default     = "users_db"
  description = "Users service database name"
}

variable "users_db_user" {
  default     = "butler"
  description = "Users service database user"
}

variable "users_db_password" {
  type        = string
  description = "Users service database password"
}