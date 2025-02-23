package usecase

import (
	"context"

	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (u productCommentUseCase) DeleteProductCommentByID(ctx context.Context, request gen.DeleteProductCommentByIDRequestObject) (gen.DeleteProductCommentByIDResponseObject, error) {

	return gen.DeleteProductCommentByID204Response{}, nil
}
