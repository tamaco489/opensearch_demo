resource "aws_route53_zone" "opensearch_demo" {
  name    = var.domain
  comment = "OpenSearchの検証で利用"
  tags = {
    Env     = var.env
    Project = var.product
    Name  = "${local.fqn}"
  }
}
