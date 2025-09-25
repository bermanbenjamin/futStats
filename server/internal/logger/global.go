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
		// Check if we're in development mode
		if os.Getenv("ENVIRONMENT") == "development" || os.Getenv("ENVIRONMENT") == "" {
			globalLogger, err = NewDevelopment()
		} else {
			globalLogger, err = NewProduction()
		}
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

// Info logs at info level using global logger
func Info(msg string, fields ...zap.Field) {
	GetGlobal().Info(msg, fields...)
}

// Debug logs at debug level using global logger
func Debug(msg string, fields ...zap.Field) {
	GetGlobal().Debug(msg, fields...)
}

// Warn logs at warn level using global logger
func Warn(msg string, fields ...zap.Field) {
	GetGlobal().Warn(msg, fields...)
}

// Error logs at error level using global logger
func Error(msg string, fields ...zap.Field) {
	GetGlobal().Error(msg, fields...)
}

// Fatal logs at fatal level using global logger
func Fatal(msg string, fields ...zap.Field) {
	GetGlobal().Fatal(msg, fields...)
}

// Panic logs at panic level using global logger
func Panic(msg string, fields ...zap.Field) {
	GetGlobal().Panic(msg, fields...)
}
