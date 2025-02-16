package usecase

import (
	"context"
	"log"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

type CreateProductCommentUseCase struct {
	opensearchApiClient *opensearchapi.Client
}

func NewCreateProductComment(
	opensearchApiClient *opensearchapi.Client,
) *CreateProductCommentUseCase {
	return &CreateProductCommentUseCase{
		opensearchApiClient,
	}
}

// CreateProductComment は商品に対してコメントを投稿します。
func (u *CreateProductCommentUseCase) CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	res, err := u.opensearchApiClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("[info] result ping to open search", res)

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
