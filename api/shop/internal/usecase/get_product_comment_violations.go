package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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

	ngComments := make([]gen.GetProductCommentViolations, len(searchResult.Hits.Hits))
	for i, hit := range searchResult.Hits.Hits {

		// 中間構造体でまずパースする
		var intermediate struct {
			Id        uint64 `json:"id"`
			Title     string `json:"title"`
			Content   string `json:"content"`
			CreatedAt string `json:"created_at"`
			ProductId uint64 `json:"product_id"`
			Rate      uint32 `json:"rate"`
			UserID    uint64 `json:"user_id"`
		}

		if err := json.Unmarshal(hit.Source, &intermediate); err != nil {
			return gen.GetProductCommentViolations500Response{}, fmt.Errorf("failed to unmarshal comment: %v", err)
		}

		commentID, err := strconv.ParseUint(hit.ID, 10, 64)
		if err != nil {
			return gen.GetProductCommentViolations500Response{}, fmt.Errorf("failed to parse product comment ID: %v", err)
		}

		// 日付文字列を time.Time に変換
		parsedTime, err := time.Parse("2006-01-02 15:04:05", intermediate.CreatedAt)
		if err != nil {
			parsedTime, err = time.Parse(time.RFC3339, intermediate.CreatedAt)
			if err != nil {
				return gen.GetProductCommentViolations500Response{}, fmt.Errorf("failed to parse CreatedAt: %v", err)
			}
		}

		ngComments[i] = gen.GetProductCommentViolations{
			Id:            commentID,
			Content:       intermediate.Content,
			Title:         intermediate.Title,
			CreatedAt:     parsedTime,
			ProductId:     intermediate.ProductId,
			ReportReasons: []gen.ReportReason{}, // NOTE: 別途RDS等で管理しているものをuidで取得する
			Rate:          intermediate.Rate,
			User: gen.CommentByUser{
				UserId:    intermediate.UserID,
				UserName:  "",                                                                          // NOTE: 別途RDS等で管理しているものをuidで取得する
				AvatarUrl: fmt.Sprintf("https://example.com/users/%d/avatar.jpg", intermediate.UserID), // NOTE: 別途RDS等で管理しているものをuidで取得する
			},
		}
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
