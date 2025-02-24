variable "env" {
  description = "The environment in which the Route53 will be created"
  type        = string
  default     = "dev"
}

variable "project" {
  description = "The project name"
  type        = string
  default     = "opensearch-demo"
}

variable "product" {
  description = "The product name"
  type        = string
  default     = "shop"
}

variable "region" {
  description = "The region in which the Route53 will be created"
  type        = string
  default     = "ap-northeast-1"
}

variable "domain" {
  description = "The domain name"
  type        = string
  default     = "example.com"
}

locals {
  fqn = "${var.env}-${var.product}"
}
