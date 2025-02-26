resource "aws_security_group" "shop_api" {
  name        = "${local.fqn}-api"
  description = "Security group for API service running in the VPC"
  vpc_id      = data.terraform_remote_state.network.outputs.vpc.id

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${local.fqn}-api"
  }
}

resource "aws_vpc_security_group_egress_rule" "shop_api_egress" {
  security_group_id = aws_security_group.shop_api.id
  description       = "Allow Lambda to access external resources (e.g., RDS, OpenSearch, APIs)"
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}
