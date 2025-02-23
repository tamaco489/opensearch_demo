package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

// GetProductCommentViolations は商品に対して投稿されたコメントの中で、予め定めたNGワードに該当するデータを取得します。
func (c *Controllers) GetProductCommentViolations(ctx *gin.Context, request gen.GetProductCommentViolationsRequestObject) (gen.GetProductCommentViolationsResponseObject, error) {

	res, err := c.productCommentUseCase.GetProductCommentViolations(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return gen.GetProductCommentViolations500Response{}, err
	}

	return res, nil
}
