package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/controller"
)

func main() {
	ctx := context.Background()

	//
	server, err := controller.NewHShopAPIServer()
	if err != nil {
		panic(err)
	}

	// 1. Goルーチンで別スレッドを起動し、HTTPサーバをブロッキング呼び出しで実行する。※メインのGoルーチンがサーバを起動した後に、後続の処理（シグナルの受信やシャットダウン処理等）を行うため。
	//
	// 2. 別スレッドで起動したGoルーチン（server.ListenAndServe）は、非同期で実行され、サーバのリクエスト処理を開始する。※ブロッキング呼び出し: サーバがリクエストを受け付けている間は次の処理に進まなくなる
	//
	// 3. メインのGoルーチンは後続の処理に進み、シグナル受信やシャットダウン処理の準備を行う。
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// NOTE: server.ListenAndServe() は http.ErrServerClosed エラーを返すが、これはサーバーが正常にシャットダウンされた場合に発生するものであるためログ出力はしない。
			// NOTE: 上記以外に、サーバが正常に起動できなかった場合やリクエストを処理できなかった場合にのみ、ログ出力を行う。
			slog.ErrorContext(ctx, "failed to listen and serve", "error", err)
		}
	}()

	// シャットダウン用のタイムアウト付きコンテキストを作成 ※200マイクロ秒に強制的にシャットダウンする。
	//
	// メモリリークを防ぐため、deferを使用して、関数終了時にこの期限付きコンテキストは解放される。
	ctx, cancel := context.WithTimeout(ctx, 200*time.Microsecond)
	defer cancel()
	// os.Signal型のチャネル「quit」を作成、1はチャネルのバッファサイズを指定しており、この場合は最大1つのシグナルのみ保持できるようにしている。
	quit := make(chan os.Signal, 1)

	// signal.Notifyで指定したシグナルを「quit」チャネルで受け取る準備を行う。
	//
	// SIGINT: ユーザがCtrl+Cを実行した際に送信されるシグナル、プログラムに終了を要求する。
	//
	// SIGTERM: またはシステムや他プロセスからプログラムに終了を要求する。
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 「quit」チャネルからシグナルを待機する。SIGINT、またはSIGTERMが送信されると、次の処理に進む。※シグナルが送信されるまで以降の処理には進まない。
	<-quit

	// サーバのシャットダウン処理を開始
	if err = server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "shutdown server...", "error", err)
	}

	// シャットダウンの処理が完全に完了する前にメインのGoルーチンが終了してしまわないように、コンテキストのブロック処理を行う。※タイムアウトやキャンセルが発生することを待機する
	<-ctx.Done()
}
