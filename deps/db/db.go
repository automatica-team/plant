package db

import (
	"automatica.team/plant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (*DB) Name() string {
	return "plant/db"
}

func init() {
	plant.Inject(&DB{})
}

type DB struct {
	*gorm.DB
}

func (d *DB) Connect(m plant.M) error {
	db, err := gorm.Open(postgres.Open(m.Get("dsn")))
	if err != nil {
		return err
	}

	*d = DB{
		DB: db,
	}

	return nil
}
