variable "prefix" {
  default = "butler"
  description = "Prefix used to name the eks cluster"
}

variable "db_name" {
  type = string
  description = "Database name"
}

variable "db_user" {
  type = string
  description = "Database username"
}

variable "db_password" {
  type = string
  description = "Database password"
}

variable "instance_type" {
  type = string
  default = "db.t2.micro"
  description = "DB instance type"
}

variable "vpc_id" {
  type = string
  description = "VPC ID"
}

variable "database_subnets" {
  type = list(string)
  description = "Database private subnets"
}

variable "vpc_cidr_block" {
  type = string
  description = "VPC CIDR Block"
}

