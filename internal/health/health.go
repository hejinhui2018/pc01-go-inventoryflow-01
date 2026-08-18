package health

import (
	"runtime"
	"time"
)

type Status struct {
	Ready     bool      `json:"ready"`
	Version   string    `json:"version"`
	Go        string    `json:"go"`
	CheckedAt time.Time `json:"checked_at"`
}

func Check(version string) Status {
	return Status{Ready: true, Version: version, Go: runtime.Version(), CheckedAt: time.Now().UTC()}
}
