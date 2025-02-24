package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/configuration"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/gen"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/library/logger"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/library/open_search"
)

func NewHShopAPIServer() (*http.Server, error) {

	// CORSの設定
	corsCnf := NewCorsConfig()

	// Ginエンジンを初期化し、カスタムログの設定、CORSの設定を追加する
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logger.LogFormatter))
	r.Use(cors.New(corsCnf))

	// リクエスト処理中にpanicが発生した場合、それをキャッチしてサーバがcrashすることを防ぐ。
	//
	// 万が一panicが発生した場合でも、サーバは適切にエラーレスポンスを返し、正常に動作し続ける。※HTTP Status Code 5xx を返す
	r.Use(gin.Recovery())

	// NOTE: OpenSearchへの認証が必要な場合はこちら
	// client, err := open_search.NewOpenSearchAPIClientWithSigner(configuration.Get().AWSConfig)
	client, err := open_search.NewOpenSearchAPIClient()
	if err != nil {
		return nil, err
	}

	apiController := NewControllers(configuration.Get().API.Env, client)
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

// NewCorsConfig は CORS の設定を定義し、アプリケーションのリクエスト制御を行います。
//
// - AllowOrigins: すべてのオリジンを許可（一旦 "*" に設定）
//
// - AllowMethods: 許可する HTTP メソッドを指定
//
// - AllowHeaders: 許可するリクエストヘッダーを指定
//
// - ExposeHeaders: クライアント側でアクセス可能なレスポンスヘッダー
//
// - AllowCredentials: 認証情報の送信を許可（false に設定）
//
// - MaxAge: プリフライトリクエストのキャッシュ時間（秒）
func NewCorsConfig() cors.Config {
	return cors.Config{
		// 許可するオリジンを指定（一旦全許可）
		AllowOrigins: []string{"*"},

		// 必要なメソッドのみ許可
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"HEAD",
			"OPTIONS",
		},

		// 許可するヘッダーを限定
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
		},

		// クライアントがアクセスできるレスポンスヘッダー
		ExposeHeaders: []string{"Content-Length"},

		// 認証情報を送信可能にする
		AllowCredentials: false,

		// プリフライトリクエストのキャッシュ時間（秒）
		MaxAge: 86400,
	}
}
