package monitor

import (
	"automatica.team/plant"
	"github.com/go-telebot/pkg/monitor"
	tele "gopkg.in/telebot.v3"
)

func (mod *Monitor) Name() string {
	return "plant/monitor"
}

type Monitor struct {
	plant.Handler
	b *plant.Bot `plant:"bot"`
}

func New() *Monitor {
	return &Monitor{}
}

func (mod *Monitor) Import(v plant.V) error {
	mon, err := monitor.New(monitor.Config{
		URL: v.GetEnv("url"),
	})
	if err != nil {
		return err
	}

	mod.Use(func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if err := next(c); err != nil {
				mon.Error(c, err.Error())
			}
			return nil
		}
	})

	mod.Use(mon.Middleware())
	return nil
}
