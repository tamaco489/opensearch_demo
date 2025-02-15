package logger

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogFormatter: HTTP Status Code に応じて出力するログレベルを切り分けます。
func LogFormatter(params gin.LogFormatterParams) string {
	switch {
	// HTTP Status Code: 500 -> エラーレベルのログを出力
	case http.StatusInternalServerError <= params.StatusCode:
		slog.ErrorContext(params.Request.Context(), params.ErrorMessage,
			"status", params.StatusCode,
			"method", params.Method,
			"path", params.Path,
			"ip", params.ClientIP,
			"latency_ms", params.Latency.Milliseconds(),
			"user_agent", params.Request.UserAgent(),
			"host", params.Request.Host,
		)
		return ""

	// HTTP Status Code: 400以上、500未満 -> 警告レベルのログを出力
	case params.StatusCode >= http.StatusBadRequest && params.StatusCode <= http.StatusInternalServerError:
		slog.WarnContext(params.Request.Context(), params.ErrorMessage,
			"status", params.StatusCode,
			"method", params.Method,
			"path", params.Path,
			"ip", params.ClientIP,
			"latency_ms", params.Latency.Milliseconds(),
			"user_agent", params.Request.UserAgent(),
			"host", params.Request.Host,
		)
		return ""

	// エラーが発生しなかった場合は通常レベルのログを出力
	default:
		slog.InfoContext(params.Request.Context(), "access",
			"status", params.StatusCode,
			"method", params.Method,
			"path", params.Path,
			"ip", params.ClientIP,
			"latency_ms", params.Latency.Milliseconds(),
			"user_agent", params.Request.UserAgent(),
			"host", params.Request.Host,
		)
		return ""
	}
}
