package queue

import (
	"context"
	"sync"
)

type Worker[T any] struct {
	Queue  *Queue[T]
	Handle func(context.Context, T) error
}

func (w Worker[T]) Run(ctx context.Context) error {
	if w.Queue == nil || w.Handle == nil {
		return nil
	}
	for {
		item, ok := w.Queue.Pop(ctx)
		if !ok {
			return ctx.Err()
		}
		if err := w.Handle(ctx, item); err != nil {
			return err
		}
	}
}
func RunMany[T any](ctx context.Context, n int, q *Queue[T], fn func(context.Context, T) error) error {
	if n < 1 {
		n = 1
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	errs := make(chan error, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); errs <- Worker[T]{Queue: q, Handle: fn}.Run(ctx) }()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil && err != context.Canceled {
			return err
		}
	}
	return nil
}
