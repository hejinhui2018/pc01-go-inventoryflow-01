package store

import "testing"

func TestPutGetAndSnapshot(t *testing.T) {
	s, err := NewFileStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Put("state", map[string]int{"x": 1}); err != nil {
		t.Fatal(err)
	}
	var out map[string]int
	if err := s.Get("state", &out); err != nil || out["x"] != 1 {
		t.Fatalf("out=%v err=%v", out, err)
	}
	if _, err := s.Snapshot("state"); err != nil {
		t.Fatal(err)
	}
}
