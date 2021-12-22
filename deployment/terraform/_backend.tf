terraform {
  backend "s3" {
    bucket         = "butler-terraform-backend"
    key            = "butler-platform"
    region         = "us-east-2"
    encrypt        = true
    dynamodb_table = "terraform-lock"
  }
}