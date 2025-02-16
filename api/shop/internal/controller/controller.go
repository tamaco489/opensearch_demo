package controller

import (
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/usecase"
)

type Controllers struct {
	env                   string
	productCommentUseCase usecase.CreateProductCommentUseCase
}

func NewControllers(
	env string,
	opensearchApiClient *opensearchapi.Client,
) *Controllers {
	return &Controllers{
		env:                   env,
		productCommentUseCase: *usecase.NewCreateProductComment(opensearchApiClient),
	}
}
