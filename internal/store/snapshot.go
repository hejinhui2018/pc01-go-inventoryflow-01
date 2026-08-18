package store

import (
	"os"
	"path/filepath"
	"time"
)

type SnapshotMeta struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Bytes     int64     `json:"bytes"`
}

func (s *FileStore) Snapshot(name string) (SnapshotMeta, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p := s.path(name)
	st, err := os.Stat(p)
	if err != nil {
		return SnapshotMeta{}, err
	}
	copyPath := filepath.Join(s.root, name+"-"+time.Now().UTC().Format("20060102T150405")+".snapshot")
	in, err := os.ReadFile(p)
	if err != nil {
		return SnapshotMeta{}, err
	}
	if err := os.WriteFile(copyPath, in, 0o600); err != nil {
		return SnapshotMeta{}, err
	}
	return SnapshotMeta{Name: copyPath, CreatedAt: time.Now().UTC(), Bytes: st.Size()}, nil
}
