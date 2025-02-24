# Lambda関数をターゲットとするAPI Gateway (HTTP API) を作成
resource "aws_apigatewayv2_api" "shop_api" {
  name          = "${local.fqn}-api"
  description   = "ショップAPI"
  protocol_type = "HTTP"

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${local.fqn}-api"
  }
}

# Lambda と API Gateway を統合するための統合リソース
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id                 = aws_apigatewayv2_api.shop_api.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.shop_api.invoke_arn
  payload_format_version = "2.0"
}

# Lambda 関数に対するルーティング設定 (/shop/v1/{proxy+})
resource "aws_apigatewayv2_route" "shop_api_route" {
  api_id    = aws_apigatewayv2_api.shop_api.id
  route_key = "ANY /shop/v1/{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

# ステージ (デフォルトステージ)
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id      = aws_apigatewayv2_api.shop_api.id
  name        = "$default"
  auto_deploy = true
}

# API Gateway とカスタムドメインのマッピング
resource "aws_apigatewayv2_api_mapping" "coral_api_mapping" {
  api_id      = aws_apigatewayv2_api.shop_api.id
  domain_name = data.terraform_remote_state.acm.outputs.shop_apigatewayv2_domain_name.id
  stage       = aws_apigatewayv2_stage.default_stage.id
}
