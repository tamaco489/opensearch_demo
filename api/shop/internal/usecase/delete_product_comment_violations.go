package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/domain/entity"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (u productCommentUseCase) DeleteProductCommentViolationsByID(ctx context.Context, request gen.DeleteProductCommentViolationByIDRequestObject) (gen.DeleteProductCommentViolationByIDResponseObject, error) {

	documentID := strconv.FormatUint(request.CommentID, 10)

	// NOTE: https://github.com/opensearch-project/opensearch-go/blob/main/opensearchapi/api_document-delete.go
	documentClient := u.opsApiClient.Document
	deleteResult, err := documentClient.Delete(
		ctx,
		opensearchapi.DocumentDeleteReq{
			Index:      entity.ProductComments.String(),
			DocumentID: documentID,
		},
	)
	if err != nil {
		if deleteResult.Inspect().Response.StatusCode == http.StatusNotFound {
			slog.ErrorContext(ctx, fmt.Sprintf("not found comment id: %v", err))
			return gen.DeleteProductCommentViolationByID404Response{}, nil
		}
		return gen.DeleteProductCommentViolationByID500Response{}, err
	}

	return gen.DeleteProductCommentViolationByID204Response{}, nil
}
