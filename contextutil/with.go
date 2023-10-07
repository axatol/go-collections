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

func WithInterrupt(ctx context.Context, signals ...os.Signal) context.Context {
	return WithInterruptCause(ctx, ErrInterrupted, signals...)
}

func WithInterruptCause(ctx context.Context, cause error, signals ...os.Signal) context.Context {
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

	return ctx
}
