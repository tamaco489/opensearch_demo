data "terraform_remote_state" "ecr" {
  backend = "s3"
  config = {
    bucket = "${var.env}-opensearch-demo-tfstate"
    key    = "ecr/terraform.tfstate"
  }
}

data "terraform_remote_state" "lambda" {
  backend = "s3"
  config = {
    bucket = "${var.env}-opensearch-demo-tfstate"
    key    = "lambda/terraform.tfstate"
  }
}

data "terraform_remote_state" "acm" {
  backend = "s3"
  config = {
    bucket = "${var.env}-opensearch-demo-tfstate"
    key    = "acm/terraform.tfstate"
  }
}
