resource "aws_ecr_repository" "shop_api" {
  name = "${local.fqn}-api"

  # 既存のタグに対して、後から上書きを可能とする設定
  image_tag_mutability = "MUTABLE"

  # イメージがpushされる度に、自動的にセキュリティスキャンを行う設定を有効にする
  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Env     = var.env
    Project = var.project
    Name  = "${local.fqn}-api"
  }
}

# ライフサイクルポリシーの設定
resource "aws_ecr_lifecycle_policy" "shop_api" {
  repository = aws_ecr_repository.shop_api.name

  policy = jsonencode(
    {
      "rules" : [
        {
          "rulePriority" : 1,
          "description" : "バージョン付きのイメージを5個保持する、6個目がアップロードされた際には古いものから順に削除されていく",
          "selection" : {
            "tagStatus" : "tagged",
            "tagPrefixList" : ["shop_api_v"],
            "countType" : "imageCountMoreThan",
            "countNumber" : 5
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 2,
          "description" : "タグが設定されていないイメージをアップロードされてから3日後に削除する",
          "selection" : {
            "tagStatus" : "untagged",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 3
          },
          "action" : {
            "type" : "expire"
          }
        },
        {
          "rulePriority" : 3,
          "description" : "タグが設定されたイメージをアップロードされてから7日後に削除する",
          "selection" : {
            "tagStatus" : "any",
            "countType" : "sinceImagePushed",
            "countUnit" : "days",
            "countNumber" : 7
          },
          "action" : {
            "type" : "expire"
          }
        }
      ]
    }
  )
}
