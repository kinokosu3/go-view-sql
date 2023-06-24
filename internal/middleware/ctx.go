package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/text/gstr"
)

// CtxRewrite 将 Request-Id 转成 Trance-Id
func CtxRewrite(r *ghttp.Request) {
	requestId := gstr.Replace(r.GetHeader("Request-Id"), "-", "")
	if requestId == "" {
		g.Log().Info(r.Context(), "[HEADER] Request-Id is none, trace-id will be generated automatically")
	} else {
		g.Log().Infof(r.Context(), "[HEADER] Request-Id is %s, trace-id is already set to it", requestId)
	}
	if requestId != "" {
		newCtx, err := gtrace.WithTraceID(r.Context(), requestId)
		if err != nil {
			g.Log().Errorf(r.Context(), "request id invalid: ", err.Error())
		} else {
			r.SetCtx(newCtx)
		}
	}

	r.Middleware.Next()
}
