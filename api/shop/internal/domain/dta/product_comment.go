package dta

import (
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/entity"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

// DTA: Data Transfer Adapter
// DTAはエンティティとDTO間でのデータ変換を担当します。
// 例えば、外部システム（API）から受け取ったデータをエンティティに変換したり、エンティティを外部システムに適した形に変換したりします。

// ToProductCommentEntity は OpenAPI のリクエストDTOをエンティティに変換します。
func ToProductCommentEntity(request gen.CreateProductCommentRequestObject, commentID, userID uint64) *entity.ProductComment {
	return entity.NewProductComment(
		commentID,
		request.ProductID,
		userID,
		request.Body.Title,
		request.Body.Content,
		request.Body.Rate,
	)
}

// ToProductCommentResponse はエンティティを API レスポンス形式に変換する
// func ToProductCommentResponse(comment *entity.ProductComment) gen.GetProductCommentByID200JSONResponse {
// 	return gen.GetProductCommentByID200JSONResponse{
// 		Id:            comment.ID,
// 		Title:         comment.Title,
// 		Content:       comment.Content,
// 		ReportReasons: []gen.ReportReason{}, // NOTE: 別途取得する場合は適宜修正
// 		User: gen.CommentByUser{
// 			UserId:    comment.UserID,
// 			UserName:  "氷織 羊", // NOTE: 別途RDS等で管理しているものを取得する
// 			AvatarUrl: fmt.Sprintf("https://example.com/users/%d/avatar.jpg", comment.UserID),
// 		},
// 		CreatedAt: comment.CreatedAt,
// 		Rate:      comment.Rate,
// 	}
// }
