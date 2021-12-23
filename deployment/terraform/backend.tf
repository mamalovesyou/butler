locals {
  create_backend = {
    "global" = 1
    "prod"   = 0
  }

}

module "tf_backend" {
  count = lookup(local.create_backend, terraform.workspace)

  source               = "git::https://github.com/DNXLabs/terraform-aws-backend?ref=1.2.1"
  bucket_prefix        = "butler"
  bucket_sse_algorithm = "AES256"
  workspaces           = ["global", "prod"]
}


terraform {
  backend "s3" {
    bucket         = "butler-terraform-backend"
    key            = "butler-platform"
    region         = "us-east-2"
    encrypt        = true
    dynamodb_table = "terraform-lock"
  }
}