package plant

import (
	tele "gopkg.in/telebot.v3"
)

type Mod interface {
	Name() string
	Import(M) error
	Handler(string) tele.HandlerFunc
}

func (p *Plant) Add(mod Mod) {
	p.mods[mod.Name()] = mod
}
