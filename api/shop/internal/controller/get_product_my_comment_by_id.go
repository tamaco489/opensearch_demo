package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetProductMyCommentByID(ctx *gin.Context, request gen.GetProductMyCommentByIDRequestObject) (gen.GetProductMyCommentByIDResponseObject, error) {

	jst := time.FixedZone("JST", 9*60*60)

	return gen.GetProductMyCommentByID200JSONResponse{
		Id:        54009221,
		Title:     "とても良い商品です",
		Content:   "この商品は非常に良いです。特にデザインが素晴らしい。",
		CreatedAt: time.Date(2025, 2, 15, 13, 45, 30, 0, jst),
		LikeCount: 15,
		ReportReasons: []gen.ReportReason{
			gen.Inappropriate,
			gen.Irrelevant,
		},
	}, nil
}
