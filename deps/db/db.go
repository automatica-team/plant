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

var logLevels = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func (d *DB) Import(m plant.M) error {
	var (
		logLevel = m.GetOr("log_level", "silent")
	)

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevels[logLevel]),
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
