package exts

import (
	"context"
	"time"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/api context value " + k.name }

type ctxKey string

var (
	ctxCancelsKey = &contextKey{"ctxCancels"}
	ctxCancelKey  = &contextKey{"ctxErr"}
)

var ctxErr error = nil

func CancelFromContext(ctx context.Context) context.CancelFunc {
	if u, ok := ctx.Value(ctxCancelKey).(context.CancelFunc); ok {
		return u
	}
	return nil
}

func CancelsFromContext(ctx context.Context) []context.CancelFunc {
	if u, ok := ctx.Value(ctxCancelsKey).([]context.CancelFunc); ok {
		return u
	}
	return nil
}

func Abort(ctx context.Context, err error) {
	if u, ok := ctx.Value(ctxCancelKey).(context.CancelFunc); ok {
		ctxErr = err
		u()
	}
}

func Error() error {
	return ctxErr
}

func WithCancel(bg context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(bg)
	cancels := getCancels(ctx, cancel)
	return context.WithValue(ctx, ctxCancelsKey, cancels), cancel
}

func WithDeadline(bg context.Context, d time.Time) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(bg, d)
	cancels := getCancels(ctx, cancel)
	return context.WithValue(ctx, ctxCancelsKey, cancels), cancel
}

func WithCancels(bg context.Context, keys ...context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(bg)
	cancels := getCancels(ctx)
	for _, c := range keys {
		_, cc := context.WithCancel(c)
		cancels = append(cancels, cc)
	}
	return context.WithValue(ctx, ctxCancelsKey, cancels), cancel
}

func Cancel(ctx context.Context) {
	cancels := []context.CancelFunc{}
	if bgCancels, ok := ctx.Value(ctxCancelsKey).([]context.CancelFunc); ok {
		cancels = append(cancels, bgCancels...)
	}
	for _, cancelFunc := range cancels {
		if cancelFunc != nil {
			cancelFunc()
		}
	}
}

func getCancels(ctx context.Context, cancelFunc ...context.CancelFunc) []context.CancelFunc {
	cancels := []context.CancelFunc{}
	if bgCancels, ok := ctx.Value(ctxCancelsKey).([]context.CancelFunc); ok {
		cancels = append(cancels, bgCancels...)
	}
	cancels = append(cancels, cancelFunc...)
	return cancels
}
