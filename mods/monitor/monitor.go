package monitor

import (
	"automatica.team/plant"
	"github.com/go-telebot/pkg/monitor"
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

func (mod *Monitor) Import(m plant.M) error {
	mon, err := monitor.New(monitor.Config{URL: m.Get("url")})
	if err != nil {
		return err
	}
	mod.Use(mon.Middleware())

	return nil
}
