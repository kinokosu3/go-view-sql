package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// PrintResponse 打印返回值
func PrintResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 只打印正确的返回
	if err := r.GetError(); err == nil {
		g.Log().Info(r.Context(), "[RESPONSE] <- status:", r.Response.Status)
		g.Log().Info(r.Context(), "[RESPONSE] <- response:", r.Response.BufferString())
	}
}
