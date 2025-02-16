package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/domain/entity"
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
//
// NOTE: 2回目のコメントが直前のコメントを上書きしてしまうため原因を調査し、修正する。
func (u *createProductCommentUseCase) CreateProductComment(ctx context.Context, request gen.CreateProductCommentRequestObject) (gen.CreateProductCommentResponseObject, error) {

	// 本来はctxから取得したsubなどでuser_idを特定する
	var userID uint64 = 25540992

	// 検索クエリをJSON文字列として構築
	// query := fmt.Sprintf(`{"query": {"match": {"product_id": %d}}, "sort": [{"created_at": {"order": "desc"}}], "size": 1}`, request.ProductID)
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
	// トランザクションなどを考慮すると、コメントデータはRDBMS等で管理し、そこからデータを取得したほうが保守的な観点では良さそう。
	searchResult, err := u.opsApiClient.Search(
		ctx,
		&opensearchapi.SearchReq{
			Indices: []string{"product_comments"},
			Body:    query,
		},
	)
	if err != nil {
		return gen.CreateCharge500Response{}, err
	}

	// 検索結果から直近のコメントIDを抽出
	var commentID uint64
	var commentIDStr string

	// hitしなかった場合は指定された商品に対してコメントが1つもないと見なし、コメントIDに1を指定
	if searchResult.Hits.Total.Value == 0 {
		commentID = 1
	}
	if searchResult.Hits.Total.Value > 0 {
		commentIDStr = searchResult.Hits.Hits[0].ID
		commentID, err = strconv.ParseUint(commentIDStr, 10, 64)
		if err != nil {
			return gen.CreateCharge500Response{}, fmt.Errorf("failed to convert comment ID to uint64: %v", err)
		}
	}

	newComment := entity.NewProductComment(
		commentID, // 現時点で最も新しいコメントIDを指定
		request.ProductID,
		userID,
		request.Body.Title,
		request.Body.Content,
		request.Body.Rate,
	)

	// JSON に変換
	commentJSON, err := json.Marshal(newComment)
	if err != nil {
		return gen.CreateProductComment500Response{}, fmt.Errorf("failed to marshal new comment: %v", err)
	}

	idxRequest := opensearchapi.IndexReq{
		Index:      "product_comments",
		DocumentID: strconv.FormatUint(newComment.ID, 10),
		Body:       bytes.NewReader(commentJSON),
		Params: opensearchapi.IndexParams{
			Refresh: "true",
		},
	}
	_, err = u.opsApiClient.Index(ctx, idxRequest)
	if err != nil {
		return gen.CreateProductComment500Response{}, fmt.Errorf("error creating product comment: %v", err)
	}

	return gen.CreateProductComment201JSONResponse{
		Id: newComment.ID,
	}, nil
}
