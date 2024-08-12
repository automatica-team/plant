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

func (d *DB) Import(v plant.V) error {
	v.SetDefault("log_level", "silent")

	var (
		logLevel = v.GetString("log_level")
		dsn      = v.GetEnv("dsn")
	)

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevels[logLevel]),
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return err
	}

	*d = DB{
		DB: db,
	}

	return nil
}
