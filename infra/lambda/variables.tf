variable "env" {
  description = "The environment in which the lambda logging iam policy will be created"
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

locals {
  fqn = "${var.env}-${var.product}"
}
