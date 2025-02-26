output "vpc" {
  value = {
    arn                = aws_vpc.opensearch_demo.arn
    id                 = aws_vpc.opensearch_demo.id
    cidr_block         = aws_vpc.opensearch_demo.cidr_block
    public_subnet_ids  = [for s in aws_subnet.public_subnet : s.id]
    private_subnet_ids = [for s in aws_subnet.private_subnet : s.id]
  }
}
