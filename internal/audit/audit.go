package audit

import (
	"fmt"
	"sync"
	"time"
)

type Entry struct {
	At     time.Time `json:"at"`
	Actor  string    `json:"actor"`
	Action string    `json:"action"`
	Ref    string    `json:"ref"`
	Detail string    `json:"detail"`
}
type Log struct {
	mu      sync.RWMutex
	entries []Entry
}

func New() *Log { return &Log{} }
func (l *Log) Record(actor, action, ref, detail string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = append(l.entries, Entry{At: time.Now().UTC(), Actor: actor, Action: action, Ref: ref, Detail: detail})
}
func (l *Log) Entries() []Entry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return append([]Entry(nil), l.entries...)
}
func (e Entry) String() string {
	return fmt.Sprintf("%s %s %s %s", e.At.Format(time.RFC3339), e.Actor, e.Action, e.Ref)
}
