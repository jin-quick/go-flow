/*
 * Name: logger.go
 * Created by: Masato Shima / Mengzhuang Jin
 * Created on: 2020/10/16
 * Description: Handler for logging.
 */

package logger

import (
	"context"
	"github.com/jin-quick/go-flow/pkg/contexts"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// Debug レベルログを出力
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if logger.Core().Enabled(zap.DebugLevel) {
		write(ctx, zap.DebugLevel, msg, fields)
	}
}

// Warn レベルログを出力
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	write(ctx, zap.WarnLevel, msg, fields)
}

// Info レベルログを出力
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	write(ctx, zap.InfoLevel, msg, fields)
}

// Error レベルログを出力
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	write(ctx, zap.ErrorLevel, msg, fields)
}

// ErrorNoStack Stack出力なしで Error レベルログを出力
func ErrorNoStack(ctx context.Context, msg string, fields ...zap.Field) {
	if ce := noStackLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(contexts.GetTraceFields(ctx)...)
		ce.Write(fields...)
	}
}

func write(ctx context.Context, level zapcore.Level, msg string, fields []zap.Field) {
	if ce := logger.Check(level, msg); ce != nil {
		fields = append(fields, contexts.GetTraceFields(ctx)...)
		ce.Write(fields...)
	}
}

// End
