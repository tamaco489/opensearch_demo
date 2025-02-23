package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/entity"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (u productCommentUseCase) GetProductCommentByID(ctx context.Context, request gen.GetProductCommentByIDRequestObject) (gen.GetProductCommentByIDResponseObject, error) {

	documentID := strconv.FormatUint(request.CommentID, 10)
	documentClient := u.opsApiClient.Document
	getResult, err := documentClient.Get(ctx, opensearchapi.DocumentGetReq{
		Index:      entity.ProductComments.String(),
		DocumentID: documentID,
	})
	if err != nil {
		return nil, err
	}

	var res opensearchapi.DocumentGetResp
	if err := json.NewDecoder(getResult.Inspect().Response.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	log.Println("[INFO] 成功 resp.ID:", res.ID)

	var source ProductCommentSource
	if err := json.Unmarshal(res.Source, &source); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source: %w", err)
	}

	log.Printf("[INFO] 成功 resp.Source: %+v", source)

	commentID, err := strconv.ParseUint(res.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05", source.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CreatedAt: %w", err)
	}

	return gen.GetProductCommentByID200JSONResponse{
		Id:            commentID,
		Title:         source.Title,
		Content:       source.Content,
		ReportReasons: []gen.ReportReason{}, // NOTE: 別途RDS等で管理しているものをuidで取得する
		User: gen.CommentByUser{
			UserId:    source.UserID,
			UserName:  "氷織 羊",                                                                // NOTE: 別途RDS等で管理しているものをuidで取得する
			AvatarUrl: fmt.Sprintf("https://example.com/users/%d/avatar.jpg", source.UserID), // NOTE: 別途RDS等で管理しているものをuidで取得する
		},
		CreatedAt: createdAt,
		Rate:      uint32(source.Rate),
	}, nil
}

// レスポンス用の構造体を定義
type ProductCommentSource struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Rate      int    `json:"rate"`
	CreatedAt string `json:"created_at"`
}
