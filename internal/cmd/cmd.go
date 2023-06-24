package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	controller "go-view-sql/internal/controller/sql"
	"go-view-sql/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse,
					middleware.CtxRewrite,
					middleware.ContentType,
					middleware.PrintRequest,
					middleware.PrintResponse,
					middleware.MiddlewareCORS,
				)
				group.Bind(
					controller.Sql,
				)
			})
			s.Run()
			return nil
		},
	}
)
