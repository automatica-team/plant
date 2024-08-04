package db

import (
	"automatica.team/plant"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (*DB) Name() string {
	return "plant/db"
}

type DB struct {
	*gorm.DB
}

func New() *DB {
	return &DB{}
}

func (d *DB) Import(m plant.M) error {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}

	db, err := gorm.Open(postgres.Open(m.Get("dsn")), config)
	if err != nil {
		return err
	}

	*d = DB{
		DB: db,
	}

	return nil
}
