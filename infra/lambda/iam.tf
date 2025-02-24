resource "aws_iam_policy" "lambda_logging" {
  name        = "${var.env}-${var.project}-lambda-logging-iam-policy"
  path        = "/"
  description = "IAM policy granting permissions for Lambda logging"
  policy      = data.aws_iam_policy_document.lambda_logging.json

  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${var.env}-${var.project}" // api, batch 双方で利用する想定のため
  }
}
