package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateProductComment は商品に対してコメントを投稿します。
//
// タイトル、本文、商品に対する評価の3つがリクエストされます。
//
// タイトル、本文は空文字以外を許容します。
//
// 商品に対する評価は1～5までの数値型を許容します。
func (c *Controllers) CreateProductComment(ctx *gin.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	err := validation.ValidateStruct(request.Body,
		validation.Field(
			&request.Body.Title,
			validation.Required,
		),
		validation.Field(
			&request.Body.Content,
			validation.Required,
		),
	)
	if err != nil {
		_ = ctx.Error(err)
		return gen.CreateProfile400Response{}, nil
	}

	// `rate` のバリデーション
	// uint32 → int にcastする必要があるため、一度ローカル変数で定義してそのポインタを渡す必要がある。
	// また、ローカル変数で定義しているため、ValidateStructではなく、Validateで個別のバリデーションチェックを行う。
	rate := int(request.Body.Rate)
	err = validation.Validate(
		&rate,
		validation.Required,
		validation.Min(1),
		validation.Max(5),
	)
	if err != nil {
		_ = ctx.Error(err)
		return gen.CreateProfile400Response{}, nil
	}

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
