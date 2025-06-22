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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Debug().Str("Logging level", level).Msg("Logger initialized")
}

func GetLogger() zerolog.Logger {
	return log.Logger
}
