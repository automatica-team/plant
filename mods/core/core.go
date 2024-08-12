package core

import (
	"automatica.team/plant"
	"automatica.team/plant/deps/db"
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

func (mod *Core) Import(v plant.V) error {
	v.SetDefault("default_locale", "en")

	var (
		defLocale = v.GetString("default_locale")
	)

	lt := mod.b.Layout

	// Middlewares
	mod.Use(mod.Logger(m))
	mod.Use(lt.Middleware(defLocale, mod.userLocale))

	// Handlers
	mod.Handle("/start", mod.onStart)

	// Auto migrate DB table
	return mod.db.AutoMigrate(&User{})
}
