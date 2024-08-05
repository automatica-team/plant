package monitor

import (
	"automatica.team/plant"
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
	return nil
}
