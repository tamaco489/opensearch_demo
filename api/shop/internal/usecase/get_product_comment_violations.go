package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/entity"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/utils/ngwords"
)

// GetProductCommentViolations: 商品に対して投稿されたコメントの中で、予め定めたNGワードに該当するデータを取得します。
func (u productCommentUseCase) GetProductCommentViolations(ctx context.Context, request gen.GetProductCommentViolationsRequestObject) (gen.GetProductCommentViolationsResponseObject, error) {

	n := ngwords.NewNGWords()
	allNGWords := n.GetAllNGWordsCombined()

	// NGワードを含む検索クエリを組み立てる
	query := buildNGWordsQuery(allNGWords)

	// OpenSearchに検索リクエストを送信
	searchResult, err := u.opsApiClient.Search(
		ctx,
		&opensearchapi.SearchReq{
			Indices: []string{entity.ProductComments.String()},
			Body:    query,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to search comments with NG words: %v", err)
	}

	ngComments := make([]gen.GetProductCommentViolations, 0, len(searchResult.Hits.Hits))
	for _, hit := range searchResult.Hits.Hits {
		var comment gen.GetProductCommentViolations

		// 中間構造体でまずパースする
		var intermediate struct {
			Content   string `json:"content"`
			CreatedAt string `json:"created_at"`
			Id        uint64 `json:"id"`
			ProductId uint64 `json:"product_id"`
			Rate      uint32 `json:"rate"`
			// ReportReasons []ReportReason `json:"report_reasons"`
			Title  string `json:"title"`
			UserID uint64 `json:"user_id"`
			// User          CommentByUser  `json:"user"`
		}

		if err := json.Unmarshal(hit.Source, &intermediate); err != nil {
			return nil, fmt.Errorf("failed to unmarshal comment: %v", err)
		}

		// 日付文字列を time.Time に変換
		parsedTime, err := time.Parse("2006-01-02 15:04:05", intermediate.CreatedAt)
		if err != nil {
			parsedTime, err = time.Parse(time.RFC3339, intermediate.CreatedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to parse CreatedAt: %v", err)
			}
		}

		// intermediate から最終的な構造体に値をセット
		comment.Content = intermediate.Content
		comment.CreatedAt = parsedTime
		comment.Id = intermediate.Id
		comment.ProductId = intermediate.ProductId
		comment.Rate = intermediate.Rate
		comment.ReportReasons = []gen.ReportReason{} // NOTE: 別途RDS等で管理しているものをuidで取得する
		comment.Title = intermediate.Title

		// user構造体
		comment.User.UserId = intermediate.UserID
		comment.User.UserName = ""                                                                           // NOTE: 別途RDS等で管理しているものをuidで取得する
		comment.User.AvatarUrl = fmt.Sprintf("https://example.com/users/%d/avatar.jpg", intermediate.UserID) // NOTE: 別途RDS等で管理しているものをuidで取得する

		// NGコメントを追加
		ngComments = append(ngComments, comment)
	}

	// NOTE: cursorの値は一旦固定値とする
	nextCursor := gen.GetProductCommentViolationsNextCursor{
		NextCursor: "NTQwMDk1MzY=",
	}

	response := gen.GetProductCommentViolations200JSONResponse{
		NgComments: ngComments,
		Metadata:   nextCursor,
	}

	return response, nil
}

func buildNGWordsQuery(ngWords []string) *strings.Reader {
	var shouldClauses []string
	for _, word := range ngWords {
		shouldClauses = append(shouldClauses, fmt.Sprintf(`{ "match_phrase": { "content": "%s" } }`, word))
		shouldClauses = append(shouldClauses, fmt.Sprintf(`{ "match_phrase": { "title": "%s" } }`, word))
	}

	queryString := fmt.Sprintf(`{
		"query": {
			"bool": {
				"should": [%s],
				"minimum_should_match": 1
			}
		},
		"size": 10
	}`, strings.Join(shouldClauses, ","))

	return strings.NewReader(queryString)
}
