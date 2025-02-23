package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) DeleteProductMyComment(ctx *gin.Context, request gen.DeleteProductMyCommentRequestObject) (gen.DeleteProductMyCommentResponseObject, error) {

	return gen.DeleteProductMyComment204Response{}, nil
}
