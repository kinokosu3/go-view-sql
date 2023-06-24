// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-view-sql/internal/model"
)

type (
	ISql interface {
		Dimensions(ctx context.Context, in *model.DimensionsInput) (out *model.DimensionsOutput, err error)
		SingleNum(ctx context.Context, in *model.SingleNumInput) (out interface{}, err error)
		KeyValue(ctx context.Context, in *model.KeyValueInput) (resList []interface{}, err error)
		ValueList(ctx context.Context, in *model.ValueListInput) (out [][]string, err error)
	}
)

var (
	localSql ISql
)

func Sql() ISql {
	if localSql == nil {
		panic("implement not found for interface ISql, forgot register?")
	}
	return localSql
}

func RegisterSql(i ISql) {
	localSql = i
}
