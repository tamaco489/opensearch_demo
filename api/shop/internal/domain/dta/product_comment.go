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
