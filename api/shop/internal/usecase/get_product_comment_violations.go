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
	query := u.buildNGWordsQuery(allNGWords)

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

	// データ変換のメソッドを呼び出して詰め替え
	ngComments, err := u.transformToGetProductCommentViolations(searchResult.Hits.Hits)
	if err != nil {
		return gen.GetProductCommentViolations500Response{}, fmt.Errorf("failed to transform search hits: %v", err)
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

// buildNGWordsQuery: 与えられたNGワードを使ってOpenSearchの検索クエリを構築する関数。
// それぞれのNGワードに対して、商品コメントの「content」と「title」フィールドを検索し、
// 該当するものを取得するためのboolクエリを組み立てます。
//
// Args:
//
//	ngWords []string: NGワードのリスト。これらのワードを商品コメントの「content」および「title」に対して検索します。
//
// Returns:
//
//	*strings.Reader: 組み立てた検索クエリを格納したReaderオブジェクト。これをOpenSearchの検索リクエストのボディとして使用します。
func (u productCommentUseCase) buildNGWordsQuery(ngWords []string) *strings.Reader {
	// "should"句に含める検索条件を格納するスライス
	var shouldClauses []string

	// 各NGワードについて、contentとtitleのフィールドに対するmatch_phrase検索を追加
	for _, word := range ngWords {
		shouldClauses = append(shouldClauses, fmt.Sprintf(`{ "match_phrase": { "content": "%s" } }`, word))
		shouldClauses = append(shouldClauses, fmt.Sprintf(`{ "match_phrase": { "title": "%s" } }`, word))
	}

	// 検索クエリをJSON形式の文字列として組み立て
	queryString := fmt.Sprintf(`{
		"query": {
			"bool": {
				"should": [%s],                      // "should"句にNGワードの条件を追加
				"minimum_should_match": 1            // 1つ以上の条件に一致すればマッチとみなす
			}
		},
		"size": 10                                  // 最大で10件の結果を返す
	}`, strings.Join(shouldClauses, ","))

	// 構築したクエリ文字列をstrings.Readerに変換して返す
	return strings.NewReader(queryString)
}

// transformToGetProductCommentViolations: OpenSearchの検索結果から、商品コメントの情報をgen.GetProductCommentViolations型に変換するメソッドです。
//
// OpenSearchで管理しているcreated_atはstring型だが、API Responseで定義している型がtime.Time型であり、
//
// そのままUnmarshalしてしまうとエラーになってしまうため、一度中間構造体へparseし、その上でcreated_atをtime.Time型に変換してAPI Responseに設定する。
func (u productCommentUseCase) transformToGetProductCommentViolations(hits []opensearchapi.SearchHit) ([]gen.GetProductCommentViolations, error) {

	ngComments := make([]gen.GetProductCommentViolations, len(hits))

	for i, hit := range hits {
		// OpenSearchで管理しているcreated_atはstring型だが、API Responseで定義している型がtime.Time型であり、
		// そのままUnmarshalしてしまうとエラーになってしまうため、一度中間構造体へparseし、その上でcreated_atをtime.Time型に変換して
		// API Responseに設定する。

		var intermediate struct {
			Id        uint64 `json:"id"`
			Title     string `json:"title"`
			Content   string `json:"content"`
			CreatedAt string `json:"created_at"` // コメント作成日時（文字列形式）
			ProductId uint64 `json:"product_id"`
			Rate      uint32 `json:"rate"`
			UserID    uint64 `json:"user_id"`
		}

		if err := json.Unmarshal(hit.Source, &intermediate); err != nil {
			return nil, fmt.Errorf("failed to unmarshal comment: %v", err)
		}

		// コメントIDを文字列からuint64型に変換
		commentID, err := strconv.ParseUint(hit.ID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse product comment ID: %v", err)
		}

		// 日付文字列をtime.Time型に変換
		parsedTime, err := time.Parse("2006-01-02 15:04:05", intermediate.CreatedAt)
		if err != nil {
			parsedTime, err = time.Parse(time.RFC3339, intermediate.CreatedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to parse CreatedAt: %v", err)
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

	return ngComments, nil
}
