package config

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, *sql.DB, error) {
	LoadEnv()

	host := MustGetEnv("DB_HOST")
	port := MustGetEnv("DB_PORT")
	user := MustGetEnv("DB_USER")
	pass := MustGetEnv("DB_PASSWORD")
	name := MustGetEnv("DB_NAME")
	ssl := GetEnv("DB_SSLMODE", "")

	var dsn string
	if ssl != "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, name, ssl)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	return db, sqlDB, nil
}

type ErrEnv string

func (e ErrEnv) Error() string { return "missing env var: " + string(e) }
