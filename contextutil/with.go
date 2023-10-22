package contextutil

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrInterrupted = errors.New("received interrupt")

	InterruptSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM, os.Interrupt}
)

// creates a context that can be cancelled by an interrupt signal
func WithInterrupt(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	return WithInterruptCause(ctx, ErrInterrupted, signals...)
}

// creates a context that can be cancelled by an interrupt signal with a reason
func WithInterruptCause(ctx context.Context, cause error, signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	if len(signals) < 1 {
		signals = InterruptSignals
	}

	ctx, cancel := context.WithCancelCause(ctx)
	notifications := make(chan os.Signal, 1)
	signal.Notify(notifications, signals...)

	go func() {
		<-notifications
		cancel(cause)
	}()

	return ctx, cancel
}
