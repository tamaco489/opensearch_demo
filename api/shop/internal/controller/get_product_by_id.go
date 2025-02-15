package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
)

func (c *Controllers) GetProductByID(ctx *gin.Context, request gen.GetProductByIDRequestObject) (gen.GetProductByIDResponseObject, error) {

	return gen.GetProductByID200JSONResponse{
		Id:              20001001,
		Name:            "プレミアムコーヒー",
		CategoryId:      10,
		CategoryName:    "飲料",
		Description:     "香り高いアラビカ種のコーヒーです。",
		Price:           500.0,
		DiscountFlag:    true,
		DiscountName:    "2024年クリスマスキャンペーン",
		DiscountRate:    20,
		DiscountedPrice: 400.0,
		StockQuantity:   50,
		VipOnly:         false,
		ImageUrl:        "https://example.com/images/20001001/product.jpg",
	}, nil
}
