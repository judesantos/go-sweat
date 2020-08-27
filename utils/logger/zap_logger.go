package logger

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zapLogger
)

func init() {

	// initialize logger subsystem
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.RFC3339TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	_log, err := logConfig.Build()

	if err != nil {
		panic(err)
	}

	log = &zapLogger{
		logger: _log,
	}
}

type zapLogger struct {
	logger *zap.Logger
}

// GetLogger - Get zap logger instance
func GetLogger() *zapLogger {
	return log
}

// Info - Log info message
func (l *zapLogger) Info(msg string, tags ...zap.Field) {
	log.logger.Info(msg, tags...)
	log.logger.Sync()
}

// Error - Log error message
func (l *zapLogger) Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.logger.Error(msg, tags...)
	log.logger.Sync()
}

// Debug - Log debug message
func (l *zapLogger) Debug(msg string, tags ...zap.Field) {
	log.logger.Debug(msg, tags...)
	log.logger.Sync()
}

func (l *zapLogger) Trace(msg string, tags ...zap.Field) {

	var fnName = "?"

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "?"
		line = -1
	} else {
		fn := runtime.FuncForPC(pc)
		fnName = fn.Name()
	}

	_msg := fmt.Sprintf("%s, %d, %s:", file, line, fnName)
	_msg = _msg + msg

	fmt.Printf("zaplogger trace: %s\n", msg)

	log.logger.Debug(_msg, tags...)
	log.logger.Sync()
}

// Warn - Log warning message
func (l *zapLogger) Warn(msg string, tags ...zap.Field) {
	log.logger.Warn(msg, tags...)
	log.logger.Sync()
}
