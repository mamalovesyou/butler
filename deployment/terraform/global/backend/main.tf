
module "tf_backend" {
  source = "git::https://github.com/DNXLabs/terraform-aws-backend?ref=1.2.1"

  bucket_prefix = "butler"
  bucket_sse_algorithm = "AES256"
  workspaces           = ["prod"]
}