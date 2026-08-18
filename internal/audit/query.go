package audit

import (
	"sort"
	"strings"
)

func Since(entries []Entry, actor string) []Entry {
	var out []Entry
	for _, e := range entries {
		if actor == "" || e.Actor == actor {
			out = append(out, e)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func Contains(entries []Entry, term string) []Entry {
	term = strings.ToLower(term)
	var out []Entry
	for _, e := range entries {
		if strings.Contains(strings.ToLower(e.Action), term) || strings.Contains(strings.ToLower(e.Detail), term) {
			out = append(out, e)
		}
	}
	return out
}
func (l *Log) Count(action string) int {
	n := 0
	for _, e := range l.Entries() {
		if e.Action == action {
			n++
		}
	}
	return n
}
