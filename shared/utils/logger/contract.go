package logger

import "context"

type Logger interface {
	Info(ctx context.Context, message string, args ...any)
	Error(ctx context.Context, message string, args ...any)
}
