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