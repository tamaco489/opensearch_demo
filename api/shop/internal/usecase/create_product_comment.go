package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/dta"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/entity"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

// CreateProductComment は商品に対して任意のコメントを投稿します。
func (u *productCommentUseCase) CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	// 本来はctxから取得したsubなどでuser_idを特定する
	var userID uint64 = 25540992

	// 検索クエリをJSON文字列として構築
	query := strings.NewReader(`{
		"query": {
			"match_all": {}
		},
		"sort": [
			{ "created_at": { "order": "desc" } }
		],
		"size": 1
	}`)

	// 検索リクエストを作成、直近のコメントIDを取得する
	// NOTE: トランザクションなどを考慮すると、コメントデータはRDBMS等で管理し、そこからデータを取得したほうが保守的な観点では良さそう。
	searchResult, err := u.opsApiClient.Search(
		ctx,
		&opensearchapi.SearchReq{
			Indices: []string{entity.ProductComments.String()},
			Body:    query,
		},
	)
	if err != nil {
		return gen.CreateCharge500Response{}, err
	}

	// 初期値の設定、検索結果が0件ではない場合にのみ値の更新を行う
	commentID := uint64(1)
	if searchResult.Hits.Total.Value > 0 {
		commentID, err = strconv.ParseUint(searchResult.Hits.Hits[0].ID, 10, 64)
		if err != nil {
			return gen.CreateCharge500Response{}, fmt.Errorf("failed to convert comment ID to uint64: %v", err)
		}
	}

	// Entityを生成し、jsonに変換
	commentEntity := dta.ToProductCommentEntity(request, commentID, userID)
	commentEntityJSON, err := json.Marshal(commentEntity)
	if err != nil {
		return gen.CreateProductComment500Response{}, fmt.Errorf("failed to marshal new comment: %v", err)
	}

	// OpenSearch にデータ投入
	idxRequest := opensearchapi.IndexReq{
		Index:      entity.ProductComments.String(),
		DocumentID: strconv.FormatUint(commentEntity.ID, 10),
		Body:       bytes.NewReader(commentEntityJSON),
		Params: opensearchapi.IndexParams{
			Refresh: "true",
			Timeout: 5 * time.Second,
		},
	}
	_, err = u.opsApiClient.Index(ctx, idxRequest)
	if err != nil {
		return gen.CreateProductComment500Response{}, fmt.Errorf("error creating product comment: %v", err)
	}

	return gen.CreateProductComment201JSONResponse{
		Id: commentEntity.ID,
	}, nil
}
