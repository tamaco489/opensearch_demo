variable "env" {
  description = "The environment in which the Route53 will be created"
  type        = string
  default     = "dev"
}

variable "product" {
  description = "The product name"
  type        = string
  default     = "opensearch-demo"
}

locals {
  fqn = "${var.env}-${var.product}"
}
