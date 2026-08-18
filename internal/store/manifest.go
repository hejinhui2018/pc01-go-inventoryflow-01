package store

import (
	"encoding/json"
	"os"
	"time"
)

type Manifest struct {
	Schema    int       `json:"schema"`
	CreatedAt time.Time `json:"created_at"`
	Records   []string  `json:"records"`
}

func (s *FileStore) WriteManifest(m Manifest) error {
	if m.Schema == 0 {
		m.Schema = 1
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().UTC()
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.root+string(os.PathSeparator)+"manifest.json", append(b, '\n'), 0o600)
}
func (s *FileStore) ReadManifest() (Manifest, error) {
	b, err := os.ReadFile(s.root + string(os.PathSeparator) + "manifest.json")
	if err != nil {
		return Manifest{}, err
	}
	var m Manifest
	err = json.Unmarshal(b, &m)
	return m, err
}
func (m Manifest) Has(name string) bool {
	for _, v := range m.Records {
		if v == name {
			return true
		}
	}
	return false
}
