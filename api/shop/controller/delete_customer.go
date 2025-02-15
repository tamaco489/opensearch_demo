package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) DeleteCustomer(ctx *gin.Context, request gen.DeleteCustomerRequestObject) (gen.DeleteCustomerResponseObject, error) {

	// NOTE: 外部APIを実行する想定であるため、300ms遅延させる。
	time.Sleep(300 * time.Millisecond)

	return gen.DeleteCustomer204Response{}, nil
}
