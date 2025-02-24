output "host_zone" {
  value = {
    id   = aws_route53_zone.opensearch_demo.id,
    name = aws_route53_zone.opensearch_demo.name
  }
}
