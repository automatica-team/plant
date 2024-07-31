package core

import "time"

type User struct {
	CreatedAt time.Time
	ID        int64 `gorm:"primarykey"`
}

func (User) TableName() string {
	return "users"
}

func (c *Core) userExists(id int64) bool {
	var count int64
	c.db.Table("users").Where("id = ?", id).Count(&count)
	return count > 0
}
