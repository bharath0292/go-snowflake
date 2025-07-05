package main

import (
	"os"

	"github.com/gin-gonic/gin"
	factory "github.com/go-snowflake/internal/factories"
	"github.com/go-snowflake/internal/factories/snowflake"
	repository "github.com/go-snowflake/internal/repositories"
	"github.com/go-snowflake/internal/services"
	"github.com/go-snowflake/pkg/utils"
	handler "github.com/go-snowflake/server/handlers"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/snowflakedb/gosnowflake"
)

func init() {
	utils.InitLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading.env file: %v", err)
	}

	requiredEnvVars := []string{
		"SNOWFLAKE_ACCOUNT",
		"SNOWFLAKE_USER",
		"SNOWFLAKE_WAREHOUSE",
		"SNOWFLAKE_DATABASE",
		"SNOWFLAKE_SCHEMA",
		"SNOWFLAKE_PRIVATE_KEY",
		"SNOWFLAKE_PASSPHRASE",
	}

	if err := utils.ValidateEnvVars(requiredEnvVars); err != nil {
		log.Logger.Fatal().Msg(err.Error())
		os.Exit(1)
	}

	_ = gosnowflake.GetLogger().SetLogLevel("warn")

}

func main() {
	r := gin.Default()

	sfConfig := snowflake.SnowflakeConfig{
		Account:           os.Getenv("SNOWFLAKE_ACCOUNT"),
		User:              os.Getenv("SNOWFLAKE_USER"),
		Warehouse:         os.Getenv("SNOWFLAKE_WAREHOUSE"),
		Database:          os.Getenv("SNOWFLAKE_DATABASE"),
		Schema:            os.Getenv("SNOWFLAKE_SCHEMA"),
		EncodedPrivateKey: os.Getenv("SNOWFLAKE_PRIVATE_KEY"),
		Passphrase:        os.Getenv("SNOWFLAKE_PASSPHRASE"),
	}

	repoFactory, err := factory.NewFactory(&sfConfig)
	if err != nil {
		log.Fatal().Msg(err.Error())
		os.Exit(1)
	}

	columnListRepo := repository.NewColumnListRepository(repoFactory.SnowflakeDb)
	columnListService := services.NewColumnListService(columnListRepo)
	columnListHandler := handler.NewColumnListHandler(columnListService)

	r.GET("/columns", columnListHandler.GetAllColumnsInfo)

	r.Run("0.0.0.0:8080")

}
