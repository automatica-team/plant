package core

import (
	"automatica.team/plant"
	"automatica.team/plant/deps/db"
	tele "gopkg.in/telebot.v3"
)

func (c *Core) Name() string {
	return "plant/core"
}

func New(b *plant.Bot, deps plant.Deps) *Core {
	d, ok := deps["plant/db"].(*db.DB)
	if !ok {
		panic("plant/core: bad dependency (plant/db)")
	}

	return &Core{
		b:  b,
		db: d,
	}
}

type Core struct {
	b  *plant.Bot
	db *db.DB
}

func (c *Core) Prepare() error {
	lt := c.b.Layout

	// Register middlewares
	c.b.Use(lt.Middleware("en"))

	// Auto migrate DB table
	return c.db.AutoMigrate(&User{})
}

func (c *Core) Handler(joint string) tele.HandlerFunc {
	return map[string]tele.HandlerFunc{
		"/start": c.onStart,
	}[joint]
}

func (c *Core) onStart(ctx tele.Context) error {
	user := &User{ID: ctx.Sender().ID}
	if !c.userExists(user.ID) {
		if err := c.db.Create(user).Error; err != nil {
			return err
		}
	}

	return ctx.Send(
		c.b.Text(ctx, "start"),
		c.b.Markup(ctx, "start"),
	)
}
