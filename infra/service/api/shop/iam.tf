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

# DynamoDBへのアクセス権を定義するポリシードキュメントを生成する
# data "aws_iam_policy_document" "shop_api" {
#   statement {
#     effect = "Allow"
#     actions = [
#       "dynamodb:PutItem",
#       "dynamodb:Query",
#       "dynamodb:UpdateItem",
#       "dynamodb:BatchWriteItem"
#     ]
#     resources = ["${aws_dynamodb_table.user_table.arn}"]
#   }
# }

# DynamoDBアクセス権限をIAM Roleに付与する
# resource "aws_iam_role_policy" "shop_api" {
#   name   = "${local.fqn}-api-dynamodb-access-policy"
#   role   = aws_iam_role.shop_api.id
#   policy = data.aws_iam_policy_document.shop_api.json
# }


# NOTE: VPC Lambda として稼働させるために最低限必要になる権限を関連付ける（ENI作成、削除、CloudWatch Logsへの書き込み等）
# https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AWSLambdaVPCAccessExecutionRole.html
resource "aws_iam_role_policy_attachment" "shop_api_execution_role" {
  role       = aws_iam_role.shop_api.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}
