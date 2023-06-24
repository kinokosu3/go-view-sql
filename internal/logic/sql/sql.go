package sql

import (
	"context"
	"errors"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"go-view-sql/internal/model"
	"go-view-sql/internal/service"
)

type sSql struct{}

func newSSql() *sSql {
	return &sSql{}
}

func init() {
	service.RegisterSql(newSSql())
}

func parserSql(ctx context.Context, sql string) (fieldList []string, err error) {
	queryInterface, err := sqlparser.Parse(sql)
	if err != nil {
		g.Log().Errorf(ctx, "Dimensions Sql Parser Error: %v", err)
		return nil, err
	}
	var query *sqlparser.Select
	switch queryInterface.(type) {
	case *sqlparser.Select:
		query = queryInterface.(*sqlparser.Select)
	default:
		return nil, errors.New("不是select语句")
	}
	for _, v := range query.SelectExprs {
		var val *sqlparser.AliasedExpr
		switch v.(type) {
		case *sqlparser.AliasedExpr:
			val = v.(*sqlparser.AliasedExpr)
		default:
			return nil, errors.New("SQL fields Parser Error")
		}
		if val.As.String() == "" {
			switch val.Expr.(type) {
			case *sqlparser.ColName:
				fieldList = append(fieldList, val.Expr.(*sqlparser.ColName).Name.String())
			case *sqlparser.FuncExpr:
				fieldList = append(fieldList, val.Expr.(*sqlparser.FuncExpr).Name.String())
			default:
				return nil, errors.New("SQL fields Parser Error")
			}
		} else {
			fieldList = append(fieldList, val.As.String())
		}
	}
	return
}

// 确保顺序

func (s *sSql) Dimensions(ctx context.Context, in *model.DimensionsInput) (out *model.DimensionsOutput, err error) {

	fieldList, err := parserSql(ctx, in.Sql)
	if err != nil {
		return nil, err
	}
	list, err := g.DB().Ctx(ctx).GetAll(ctx, in.Sql)
	if err != nil {
		g.Log().Errorf(ctx, "Dimensions Error: %v", err)
		return nil, err
	}
	sourceList := make([]interface{}, 0)

	for _, v := range list {
		val := v.GMap()
		source := make(map[string]string, 0)
		for _, key := range fieldList {
			value := val.Get(key)
			source[key] = gconv.String(value)
		}
		sourceList = append(sourceList, source)
	}
	out = &model.DimensionsOutput{
		Dimensions: fieldList,
		Source:     sourceList,
	}
	return
}

func (s *sSql) SingleNum(ctx context.Context, in *model.SingleNumInput) (out interface{}, err error) {
	one, err := g.DB().Ctx(ctx).GetOne(ctx, in.Sql)
	if err != nil {
		g.Log().Errorf(ctx, "SingleNum Error: %v", err)
		return nil, err
	}
	dataMap := one.GMap()
	if len(dataMap.Keys()) < 1 {
		return nil, nil
	}
	out = dataMap.Get(dataMap.Keys()[0])
	return
}

func (s *sSql) KeyValue(ctx context.Context, in *model.KeyValueInput) (resList []interface{}, err error) {

	fieldList, err := parserSql(ctx, in.Sql)
	if err != nil {
		return nil, err
	}
	list, err := g.DB().Ctx(ctx).GetAll(ctx, in.Sql)
	if err != nil {
		g.Log().Errorf(ctx, "KeyValue Error: %v", err)
		return nil, err
	}
	for _, v := range list {
		val := v.GMap()
		res := make(map[string]string, 0)
		for _, key := range fieldList {
			value := val.Get(key)
			res[key] = gconv.String(value)
		}
		resList = append(resList, res)
	}
	return
}

func (s *sSql) ValueList(ctx context.Context, in *model.ValueListInput) (out [][]string, err error) {
	fieldList, err := parserSql(ctx, in.Sql)
	if err != nil {
		return nil, err
	}
	list, err := g.DB().Ctx(ctx).GetAll(ctx, in.Sql)
	if err != nil {
		g.Log().Errorf(ctx, "ValueList Error: %v", err)
		return nil, err
	}
	for _, v := range list {
		val := v.GMap()
		res := make([]string, 0)
		for _, key := range fieldList {
			value := val.Get(key)
			res = append(res, gconv.String(value))
		}
		out = append(out, res)
	}

	return
}
