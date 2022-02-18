package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}

func (db *DB) Migrate(dst ...interface{}) error {
	err := db.Migrator().DropTable(dst...)
	if err != nil {
		return nil
	}

	return db.DB.AutoMigrate(dst...)
}
