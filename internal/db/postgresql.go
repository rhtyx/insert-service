package db

import (
	"github.com/rhtyx/insert-service.git/internal/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQLConn() *gorm.DB {
	conn, err := openPostgreSQLConn(config.DSN)
	if err != nil {
		log.Error(err.Error())
	}

	return conn
}

func openPostgreSQLConn(dsn string) (*gorm.DB, error) {
	log.Println("open postgreSQL connection")

	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	conn.SetMaxIdleConns(config.DBMaxIdleConns)
	conn.SetConnMaxLifetime(config.DBMaxLifetime)
	conn.SetMaxOpenConns(config.DBMaxConns)

	return db, err
}
