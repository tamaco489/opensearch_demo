# AWS Lambdaが指定されたロールを引き受けるためのポリシードキュメントを生成する
# そのロールに付与された権限を使用できるようにするために必要な設定
# これにより、Lambda関数が特定のAWSリソースにアクセスするために必要な権限を持つことができるようになる
data "aws_iam_policy_document" "lambda_execution_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

# IAM Roleの作成
resource "aws_iam_role" "shop_api" {
  name               = "${local.fqn}-api-iam-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_assume_role.json
  tags = {
    Env     = var.env
    Project = var.project
    Name    = "${local.fqn}-api"
  }
}

# NOTE: VPC Lambda として稼働させるために最低限必要になる権限を関連付ける（ENI作成、削除、CloudWatch Logsへの書き込み等）
# https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AWSLambdaVPCAccessExecutionRole.html
resource "aws_iam_role_policy_attachment" "shop_api_execution_role" {
  role       = aws_iam_role.shop_api.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

# AWS IAMポリシーの作成
# このポリシーは、LambdaがOpenSearch Serverlessにアクセスできるようにする
# Lambda関数に対して、特定のインデックスに対する読み書き権限を付与
resource "aws_iam_policy" "opensearch_access_policy" {
  name        = "${var.env}-${var.project}-opensearch-access-policy"
  description = "Allow Lambda to access OpenSearch Serverless"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Action = [
        "aoss:ReadDocument",
        "aoss:WriteDocument"
      ],
      Resource = ["arn:aws:aoss:${var.region}:${var.aws_account_id}:collection/${local.collection_name}/index/product_comments"]
    }]
  })
}

# AWS Lambdaが指定されたロールを引き受けるためのポリシードキュメントを生成する
# これにより、Lambda関数が指定されたIAMロールを引き受け、指定されたリソースにアクセスするための権限を持つことができるようになる
# Lambda関数がOpenSearch Serverlessにアクセスするために必要な設定
resource "aws_iam_role_policy_attachment" "shop_api_opensearch_attach" {
  role       = aws_iam_role.shop_api.name
  policy_arn = aws_iam_policy.opensearch_access_policy.arn
}

