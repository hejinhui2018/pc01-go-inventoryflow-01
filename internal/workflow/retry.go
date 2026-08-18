package workflow

import (
	"context"
	"time"
)

func Retry(ctx context.Context, attempts int, delay time.Duration, fn func() error) error {
	if attempts < 1 {
		attempts = 1
	}
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return err
}
func Backoff(base time.Duration, attempt int) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 10 {
		attempt = 10
	}
	return base * time.Duration(1<<attempt)
}
