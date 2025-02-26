resource "aws_eip" "nat_gw" {
  domain = "vpc"

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${var.env}-${var.project}-eip"
  }
}
