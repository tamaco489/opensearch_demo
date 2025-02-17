package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetProductCommentViolations(ctx *gin.Context, request gen.GetProductCommentViolationsRequestObject) (gen.GetProductCommentViolationsResponseObject, error) {

	// NOTE: 大量の商品データ、及びコメントデータを取得する想定とし、1500ms(1.5秒) 遅延させる。
	time.Sleep(1500 * time.Millisecond)

	jst := time.FixedZone("JST", 9*60*60)

	ngComments := []gen.GetProductCommentViolations{
		{
			Id:        54009221,
			Title:     "とても良い商品です",
			Content:   "この商品は非常に良いです。特にデザインが素晴らしい。",
			CreatedAt: time.Date(2025, 2, 15, 13, 45, 30, 0, jst),
			LikeCount: 15,
			ReportReasons: []gen.ReportReason{
				gen.Inappropriate,
				gen.Irrelevant,
			},
			User: gen.CommentByUser{
				UserId:    12345,
				UserName:  "氷織 羊",
				AvatarUrl: "https://example.com/avatar.jpg",
			},
		},
		{
			Id:        54009226,
			Title:     "デザインが気に入った",
			Content:   "商品は思った通りのデザインでとても気に入りました。",
			CreatedAt: time.Date(2025, 2, 16, 14, 25, 30, 0, jst),
			LikeCount: 10,
			ReportReasons: []gen.ReportReason{
				gen.Fake,
			},
			User: gen.CommentByUser{
				UserId:    23456,
				UserName:  "御影 玲王",
				AvatarUrl: "https://example.com/avatar2.jpg",
			},
		},
		{
			Id:        54009533,
			Title:     "ちょっと期待外れ",
			Content:   "使用感は良かったが、価格に見合わないかもしれない。",
			CreatedAt: time.Date(2025, 2, 17, 15, 15, 30, 0, jst),
			LikeCount: 5,
			ReportReasons: []gen.ReportReason{
				gen.Other,
			},
			User: gen.CommentByUser{
				UserId:    34567,
				UserName:  "乙夜 影汰",
				AvatarUrl: "https://example.com/avatar3.jpg",
			},
		},
	}

	nextCursor := gen.GetProductCommentViolationsNextCursor{
		NextCursor: "NTQwMDk1MzY=",
	}

	response := gen.GetProductCommentViolations200JSONResponse{
		NgComments: ngComments,
		Metadata:   nextCursor,
	}

	return response, nil
}
