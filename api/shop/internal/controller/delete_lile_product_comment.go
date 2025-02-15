package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) DeleteLikeProductComment(ctx *gin.Context, request gen.DeleteLikeProductCommentRequestObject) (gen.DeleteLikeProductCommentResponseObject, error) {

	return gen.DeleteLikeProductComment204Response{}, nil
}
