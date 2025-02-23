package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetProductMyComment(ctx *gin.Context, request gen.GetProductMyCommentRequestObject) (gen.GetProductMyCommentResponseObject, error) {

	jst := time.FixedZone("JST", 9*60*60)

	return gen.GetProductMyComment200JSONResponse{
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
