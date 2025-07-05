package entities

import "github.com/go-snowflake/pkg/enums"

type ColumnInfo struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	DataType enums.DataType `json:"dataType"`
}
