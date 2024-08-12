package core

import tele "gopkg.in/telebot.v3"

func (mod *Core) onStart(c tele.Context) error {
	user := &User{
		ID:   c.Sender().ID,
		Lang: c.Sender().LanguageCode,
	}

	if !mod.userExists(user.ID) {
		if err := mod.db.Create(user).Error; err != nil {
			return err
		}
	}

	return c.Send(
		mod.b.Text(c, "start"),
		mod.b.Markup(c, "start"),
	)
}
