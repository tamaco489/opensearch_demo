package usecase

import (
	"context"
	"log"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

type createProductCommentUseCase struct {
	opsApiClient *opensearchapi.Client
}

func NewCreateProductComment(opsApiClient *opensearchapi.Client) *createProductCommentUseCase {
	return &createProductCommentUseCase{
		opsApiClient: opsApiClient,
	}
}

// CreateProductComment は商品に対してコメントを投稿します。
func (u *createProductCommentUseCase) CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	res, err := u.opsApiClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("[info] result ping to open search", res)

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
