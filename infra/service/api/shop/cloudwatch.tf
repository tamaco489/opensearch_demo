resource "aws_cloudwatch_log_group" "shop_api" {
  name              = "/aws/lambda/${aws_lambda_function.shop_api.function_name}"
  retention_in_days = 3

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${local.fqn}-api"
  }
}
