package core

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

type User struct {
	CreatedAt time.Time
	ID        int64  `gorm:"primaryKey"`
	Lang      string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

// TODO: handle db errors

func (mod *Core) userLocale(r tele.Recipient) (lang string) {
	mod.db.
		Table("users").
		Select("lang").
		Where("id = ?", r.Recipient()).
		Find(&lang)
	return lang
}

func (mod *Core) userExists(id int64) bool {
	var count int64
	mod.db.Table("users").Where("id = ?", id).Count(&count)
	return count > 0
}
