package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) CreateLikeProductComment(ctx *gin.Context, request gen.CreateLikeProductCommentRequestObject) (gen.CreateLikeProductCommentResponseObject, error) {

	return gen.CreateLikeProductComment204Response{}, nil
}
