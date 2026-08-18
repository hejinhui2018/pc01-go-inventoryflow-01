package config

import "time"

type Config struct {
	Port           int
	DataDir        string
	RequestTimeout time.Duration
	MaxBodyBytes   int64
	ReconcileBatch int
	WarehouseCode  string
}

func Default() Config {
	return Config{Port: 8080, DataDir: "./data", RequestTimeout: 5 * time.Second, MaxBodyBytes: 1 << 20, ReconcileBatch: 100, WarehouseCode: "WH-SH-01"}
}

func (c Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return ErrInvalidPort
	}
	if c.DataDir == "" {
		return ErrEmptyDataDir
	}
	if c.RequestTimeout <= 0 {
		return ErrInvalidTimeout
	}
	if c.MaxBodyBytes <= 0 {
		return ErrInvalidBodyLimit
	}
	if c.ReconcileBatch <= 0 {
		return ErrInvalidBatch
	}
	if c.WarehouseCode == "" {
		return ErrEmptyWarehouse
	}
	return nil
}
