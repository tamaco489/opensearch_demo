package controller

import (
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/usecase"
)

type Controllers struct {
	env                   string
	productCommentUseCase usecase.IProductCommentUseCase
}

func NewControllers(env string, opensearchApiClient *opensearchapi.Client) *Controllers {
	productCommentUseCase := usecase.NewCreateProductComment(opensearchApiClient)
	return &Controllers{
		env:                   env,
		productCommentUseCase: productCommentUseCase,
	}
}
