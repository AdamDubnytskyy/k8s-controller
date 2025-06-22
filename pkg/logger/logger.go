// Logger package implements structured logging by leveraging [zerolog](https://github.com/rs/zerolog)
package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
	Init sets Global logging level
	You can set the Global logging level to any of these options using the SetGlobalLevel function in the zerolog package, 
	passing in one of the given constants above, e.g. zerolog.InfoLevel would be the "info" level.
*/
func Init(level string) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr)}

	switch strings.ToLower(lvl) {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}

	log.Debug().Str("Logging level", level).Msg("Logger initialized")
}

func GetLogger() zerolog.Logger {
	return log.Logger
}