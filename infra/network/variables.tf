# =================================================================
# general
# =================================================================
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

variable "region" {
  description = "The region in which the network will be created"
  type        = string
  default     = "ap-northeast-1"
}

locals {
  fqn = "${var.env}-${var.product}"
}


# =================================================================
# network
# =================================================================
variable "vpc_cidr_block" {
  description = "The cidr block for the vpc"
  type        = string
  default     = "10.1.0.0/16"
}

# variable "public_subnet" {
#   description = "The cidr block for the public subnet"
#   type        = map(map(string))
#   default = {
#     a = {
#       az   = "a"
#       cidr = "10.1.11.0/24"
#     }
#     d = {
#       az   = "d"
#       cidr = "10.1.12.0/24"
#     }
#   }
# }

variable "private_subnet" {
  description = "The cidr block for the private subnet"
  type        = map(map(string))
  default = {
    a = {
      az   = "a"
      cidr = "10.1.21.0/24"
    }
    d = {
      az   = "d"
      cidr = "10.1.22.0/24"
    }
  }
}
