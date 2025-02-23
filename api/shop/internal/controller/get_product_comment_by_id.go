package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetProductCommentByID(ctx *gin.Context, request gen.GetProductCommentByIDRequestObject) (gen.GetProductCommentByIDResponseObject, error) {

	res, err := c.productCommentUseCase.GetProductCommentByID(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return gen.GetProductCommentByID500Response{}, err
	}

	return res, nil
}
