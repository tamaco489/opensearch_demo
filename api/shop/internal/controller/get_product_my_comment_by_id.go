package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetProductMyCommentByID(ctx *gin.Context, request gen.GetProductMyCommentByIDRequestObject) (gen.GetProductMyCommentByIDResponseObject, error) {

	return gen.GetProductMyCommentByID200JSONResponse{}, nil
}
