variable "region" {
  default     = "us-east-2"
  description = "AWS region"
}

variable "prefix" {
  default     = "butler"
  description = "Prefix used to name aws entities"
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
  description = "Users database password"
}

variable "connectors_db_name" {
  default     = "connectors_db"
  description = "Connectors service database name"
}

variable "connectors_db_user" {
  default     = "butler"
  description = "Connectors service database user"
}

variable "connectors_db_password" {
  description = "Connectors database password"
}

variable "airbyte_db_name" {
  default     = "airbyte_db"
  description = "Airbyte service database name"
}

variable "airbyte_db_user" {
  default     = "butler"
  description = "Airbyte service database user"
}

variable "airbyte_db_password" {
  description = "Airbyte database password"
}

variable "eks_map_users" {
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