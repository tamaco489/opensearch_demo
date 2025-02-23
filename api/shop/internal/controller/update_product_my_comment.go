package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) UpdateProductMyComment(ctx *gin.Context, request gen.UpdateProductMyCommentRequestObject) (gen.UpdateProductMyCommentResponseObject, error) {

	return gen.UpdateProductMyComment204Response{}, nil
}
