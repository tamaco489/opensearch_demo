package usecase

import (
	"context"

	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

type CreateProductCommentUseCase struct {
	// db, redis等のインスタンスは外部から注入できるようにする
}

func NewCreateProductComment() *CreateProductCommentUseCase {
	return &CreateProductCommentUseCase{}
}

// CreateProductComment は商品に対してコメントを投稿します。
func (u *CreateProductCommentUseCase) CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
