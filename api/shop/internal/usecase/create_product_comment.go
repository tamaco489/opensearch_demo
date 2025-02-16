package usecase

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

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

	// 本来はctxから取得したsubなどでuser_idを特定する
	var userID uint64 = 25540992
	log.Println("[INFO] product_id:", request.ProductID, "user_id:", userID)

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

	// hitしなかった場合は指定された商品に対してコメントが1つもないと見なし、コメントIDに1を指定
	if len(searchResult.Hits.Hits) == 0 {
		commentID = 1
	}
	if len(searchResult.Hits.Hits) > 0 {
		commentIDStr := searchResult.Hits.Hits[0].ID
		commentID, err = strconv.ParseUint(commentIDStr, 10, 64)
		if err != nil {
			return gen.CreateCharge500Response{}, fmt.Errorf("failed to convert comment ID to uint64: %v", err)
		}
	}

	log.Println("[INFO] comment_id:", commentID)

	return gen.CreateProductComment201JSONResponse{
		Id: 70235591,
	}, nil
}
