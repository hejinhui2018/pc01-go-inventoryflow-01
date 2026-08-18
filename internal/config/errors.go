package config

import "errors"

var (
	ErrInvalidPort      = errors.New("port must be between 1 and 65535")
	ErrEmptyDataDir     = errors.New("data directory is required")
	ErrInvalidTimeout   = errors.New("request timeout must be positive")
	ErrInvalidBodyLimit = errors.New("body limit must be positive")
	ErrInvalidBatch     = errors.New("reconcile batch must be positive")
	ErrEmptyWarehouse   = errors.New("warehouse code is required")
)
