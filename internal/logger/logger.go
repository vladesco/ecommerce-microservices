package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type LogLevel string

const (
	FATAL = "FATAL"
	ERROR = "ERROR"
	WARN  = "WARN"
	INFO  = "INFO"
	DEBUG = "DEBUG"
	TRACE = "TRACE"
)

type LogConfig struct {
	Level LogLevel `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

func NewLogger(config LogConfig) zerolog.Logger {
	logLevel := convertLogLevel(config.Level)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if logLevel <= zerolog.WarnLevel {
		return zerolog.New(os.Stdout).
			Level(logLevel).
			With().
			Timestamp().
			Logger()
	} else {
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "03:04:05.000PM"}).
			Level(logLevel).
			With().
			Timestamp().
			Logger()
	}
}

func convertLogLevel(level LogLevel) zerolog.Level {
	switch level {
	case FATAL:
		return zerolog.FatalLevel
	case ERROR:
		return zerolog.ErrorLevel
	case WARN:
		return zerolog.WarnLevel
	case INFO:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case TRACE:
		return zerolog.TraceLevel
	default:
		return zerolog.Disabled
	}
}
