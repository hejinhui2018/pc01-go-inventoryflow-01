package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type FileStore struct {
	root string
	mu   sync.RWMutex
}

func NewFileStore(root string) (*FileStore, error) {
	if root == "" {
		return nil, errors.New("store root is required")
	}
	if err := os.MkdirAll(root, 0o750); err != nil {
		return nil, err
	}
	return &FileStore{root: root}, nil
}
func (s *FileStore) path(name string) string { return filepath.Join(s.root, name+".json") }
func (s *FileStore) Put(name string, value any) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(s.root, ".store-*")
	if err != nil {
		return err
	}
	nerr := func() error {
		if err := tmp.Chmod(0o600); err != nil {
			return err
		}
		if _, err := tmp.Write(append(b, '\n')); err != nil {
			return err
		}
		if err := tmp.Sync(); err != nil {
			return err
		}
		if err := tmp.Close(); err != nil {
			return err
		}
		return os.Rename(tmp.Name(), s.path(name))
	}()
	if nerr != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return fmt.Errorf("put %s: %w", name, nerr)
	}
	return nil
}
func (s *FileStore) Get(name string, out any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, err := os.ReadFile(s.path(name))
	if errors.Is(err, os.ErrNotExist) {
		return ErrNotFound
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, out); err != nil {
		return fmt.Errorf("decode %s: %w", name, err)
	}
	return nil
}
func (s *FileStore) Delete(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := os.Remove(s.path(name))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}
func (s *FileStore) Exists(name string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, err := os.Stat(s.path(name))
	return err == nil
}
