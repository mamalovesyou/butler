locals {
  cluster_name    = "${var.prefix}-eks-${terraform.workspace}"
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.cluster.token
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = local.cluster_name
  cluster_version = "1.21"
  subnets         = var.private_subnets

  vpc_id = var.vpc_id

  tags = {
    Environment = terraform.workspace
    Organization   = "butlerhq"
  }

  worker_groups = [
    {
      instance_type        = "t2.small"
      asg_max_size         = 6
      asg_desired_capacity = 3
      asg_min_size         = 1
      root_volume_type     = "gp2"
      subnets              = var.private_subnets
    }
  ]

  manage_aws_auth = true
  enable_irsa     = true
}

data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}

# IAM Role Service Account for the cert-manager
module "cert_manager_irsa" {
  source                        = "terraform-aws-modules/iam/aws//modules/iam-assumable-role-with-oidc"
  version                       = "3.6.0"
  create_role                   = true
  role_name                     = "cert_manager-irsa-${terraform.workspace}"
  provider_url                  = replace(module.eks.cluster_oidc_issuer_url, "https://", "")
  role_policy_arns              = [aws_iam_policy.cert_manager_policy.arn]
  oidc_fully_qualified_subjects = ["system:serviceaccount:cert-manager:cert-manager"]
}

resource "aws_iam_policy" "cert_manager_policy" {
  name        = "cert-manager-policy-${terraform.workspace}"
  path        = "/"
  description = "Policy, which allows CertManager to create Route53 records"

  policy = jsonencode({
    "Version": "2012-10-17",

    "Statement": [
      {
        "Effect": "Allow",
        "Action": "route53:GetChange",
        "Resource": "arn:aws:route53:::change/*"
      },
      {
        "Effect": "Allow",
        "Action": [
          "route53:ChangeResourceRecordSets",
          "route53:ListResourceRecordSets"
        ],
        "Resource": "arn:aws:route53:::hostedzone/*"
      },
      {
        "Effect": "Allow",
        "Action": "route53:ListHostedZonesByName",
        "Resource": "*"
      }
    ]
  })
}
