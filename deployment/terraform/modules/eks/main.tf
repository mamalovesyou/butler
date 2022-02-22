provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  token                  = data.aws_eks_cluster_auth.cluster.token
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
}

module "eks" {
  source                 = "terraform-aws-modules/eks/aws"
  version                = "~> 17.24"
  cluster_name           = var.cluster_name
  cluster_version        = "1.21"
  subnets                = var.private_subnets
  cluster_create_timeout = "30m" # need to increase module defaults
  write_kubeconfig       = false # Disabled permanent writing of config files

  vpc_id = var.vpc_id

  tags = {
    Environment  = terraform.workspace
    Organization = "butlerhq"
  }

  node_groups_defaults = {
    ami_type  = "AL2_x86_64"
    disk_size = 50
  }

  node_groups = {
    t3Large = {
      desired_capacity = 1
      max_capacity     = 10
      min_capacity     = 1

      subnets = var.private_subnets

      instance_types = ["t3.large"]
      capacity_type  = "SPOT"
      k8s_labels = {
        Environment = "${terraform.workspace}"
      }
    }
  }

  enable_irsa = true
  map_users   = var.map_users
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
    "Version" : "2012-10-17",

    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : "route53:GetChange",
        "Resource" : "arn:aws:route53:::change/*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "route53:ChangeResourceRecordSets",
          "route53:ListResourceRecordSets"
        ],
        "Resource" : "arn:aws:route53:::hostedzone/*"
      },
      {
        "Effect" : "Allow",
        "Action" : "route53:ListHostedZonesByName",
        "Resource" : "*"
      }
    ]
  })
}
