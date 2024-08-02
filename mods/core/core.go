package core

import (
	"automatica.team/plant"
	"automatica.team/plant/deps/db"

	tele "gopkg.in/telebot.v3"
)

func (mod *Core) Name() string {
	return "plant/core"
}

type Core struct {
	plant.Handler
	b  *plant.Bot `plant:"bot"`
	db *db.DB     `plant:"dep:plant/db"`
}

func New() *Core {
	return &Core{}
}

func (mod *Core) Import(_ plant.M) error {
	// Middlewares
	mod.Use(mod.b.Layout.Middleware("en"))

	// Handlers
	mod.Handle("/start", mod.onStart)

	// Auto migrate DB table
	return mod.db.AutoMigrate(&User{})
}

func (mod *Core) onStart(c tele.Context) error {
	user := User{ID: c.Sender().ID}
	if !mod.userExists(user.ID) {
		if err := mod.db.Create(&user).Error; err != nil {
			return err
		}
	}

	return c.Send(
		mod.b.Text(c, "start"),
		mod.b.Markup(c, "start"),
	)
}
