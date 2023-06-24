package model

type DimensionsInput struct {
	Sql string `json:"sql"`
}

type DimensionsOutput struct {
	Dimensions []string      `json:"dimensions"`
	Source     []interface{} `json:"source"`
}

type SingleNumInput struct {
	Sql string `json:"sql"`
}

type KeyValueInput struct {
	Sql string `json:"sql"`
}

type ValueListInput struct {
	Sql string `json:"sql"`
}
