package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/api"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/config"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/store"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/workflow"
)

func main() {
	port := flag.Int("port", 8080, "HTTP listen port")
	data := flag.String("data", "./data", "data directory")
	flag.Parse()
	if err := run(*port, *data); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func run(port int, dataDir string) error {
	cfg := config.Default()
	cfg.Port = port
	cfg.DataDir = dataDir
	st, err := store.NewFileStore(cfg.DataDir)
	if err != nil {
		return fmt.Errorf("create store: %w", err)
	}
	svc := workflow.NewService(st)
	server := api.NewServer(svc, cfg)
	return server.ListenAndServe()
}
