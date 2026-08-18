package queue

import (
	"context"
	"sync"
)

type Queue[T any] struct {
	mu    sync.Mutex
	items []T
	wake  chan struct{}
}

func New[T any]() *Queue[T] { return &Queue[T]{wake: make(chan struct{}, 1)} }
func (q *Queue[T]) Push(v T) {
	q.mu.Lock()
	q.items = append(q.items, v)
	q.mu.Unlock()
	select {
	case q.wake <- struct{}{}:
	default:
	}
}
func (q *Queue[T]) Pop(ctx context.Context) (T, bool) {
	for {
		q.mu.Lock()
		if len(q.items) > 0 {
			v := q.items[0]
			q.items = q.items[1:]
			q.mu.Unlock()
			return v, true
		}
		q.mu.Unlock()
		select {
		case <-q.wake:
		case <-ctx.Done():
			var zero T
			return zero, false
		}
	}
}
func (q *Queue[T]) Len() int { q.mu.Lock(); defer q.mu.Unlock(); return len(q.items) }
