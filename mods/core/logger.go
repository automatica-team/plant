package core

import (
	"log/slog"
	"os"

	"automatica.team/plant"

	"github.com/go-telebot/pkg/monitor"
	tele "gopkg.in/telebot.v3"
)

func (mod *Core) Logger(m plant.M) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mod.Use(func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if update, ok := monitor.NewUpdate(c); ok {
				logger.Info("info", update)
			}
			return next(c)
		}
	})

	return nil
}
