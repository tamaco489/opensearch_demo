resource "aws_route53_zone" "opensearch_demo" {
  name    = var.domain
  comment = "OpenSearchの検証で利用"
  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${var.env}-${var.project}" // api, batch 双方で利用する想定のため
  }
}
