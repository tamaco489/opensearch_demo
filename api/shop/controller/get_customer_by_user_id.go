package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetCustomerByUserID(ctx *gin.Context, request gen.GetCustomerByUserIDRequestObject) (gen.GetCustomerByUserIDResponseObject, error) {

	// NOTE: 外部APIを実行する想定であるため、400ms遅延させる。
	time.Sleep(400 * time.Millisecond)

	return gen.GetCustomerByUserID200JSONResponse{
		Id: "xyz12345",
	}, nil
}
