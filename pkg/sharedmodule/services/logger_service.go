package services

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerService provides structured logging capabilities.
type LoggerService interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	With(fields ...Field) LoggerService
	Sync() error
}

// Field represents a structured logging field.
type Field = zap.Field

// Field constructors for structured logging.
var (
	String   = zap.String
	Int      = zap.Int
	Int64    = zap.Int64
	Float64  = zap.Float64
	Bool     = zap.Bool
	Err      = zap.Error
	Any      = zap.Any
	Duration = zap.Duration
)

// LoggerServiceZap is a zap-based implementation of LoggerService.
type LoggerServiceZap struct {
	logger *zap.Logger
}

var _ LoggerService = &LoggerServiceZap{}

// NewLoggerService creates a new logger configured based on the environment.
// In development mode, it uses a human-readable console format.
// In production mode, it uses JSON format for structured logging.
func NewLoggerService(debug bool) *LoggerServiceZap {
	var logger *zap.Logger
	var err error

	if debug || isDevEnvironment() {
		// Development: human-readable console output
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = config.Build()
	} else {
		// Production: JSON output for log aggregation
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = config.Build()
	}

	if err != nil {
		// Fallback to nop logger if configuration fails
		logger = zap.NewNop()
	}

	return &LoggerServiceZap{logger: logger}
}

// NewLoggerServiceNop creates a no-op logger for testing.
func NewLoggerServiceNop() *LoggerServiceZap {
	return &LoggerServiceZap{logger: zap.NewNop()}
}

func isDevEnvironment() bool {
	env := os.Getenv("APP_ENV")
	switch env {
	case "test", "local", "dev", "development":
		return true
	}
	return false
}

func (l *LoggerServiceZap) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *LoggerServiceZap) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *LoggerServiceZap) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *LoggerServiceZap) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *LoggerServiceZap) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

// With creates a child logger with the given fields attached.
func (l *LoggerServiceZap) With(fields ...Field) LoggerService {
	return &LoggerServiceZap{logger: l.logger.With(fields...)}
}

// Sync flushes any buffered log entries.
func (l *LoggerServiceZap) Sync() error {
	return l.logger.Sync()
}
