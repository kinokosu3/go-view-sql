package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// PrintRequest 打印请求参数
func PrintRequest(r *ghttp.Request) {
	g.Log().Info(r.Context(), "[REQUEST] -> path:", r.RequestURI)
	g.Log().Info(r.Context(), "[REQUEST] -> method:", r.Method)
	g.Log().Info(r.Context(), "[REQUEST] -> body:", r.GetRequestMap())

	r.Middleware.Next()
}
