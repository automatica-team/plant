package plant

import (
	"fmt"

	"golang.org/x/exp/maps"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	*layout.Layout
	*tele.Bot

	h map[any][]tele.HandlerFunc
	m map[any][]tele.MiddlewareFunc
	c map[any][]func()
}

func (p *Plant) composeBot() (*Bot, error) {
	lt, err := layout.New(p.Config.Bot.File)
	if err != nil {
		return nil, fmt.Errorf("composeBot: %w", err)
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, fmt.Errorf("composeBot: %w", err)
	}

	return &Bot{
		Layout: lt,
		Bot:    b,

		h: make(map[any][]tele.HandlerFunc),
		m: make(map[any][]tele.MiddlewareFunc),
		c: make(map[any][]func()),
	}, nil
}

func (p *Plant) exposeBot(b *Bot) error {
	for _, mod := range p.mods {
		for _, end := range mod.Expose() {
			b.handle(end)
		}
	}
	return nil
}

func (b *Bot) Start() {
	b.callOn(Startup)
	b.Bot.Start()
}

func (b *Bot) handle(end any) {
	b.Handle(end, func(c tele.Context) error {
		for _, h := range b.h[end] {
			if err := h(c); err != nil {
				return err
			}
		}
		return nil
	}, b.m[end]...)
}

func (b *Bot) callOn(on On) {
	for _, do := range b.c[on] {
		do()
	}
}

type On = string

const (
	Startup On = "plant:startup"
)

type Handler struct {
	b *Bot
}

func (h *Handler) On(on On, do func()) {
	h.b.c[on] = append(h.b.c[on], do)
}

func (h *Handler) Use(middle ...tele.MiddlewareFunc) {
	h.b.Use(middle...)
}

func (h *Handler) Handle(end any, handler tele.HandlerFunc, m ...tele.MiddlewareFunc) {
	h.b.h[end] = append(h.b.h[end], handler)
	if len(m) > 0 {
		h.b.m[end] = append(h.b.m[end], m...)
	}
}

func (h *Handler) Expose() []any {
	return maps.Keys(h.b.h)
}
