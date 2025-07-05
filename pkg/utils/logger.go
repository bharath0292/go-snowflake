package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initializes zerolog with console + file output using rotation.
func InitLogger() {
	// Use RFC3339 time format for better compatibility (e.g. Prometheus, Loki, Grafana)
	zerolog.TimeFieldFormat = time.RFC3339

	// Optional: Enable human-friendly console writer in dev
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	consoleWriter.FormatLevel = func(i any) string {
		if i == nil {
			return ""
		}
		return " | " + strings.ToUpper(i.(string))[:3] + " | "
	}
	consoleWriter.FormatCaller = func(i any) string {
		if i == nil {
			return ""
		}
		return i.(string) + " | "
	}

	// File logger with rotation
	logDir := "./logs"
	os.MkdirAll(logDir, 0755)

	rotatingLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, "server.log"),
		MaxSize:    50,   // MB
		MaxBackups: 30,   // Keep last 30 logs
		MaxAge:     7,    // Days
		Compress:   true, // gzip
	}

	// Setup global logger to write to both console and file
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter, rotatingLogger)).
		With().
		Timestamp().
		Caller().
		Logger()

	log.Info().Msg("Zerolog initialized successfully.")
}
