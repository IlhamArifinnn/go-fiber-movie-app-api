package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zdacoder/go-fiber-movie-app-api/config"
)

func Init(config *config.Config) {
	var level zerolog.Level

	switch strings.ToLower(config.LogLevel) {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	if config.AppEnv == "development" {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		}
		log.Logger = log.Output(consoleWriter)
	} else {
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	log.Logger = log.Logger.With().Str("service", "movie-app-api").Logger()

	log.Info().Msgf("Logger initialized with level: %s", level)
}
