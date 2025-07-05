package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-snowflake/internal/entities"
	"github.com/go-snowflake/internal/factories/snowflake"
	"github.com/go-snowflake/pkg/enums"
)

type ColumnListRepository struct {
	db *snowflake.SnowflakeClient
}

func NewColumnListRepository(db *snowflake.SnowflakeClient) *ColumnListRepository {
	return &ColumnListRepository{
		db: db,
	}
}

func getColumnDataType(colType string) enums.DataType {
	colType = strings.ToUpper(colType)

	switch {
	case strings.HasPrefix(colType, "VARCHAR"):
		return enums.String
	case strings.HasPrefix(colType, "CHAR"):
		return enums.String
	case strings.HasPrefix(colType, "TEXT"):
		return enums.String
	case strings.HasPrefix(colType, "FLOAT"):
		return enums.Number
	case strings.HasPrefix(colType, "DOUBLE"):
		return enums.Number
	case strings.HasPrefix(colType, "INT"):
		return enums.Number
	case strings.HasPrefix(colType, "DECIMAL"):
		return enums.Number
	case strings.HasPrefix(colType, "BOOL"):
		return enums.Boolean
	case strings.HasPrefix(colType, "DATE"):
		return enums.Date
	case strings.HasPrefix(colType, "TIMESTAMP"):
		return enums.Date
	default:
		return enums.String
	}
}

func (r *ColumnListRepository) FetchAllColumns(ctx context.Context, tableName string) ([]*entities.ColumnInfo, error) {
	query := fmt.Sprintf(`DESCRIBE TABLE %s`, tableName)

	db := r.db.GetDB()
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	columnList := []*entities.ColumnInfo{}

	for rows.Next() {
		var colName, colType string

		values := make([]any, len(columns))
		values[0] = &colName
		values[1] = &colType
		for i := 2; i < len(columns); i++ {
			var v any
			values[i] = &v
		}
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		columnList = append(columnList, &entities.ColumnInfo{
			ID:       colName,
			Name:     colName,
			DataType: getColumnDataType(colType),
		})

	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return columnList, nil
}
