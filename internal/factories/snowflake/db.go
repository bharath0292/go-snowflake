package snowflake

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-snowflake/pkg/utils"

	"github.com/rs/zerolog/log"
	"github.com/snowflakedb/gosnowflake"
)

type SnowflakeConfig struct {
	Account, User, Warehouse, Database, Schema, EncodedPrivateKey, Passphrase string
}

type SnowflakeClient struct {
	client *sql.DB
}

func createDSN(config *SnowflakeConfig) (*string, error) {
	sfConfig := gosnowflake.Config{
		Account:   config.Account,
		User:      config.User,
		Warehouse: config.Warehouse,
		Database:  config.Database,
		Schema:    config.Schema,
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(config.EncodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %s", err)
	}

	privateKey, err := utils.ParsePrivateKey(privateKeyBytes, []byte(config.Passphrase))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	sfConfig.Authenticator = gosnowflake.AuthTypeJwt
	sfConfig.PrivateKey = privateKey

	dsn, err := gosnowflake.DSN(&sfConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create DSN from Config: %v, err: %v", config, err)
	}

	return &dsn, nil
}

func NewSnowflakeClient(config *SnowflakeConfig) (*SnowflakeClient, error) {
	connString, err := createDSN(config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := sql.Open("snowflake", *connString)

	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)                  // max open connections
	sqlDB.SetMaxIdleConns(5)                   // max idle connections
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // recycle idle conns every 5 min
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // optional: recycle all conns every hour

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	log.Trace().Msg("Snowflake connection established")
	return &SnowflakeClient{client: sqlDB}, nil
}

func (s *SnowflakeClient) GetDB() *sql.DB {
	return s.client
}

func (s *SnowflakeClient) CloseDB() error {
	return s.client.Close()
}
