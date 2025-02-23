package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/entity"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (u productCommentUseCase) GetProductCommentByID(ctx context.Context, request gen.GetProductCommentByIDRequestObject) (gen.GetProductCommentByIDResponseObject, error) {

	documentID := strconv.FormatUint(request.CommentID, 10)

	// NOTE: https: //github.com/opensearch-project/opensearch-go/blob/main/opensearchapi/api_document-get.go
	documentClient := u.opsApiClient.Document
	getResult, err := documentClient.Get(ctx, opensearchapi.DocumentGetReq{
		Index:      entity.ProductComments.String(),
		DocumentID: documentID,
	})
	if err != nil {
		return gen.GetProductCommentByID500Response{}, fmt.Errorf("failed to get product comment by id: %v", err)
	}

	var res opensearchapi.DocumentGetResp
	if err := json.NewDecoder(getResult.Inspect().Response.Body).Decode(&res); err != nil {
		return gen.GetProductCommentByID500Response{}, fmt.Errorf("failed to decode response: %w", err)
	}

	var source entity.ProductComment
	if err := json.Unmarshal(res.Source, &source); err != nil {
		return gen.GetProductCommentByID500Response{}, fmt.Errorf("failed to unmarshal source: %w", err)
	}

	// comment_id を string → uint64に変換
	commentID, err := strconv.ParseUint(res.ID, 10, 64)
	if err != nil {
		return gen.GetProductCommentByID500Response{}, fmt.Errorf("failed to parse comment_id: %w", err)
	}

	// created_at をtime.Time型に変換
	createdAt, err := time.Parse("2006-01-02 15:04:05", source.CreatedAt)
	if err != nil {
		return gen.GetProductCommentByID500Response{}, fmt.Errorf("failed to parse created_at: %w", err)
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
		ProductId: source.ProductID,
		Rate:      uint32(source.Rate),
	}, nil
}
