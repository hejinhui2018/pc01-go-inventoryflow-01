package metrics

import (
	"sync"
	"time"
)

type Sample struct {
	Name  string
	Value int64
	At    time.Time
}
type Registry struct {
	mu      sync.RWMutex
	values  map[string]int64
	history []Sample
}

func NewRegistry() *Registry { return &Registry{values: map[string]int64{}} }
func (r *Registry) Set(name string, value int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.values[name] = value
	r.history = append(r.history, Sample{Name: name, Value: value, At: time.Now().UTC()})
}
func (r *Registry) Add(name string, delta int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.values[name] += delta
	r.history = append(r.history, Sample{Name: name, Value: r.values[name], At: time.Now().UTC()})
}
func (r *Registry) Get(name string) (int64, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.values[name]
	return v, ok
}
func (r *Registry) History() []Sample {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]Sample(nil), r.history...)
}
