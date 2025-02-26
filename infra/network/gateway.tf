resource "aws_internet_gateway" "opensearch_demo" {
  vpc_id = aws_vpc.opensearch_demo.id

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${var.env}-${var.project}-internet-gw"
  }
}

# コスト削減のため、NAT Gatewayは1つにする
resource "aws_nat_gateway" "opensearch_demo" {
  allocation_id = aws_eip.nat_gw.id
  # subnet_id     = aws_subnet.public_subnet["a"].id # NOTE: 外部へ出ていくときもしかしたらpublicに配置しないと動かないかもしれない
  subnet_id     = aws_subnet.private_subnet["a"].id

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${var.env}-${var.project}-nat-gw"
  }
}
