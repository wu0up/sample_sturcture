package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(debug bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// level
	log.Debug().Msg("debug")
	log.Info().Msg("info")
	log.Log().Str("type", "empty msg & level").Send()
	//設定global logger，把想要的欄位加上
	log.Logger = log.With().
		Caller().
		Logger()
	log.Info().Msg("info with default fields")
}

func InitConsoleLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%-5s]", i))
	}
	output.FormatCaller = func(i interface{}) string {
		return filepath.Base(fmt.Sprintf("%s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		if i != nil {
			return fmt.Sprintf("| %s |", i)
		} else {
			return "| |"
		}
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	log.Logger = zerolog.New(output).
		Level(zerolog.TraceLevel).
		With().
		Caller().
		Timestamp().
		Logger()
	log.Info().Str("for", "development").Msg("pretty logging")
}
