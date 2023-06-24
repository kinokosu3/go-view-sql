package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "go-view-sql/api/hello/v1"
	"go-view-sql/internal/model"
	"go-view-sql/internal/service"
)

var (
	Sql = cSql{}
)

type cSql struct{}

func (c *cSql) Dimensions(ctx context.Context, req *v1.DimensionsReq) (res *v1.DimensionsRes, err error) {
	out, err := service.Sql().Dimensions(ctx, &model.DimensionsInput{
		Sql: req.Sql,
	})
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: "ok",
		Data:    out,
	})
	return
}

func (c *cSql) SingleNum(ctx context.Context, req *v1.SingleNumReq) (res *v1.SingleNumRes, err error) {
	out, err := service.Sql().SingleNum(ctx, &model.SingleNumInput{
		Sql: req.Sql,
	})
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: "ok",
		Data:    out,
	})
	return

}
func (c *cSql) KeyValue(ctx context.Context, req *v1.KeyValueReq) (res *v1.SingleNumRes, err error) {
	out, err := service.Sql().KeyValue(ctx, &model.KeyValueInput{
		Sql: req.Sql,
	})
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: "ok",
		Data:    out,
	})
	return

}
func (c *cSql) ValueList(ctx context.Context, req *v1.ValueListReq) (res *v1.ValueListRes, err error) {
	out, err := service.Sql().ValueList(ctx, &model.ValueListInput{
		Sql: req.Sql,
	})
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: "ok",
		Data:    out,
	})
	return

}
