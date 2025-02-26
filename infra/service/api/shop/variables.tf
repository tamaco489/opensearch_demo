variable "env" {
  description = "The environment in which the shop api lambda will be created"
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
  description = "The region in which the shop api lambda will be created"
  type        = string
  default     = "ap-northeast-1"
}

variable "aws_account_id" {
  description = "The aws account id"
  type        = string
  default     = "1234567890"
}

locals {
  fqn             = "${var.env}-${var.product}"
  collection_name = "${var.env}-${var.project}-collection"
}
