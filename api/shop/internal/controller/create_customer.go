package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) CreateCustomer(ctx *gin.Context, request gen.CreateCustomerRequestObject) (gen.CreateCustomerResponseObject, error) {

	// NOTE: 外部APIを実行する想定であるため、150ms遅延させる。
	time.Sleep(150 * time.Millisecond)

	return gen.CreateCustomer201JSONResponse{
		Id: "xyz12345",
	}, nil
}
