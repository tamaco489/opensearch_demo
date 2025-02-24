resource "aws_iam_policy" "lambda_logging" {
  name        = "${local.fqn}-lambda-logging-iam-policy"
  path        = "/"
  description = "IAM policy granting permissions for Lambda logging"
  policy      = data.aws_iam_policy_document.lambda_logging.json

  tags = {
    Env     = var.env
    Project = var.product
    Name    = "${local.fqn}-api"
  }
}
