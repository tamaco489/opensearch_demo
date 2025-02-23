package usecase

import (
	"context"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

// IProductCommentUseCase は、商品コメントに関するユースケースを提供するインターフェースです。
type IProductCommentUseCase interface {
	// ユーザ向けAPI: GetProductCommentByID（商品に対しての詳細なコメントを取得します）
	GetProductCommentByID(ctx context.Context, request gen.GetProductCommentByIDRequestObject) (gen.GetProductCommentByIDResponseObject, error)

	// ユーザ向けAPI: CreateProductComment（商品に対して任意のコメントを投稿します）
	CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error)

	// 管理者向けAPI: GetProductCommentViolations（商品に対して投稿されたコメントの中で、予め定めたNGワードに該当するデータを取得します）
	GetProductCommentViolations(ctx context.Context, request gen.GetProductCommentViolationsRequestObject) (gen.GetProductCommentViolationsResponseObject, error)
}

// productCommentUseCase は、商品コメントに関連するユースケースを処理する具体的な実装です。
//
// # IProductCommentUseCase インターフェースを満たし、OpenSearch を使用して商品コメントの作成や
//
// NG ワードを含むコメントの取得などの機能を提供します。
type productCommentUseCase struct {
	// OpenSearch API クライアントを使用してデータの検索・登録を行います。
	opsApiClient *opensearchapi.Client
}

// NewCreateProductComment は、productCommentUseCase のコンストラクタ関数です。
//
// 新しい productCommentUseCase インスタンスを生成し、IProductCommentUseCase を実装したものを返します。
func NewCreateProductComment(opsApiClient *opensearchapi.Client) IProductCommentUseCase {
	return &productCommentUseCase{
		opsApiClient: opsApiClient,
	}
}
