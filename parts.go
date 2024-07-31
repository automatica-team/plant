package plant

import (
	tele "gopkg.in/telebot.v3"
)

type Part interface {
	Name() string
	Prepare() error
	Handler(string) tele.HandlerFunc
}

func (p *Plant) Add(part Part) {
	p.parts[part.Name()] = part
}
