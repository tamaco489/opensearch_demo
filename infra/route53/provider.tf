provider "aws" {}

terraform {
  required_version = "1.9.5"
  backend "s3" {
    bucket = "dev-opensearch-demo-tfstate"
    key    = "route53/terraform.tfstate"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}
