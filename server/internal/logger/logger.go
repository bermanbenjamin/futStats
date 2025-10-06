package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap.Logger with additional functionality
type Logger struct {
	*zap.Logger
}

// Config holds logger configuration
type Config struct {
	Level      string
	Format     string // json or console
	OutputPath string
	ErrorPath  string
}

// New creates a new logger instance
func New(config Config) (*Logger, error) {
	// Set default values
	if config.Level == "" {
		config.Level = "debug"
	}
	if config.Format == "" {
		config.Format = "json"
	}
	if config.OutputPath == "" {
		config.OutputPath = "stdout"
	}
	if config.ErrorPath == "" {
		config.ErrorPath = "stderr"
	}

	// Parse log level
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	// Create encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create encoder
	var encoder zapcore.Encoder
	if config.Format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// Create write syncer
	var writeSyncer zapcore.WriteSyncer
	if config.OutputPath == "stdout" {
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else {
		file, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		writeSyncer = zapcore.AddSync(file)
	}

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{Logger: logger}, nil
}

// NewDevelopment creates a development logger with console output
func NewDevelopment() (*Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{Logger: logger}, nil
}

// NewProduction creates a production logger with JSON output
func NewProduction() (*Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{Logger: logger}, nil
}

// WithFields creates a new logger with additional fields
func (l *Logger) WithFields(fields ...zap.Field) *Logger {
	return &Logger{Logger: l.Logger.With(fields...)}
}

// WithRequestID adds request ID to logger
func (l *Logger) WithRequestID(requestID string) *Logger {
	return l.WithFields(zap.String("request_id", requestID))
}

// WithUserID adds user ID to logger
func (l *Logger) WithUserID(userID string) *Logger {
	return l.WithFields(zap.String("user_id", userID))
}

// WithComponent adds component name to logger
func (l *Logger) WithComponent(component string) *Logger {
	return l.WithFields(zap.String("component", component))
}

// LogHTTPRequest logs HTTP request details
func (l *Logger) LogHTTPRequest(method, path, userAgent, remoteAddr string, statusCode int, duration time.Duration, requestID string) {
	l.Info("HTTP request",
		zap.String("method", method),
		zap.String("path", path),
		zap.String("user_agent", userAgent),
		zap.String("remote_addr", remoteAddr),
		zap.Int("status_code", statusCode),
		zap.Duration("duration", duration),
		zap.String("request_id", requestID),
	)
}

// LogDatabaseQuery logs database query details
func (l *Logger) LogDatabaseQuery(query string, duration time.Duration, rowsAffected int64) {
	l.Debug("Database query",
		zap.String("query", query),
		zap.Duration("duration", duration),
		zap.Int64("rows_affected", rowsAffected),
	)
}

// LogAuthEvent logs authentication events
func (l *Logger) LogAuthEvent(event, userID, email string, success bool) {
	level := zap.InfoLevel
	if !success {
		level = zap.WarnLevel
	}

	l.Log(level, "Authentication event",
		zap.String("event", event),
		zap.String("user_id", userID),
		zap.String("email", email),
		zap.Bool("success", success),
	)
}

// LogBusinessEvent logs business logic events
func (l *Logger) LogBusinessEvent(event, entityType, entityID string, metadata map[string]interface{}) {
	fields := []zap.Field{
		zap.String("event", event),
		zap.String("entity_type", entityType),
		zap.String("entity_id", entityID),
	}

	// Add metadata fields
	for key, value := range metadata {
		fields = append(fields, zap.Any(key, value))
	}

	l.Info("Business event", fields...)
}

// LogError logs errors with context
func (l *Logger) LogError(err error, message string, fields ...zap.Field) {
	allFields := append([]zap.Field{zap.Error(err)}, fields...)
	l.Error(message, allFields...)
}

// LogPanic logs panic with context
func (l *Logger) LogPanic(err error, message string, fields ...zap.Field) {
	allFields := append([]zap.Field{zap.Error(err)}, fields...)
	l.Panic(message, allFields...)
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	return l.Logger.Sync()
}
