package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Filed struct {
	Name  string
	Value interface{}
}

type Interface interface {
	Debug(message interface{}, args ...Filed)
	Info(message interface{}, args ...Filed)
	Warn(message interface{}, args ...Filed)
	Error(message interface{}, args ...Filed)
	Fatal(message interface{}, args ...Filed)
}

type Logger struct {
	logger *zap.Logger
}

var _ Interface = (*Logger)(nil)

func New(file, level string) (*Logger, error) {
	var atomicLevel zap.AtomicLevel

	switch level {
	case "warn":
		atomicLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		atomicLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "info":
		atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "debug":
		atomicLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	outputPaths := []string{"stdout"}
	if file != "" {
		outputPaths = append(outputPaths, file)
	}

	cfg := zap.Config{
		Level:       atomicLevel,
		Encoding:    "console",
		OutputPaths: outputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime: zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05 -0700"))
			}),
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Logger{logger: logger}, nil
}

func (l *Logger) Debug(message interface{}, args ...Filed) {
	msg := l.msg(message)
	zapArgs := l.args(args...)

	l.logger.Debug(msg, zapArgs...)
}

func (l *Logger) Info(message interface{}, args ...Filed) {
	msg := l.msg(message)
	zapArgs := l.args(args...)

	l.logger.Info(msg, zapArgs...)
}

func (l *Logger) Warn(message interface{}, args ...Filed) {
	msg := l.msg(message)
	zapArgs := l.args(args...)

	l.logger.Warn(msg, zapArgs...)
}

func (l *Logger) Error(message interface{}, args ...Filed) {
	msg := l.msg(message)
	zapArgs := l.args(args...)

	l.logger.Error(msg, zapArgs...)
}

func (l *Logger) Fatal(message interface{}, args ...Filed) {
	msg := l.msg(message)
	zapArgs := l.args(args...)

	l.logger.Fatal(msg, zapArgs...)
}

func (l *Logger) msg(message interface{}) string {
	switch message := message.(type) {
	case error:
		return message.Error()
	case string:
		return message
	}

	return fmt.Sprint(message)
}

func (l *Logger) args(args ...Filed) []zap.Field {
	var zapArgs []zap.Field
	for _, arg := range args {
		switch a := arg.Value.(type) {
		case string:
			zapArgs = append(zapArgs, zap.String(arg.Name, a))
		case int:
			zapArgs = append(zapArgs, zap.Int(arg.Name, a))
		default:
			zapArgs = append(zapArgs, zap.String(arg.Name, fmt.Sprint(arg.Value)))
		}
	}

	return zapArgs
}
