package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetCreditCards(ctx *gin.Context, request gen.GetCreditCardsRequestObject) (gen.GetCreditCardsResponseObject, error) {

	// NOTE: 外部APIを実行する想定であるため、2000ms(2秒) 遅延させる。
	time.Sleep(2000 * time.Millisecond)

	creditCards := []gen.CreditCardList{
		{
			IsDefault:        true,
			MaskedCardNumber: "******123",
		},
		{
			IsDefault:        false,
			MaskedCardNumber: "******567",
		},
		{
			IsDefault:        false,
			MaskedCardNumber: "******890",
		},
	}

	return gen.GetCreditCards200JSONResponse(creditCards), nil
}
