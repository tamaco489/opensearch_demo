package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) DeleteProductCommentViolationByID(ctx *gin.Context, request gen.DeleteProductCommentViolationByIDRequestObject) (gen.DeleteProductCommentViolationByIDResponseObject, error) {

	res, err := c.productCommentUseCase.DeleteProductCommentViolationsByID(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return gen.DeleteProductCommentViolationByID500Response{}, err
	}

	return res, nil
}
