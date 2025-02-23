package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) CreateCharge(ctx *gin.Context, request gen.CreateChargeRequestObject) (gen.CreateChargeResponseObject, error) {

	// NOTE: 外部APIを実行する想定であるため、1000ms遅延させる。
	time.Sleep(1000 * time.Millisecond)

	return gen.CreateCharge204Response{}, nil
}
