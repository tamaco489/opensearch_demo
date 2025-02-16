package usecase

import (
	"context"
	"log"

	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/configuration"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
	open_search "github.com/tamaco489/elasticsearch_demo/api/shop/internal/library/aws_open_search"
)

type CreateProductCommentUseCase struct {
	// db, redis等のインスタンスは外部から注入できるようにする
}

func NewCreateProductComment() *CreateProductCommentUseCase {
	return &CreateProductCommentUseCase{}
}

// CreateProductComment は商品に対してコメントを投稿します。
func (u *CreateProductCommentUseCase) CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	awsCfg := configuration.Get().AWSConfig
	client, err := open_search.NewOpenSearchAPIClient(awsCfg)
	if err != nil {
		return nil, err
	}

	res, err := client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("[info] result ping to open search", res)

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
