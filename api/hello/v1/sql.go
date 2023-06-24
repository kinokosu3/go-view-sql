package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DimensionsReq struct {
	g.Meta `path:"/dimensions" tags:"sql" method:"post" summary:"get dimensions data"`
	Sql    string `json:"sql"`
}

type DimensionsRes struct {
	g.Meta `mime:"application/json"`
}

type SingleNumReq struct {
	g.Meta `path:"/single-num" tags:"sql" method:"post" summary:"get single data"`
	Sql    string `json:"sql"`
}
type SingleNumRes struct {
	g.Meta `mime:"application/json"`
}
type KeyValueReq struct {
	g.Meta `path:"/key-value" tags:"sql" method:"post" summary:"get key-value data"`
	Sql    string `json:"sql"`
}
type KeyValueRes struct {
	g.Meta `mime:"application/json"`
}
type ValueListReq struct {
	g.Meta `path:"/value-list" tags:"sql" method:"post" summary:"get value-list data"`
	Sql    string `json:"sql"`
}
type ValueListRes struct {
	g.Meta `mime:"application/json"`
}
