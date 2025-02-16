package usecase

import (
	"context"

	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

type ICreateProductComment interface {
	CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error)
}
