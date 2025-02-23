package usecase

import (
	"context"
	"time"

	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (u productCommentUseCase) GetProductCommentByID(ctx context.Context, request gen.GetProductCommentByIDRequestObject) (gen.GetProductCommentByIDResponseObject, error) {

	// NOTE: 大量の商品データ、及びコメントデータを取得する想定とし、1500ms(1.5秒) 遅延させる。
	time.Sleep(1500 * time.Millisecond)

	// NOTE: OpenSearchによる商品情報の取得処理が実装されるまでは一旦固定値を返す。

	jst := time.FixedZone("JST", 9*60*60)

	return gen.GetProductCommentByID200JSONResponse{
		Id:        request.CommentID,
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
	}, nil
}
