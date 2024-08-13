package plant

import (
	"fmt"

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
	for _, end := range p.Bot.Expose {
		b.handle(end)
	}
	return nil
}

func (b *Bot) handle(end string) {
	b.Handle(end, func(c tele.Context) error {
		for _, h := range b.h[end] {
			if err := h(c); err != nil {
				return err
			}
		}
		return nil
	}, b.m[end]...)
}

func (b *Bot) Start() error {
	for _, v := range b.c[Startup] {
		v()
	}

	return b.Start()
}

type On = string

const (
	Startup On = "Startup"
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

func (h *Handler) Expose() []string {
	return nil
}
