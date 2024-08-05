package widgets

import tele "gopkg.in/telebot.v3"

type SendEditable interface {
	tele.Sendable
	tele.Editable
}
