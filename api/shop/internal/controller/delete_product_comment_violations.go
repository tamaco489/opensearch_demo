package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) DeleteProductCommentByID(ctx *gin.Context, request gen.DeleteProductCommentByIDRequestObject) (gen.DeleteProductCommentByIDResponseObject, error) {

	res, err := c.productCommentUseCase.DeleteProductCommentByID(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return gen.DeleteProductCommentByID500Response{}, err
	}

	return res, nil
}
