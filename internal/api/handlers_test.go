package api

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/config"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/store"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/workflow"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	st, _ := store.NewFileStore(t.TempDir())
	h := NewHandler(workflow.NewService(st), config.Default())
	rr := httptest.NewRecorder()
	h.Health(rr, httptest.NewRequest("GET", "/health", nil))
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
