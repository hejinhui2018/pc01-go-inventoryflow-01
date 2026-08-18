package clock

import "time"

type Window struct {
	Start time.Time
	End   time.Time
}

func NewWindow(start, end time.Time) Window { return Window{Start: start, End: end} }
func (w Window) Valid() bool                { return !w.Start.IsZero() && !w.End.IsZero() && !w.End.Before(w.Start) }
func (w Window) Contains(t time.Time) bool  { return w.Valid() && !t.Before(w.Start) && t.Before(w.End) }
func (w Window) Duration() time.Duration {
	if !w.Valid() {
		return 0
	}
	return w.End.Sub(w.Start)
}
func SplitDay(start time.Time, days int) []Window {
	if days < 1 {
		return nil
	}
	out := make([]Window, days)
	for i := range out {
		a := start.AddDate(0, 0, i)
		out[i] = Window{Start: a, End: a.Add(24 * time.Hour)}
	}
	return out
}
