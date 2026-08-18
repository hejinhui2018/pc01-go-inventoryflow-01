package store

import "sync"

type KeyLock struct {
	mu   sync.Mutex
	keys map[string]*sync.Mutex
}

func NewKeyLock() *KeyLock { return &KeyLock{keys: make(map[string]*sync.Mutex)} }
func (l *KeyLock) Lock(key string) func() {
	l.mu.Lock()
	item := l.keys[key]
	if item == nil {
		item = &sync.Mutex{}
		l.keys[key] = item
	}
	l.mu.Unlock()
	item.Lock()
	return item.Unlock
}
