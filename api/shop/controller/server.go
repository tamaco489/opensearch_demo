package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/gen"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/library/logger"
)

func NewHShopAPIServer() (*http.Server, error) {

	// CORSの設定
	cnf := cors.DefaultConfig()
	cnf.AllowOrigins = []string{"*"}
	cnf.AllowHeaders = append(cnf.AllowHeaders, "Authorization", "Access-Control-Allow-Origin")

	// Ginエンジンを初期化し、カスタムログの設定、CORSの設定を追加する
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logger.LogFormatter))
	r.Use(cors.New(cnf))

	// リクエスト処理中にpanicが発生した場合、それをキャッチしてサーバがcrashすることを防ぐ。
	//
	// 万が一panicが発生した場合でも、サーバは適切にエラーレスポンスを返し、正常に動作し続ける。※HTTP Status Code 5xx を返す
	r.Use(gin.Recovery())

	// NOTE: envは一旦固定値で渡す
	env := "dev"
	apiController := NewControllers(env)
	strictServer := gen.NewStrictHandler(apiController, nil)

	gen.RegisterHandlersWithOptions(
		r,
		strictServer,
		gen.GinServerOptions{
			BaseURL:     "/shop/",
			Middlewares: []gen.MiddlewareFunc{},
			ErrorHandler: func(ctx *gin.Context, err error, i int) {
				_ = ctx.Error(err)
				ctx.JSON(i, gin.H{"msg": err.Error()})
			},
		},
	)

	// NOTE: portは一旦固定値で渡す
	port := "8080"
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", port),
	}

	return server, nil
}
