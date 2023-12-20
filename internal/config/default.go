package config

import (
	"time"
)

const (
	DBMaxIdleConns int           = 4
	DBMaxConns     int           = 100
	DBMaxLifetime  time.Duration = 1 * time.Hour
	TotalWorker    int           = 100
	CSVName        string        = "majestic_million.csv"
	DSN            string        = "postgres://root:root@localhost:5432/insert_service?sslmode=disable&application_name=insert_service"
)
