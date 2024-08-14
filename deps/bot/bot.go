package bot

import (
	"automatica.team/plant"
	tele "gopkg.in/telebot.v3"
)

func (*Bot) Name() string {
	return "plant/bot"
}

type Bot struct {
	*tele.Bot
}

func New() *Bot {
	return &Bot{}
}

func (b *Bot) Import(m plant.M) error {
	b.Bot.Start()
	return nil
}
