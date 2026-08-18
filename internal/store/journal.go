package store

import (
	"bufio"
	"encoding/json"
	"os"
)

type Journal struct{ path string }

func NewJournal(root, name string) *Journal {
	return &Journal{path: root + string(os.PathSeparator) + name + ".jsonl"}
}
func (j *Journal) Append(value any) error {
	f, err := os.OpenFile(j.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = f.Write(append(b, '\n'))
	if err == nil {
		err = f.Sync()
	}
	return err
}
func (j *Journal) Read(out func([]byte) error) error {
	f, err := os.Open(j.path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if err := out(append([]byte(nil), sc.Bytes()...)); err != nil {
			return err
		}
	}
	return sc.Err()
}
func (j *Journal) Path() string { return j.path }
