package usecase

import (
	"context"

	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

type IProductCommentUseCase interface {
	CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error)
}

var (
	_ (IProductCommentUseCase) = (*CreateProductCommentUseCase)(nil)
)
