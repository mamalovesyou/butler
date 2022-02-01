data "aws_caller_identity" "current" {}
data "aws_partition" "current" {}

locals {
  role_to_user_map = {
    eks_admin = "admin",
  }

  role_map_obj = [
    for role_name, user in local.role_to_user_map : {
      rolearn  = "arn:${data.aws_partition.current.partition}:iam::${data.aws_caller_identity.current.account_id}:role/${role_name}"
      username = user
      groups   = contains(tolist([user]), "admin") ? tolist(["system:masters"]) : tolist(["none"])
    }
  ]
}

resource "aws_iam_role" "cluster_admin" {
  name = "${var.cluster_name}-admin"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          AWS = var.assume_developer_role
        }
      },
    ]
  })

  inline_policy {
    name = "${var.cluster_name}-admin-policy"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action   = ["eks:DescribeCluster"]
          Effect   = "Allow"
          Resource = "*"
        },
      ]
    })
  }
}
