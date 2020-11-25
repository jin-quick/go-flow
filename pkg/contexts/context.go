/**
 * Name: context.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/19
 * Description:
 */

package contexts

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const (
	traceKey contextKey = "trace_key"
)

// SetTraceFields トレース情報をcontextに追加する
func SetTraceFields(ctx context.Context, traceID, spanID uint64) context.Context {
	return context.WithValue(ctx, traceKey, []zap.Field{
		zap.Uint64("dd.trace_id", traceID),
		zap.Uint64("dd.span_id", spanID),
	})
}

// GetTraceFields contextにトレース情報をzap.Field配列へ変換して返却
func GetTraceFields(ctx context.Context) []zap.Field {
	t, ok := ctx.Value(traceKey).([]zap.Field)
	if !ok {
		return nil
	}
	return t
}

// End
