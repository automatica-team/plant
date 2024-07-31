package plant

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	*layout.Layout
	*tele.Bot
}

func (p *Plant) Compose() (*Bot, error) {
	lt, err := layout.New(p.Bot.Config)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	return &Bot{
		Layout: lt,
		Bot:    b,
	}, nil
}
