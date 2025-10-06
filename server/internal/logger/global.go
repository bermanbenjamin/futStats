package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

var (
	globalLogger *Logger
	once         sync.Once
)

// InitGlobal initializes the global logger
func InitGlobal(config Config) error {
	var err error
	once.Do(func() {
		// Always use the config provided, but ensure it outputs to stdout/stderr for Railway
		config.OutputPath = "stdout"
		config.ErrorPath = "stderr"

		// Force JSON format in production for better log parsing
		environment := os.Getenv("ENVIRONMENT")
		if environment == "production" || environment == "railway" {
			config.Format = "json"
		}

		globalLogger, err = New(config)
	})
	return err
}

// GetGlobal returns the global logger instance
func GetGlobal() *Logger {
	if globalLogger == nil {
		// Fallback to development logger if not initialized
		globalLogger, _ = NewDevelopment()
	}
	return globalLogger
}

// SetGlobal sets the global logger instance
func SetGlobal(logger *Logger) {
	globalLogger = logger
}

// DEPRECATED: Use dependency injection instead of global functions
// These functions are kept for backward compatibility but should not be used in new code

// Info logs at info level using global logger
// DEPRECATED: Use logger instance from dependency injection
func Info(msg string, fields ...zap.Field) {
	GetGlobal().Info(msg, fields...)
}

// Debug logs at debug level using global logger
// DEPRECATED: Use logger instance from dependency injection
func Debug(msg string, fields ...zap.Field) {
	GetGlobal().Debug(msg, fields...)
}

// Warn logs at warn level using global logger
// DEPRECATED: Use logger instance from dependency injection
func Warn(msg string, fields ...zap.Field) {
	GetGlobal().Warn(msg, fields...)
}

// Error logs at error level using global logger
// DEPRECATED: Use logger instance from dependency injection
func Error(msg string, fields ...zap.Field) {
	GetGlobal().Error(msg, fields...)
}

// Fatal logs at fatal level using global logger
// DEPRECATED: Use logger instance from dependency injection
func Fatal(msg string, fields ...zap.Field) {
	GetGlobal().Fatal(msg, fields...)
}

// Panic logs at panic level using global logger
// DEPRECATED: Use logger instance from dependency injection
func Panic(msg string, fields ...zap.Field) {
	GetGlobal().Panic(msg, fields...)
}
