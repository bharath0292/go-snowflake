package factory

import (
	"fmt"

	"github.com/go-snowflake/internal/factories/snowflake"
)

type Factory struct {
	SnowflakeDb *snowflake.SnowflakeClient
}

func NewFactory(sfConfig *snowflake.SnowflakeConfig) (*Factory, error) {

	snowflakeClient, err := snowflake.NewSnowflakeClient(sfConfig)
	if err != nil {
		return nil, fmt.Errorf("snowflake connection failed: %w", err)
	}

	return &Factory{
		SnowflakeDb: snowflakeClient,
	}, nil
}
