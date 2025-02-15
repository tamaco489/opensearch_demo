package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) CreateProductComment(ctx *gin.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
