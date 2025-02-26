variable "env" {
  description = "The environment in which the open search will be created"
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

variable "index" {
  description = "The open search index name"
  type        = string
  default     = "product_comments"
}

locals {
  fqn             = "${var.env}-${var.product}"
  collection_name = "${var.env}-${var.project}-collection"
}
