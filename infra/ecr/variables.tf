variable "env" {
  description = "The environment in which the ecr will be created"
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
