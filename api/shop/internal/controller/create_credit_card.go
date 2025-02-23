package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) CreateCreditCard(ctx *gin.Context, request gen.CreateCreditCardRequestObject) (gen.CreateCreditCardResponseObject, error) {

	// NOTE: 外部APIを実行する想定であるため、500ms遅延させる。※非同期的に即時で204返却しても良いが負荷テストのシナリオ的に用意
	time.Sleep(500 * time.Millisecond)

	return gen.CreateCreditCard204Response{}, nil
}
