package core

import (
	"io/ioutil"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func (mod *Core) onLanguage(c tele.Context) error {
	langs, err := mod.getAvailableLanguages()
	if err != nil {
		return err
	}

	keyboard := &tele.ReplyMarkup{InlineKeyboard: make([][]tele.InlineButton, len(langs))}
	for _, lang := range langs {
		btn := tele.InlineButton{
			Text: lang,
			Data: "set_language_" + lang,
		}
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tele.InlineButton{btn})
	}
	return c.Send("Choose your language:", keyboard)
}

func (mod *Core) onSetLanguage(c tele.Context) error {
	language := strings.TrimSpace(strings.TrimPrefix(c.Callback().Data, "set_language_"))

	if err := mod.db.
		Model(&User{}).
		Where("id = ?", c.Sender().ID).
		Update("lang", language).
		Error; err != nil {
		return err
	}

	return c.Edit("Language updated to " + language)
}

func (mod *Core) getAvailableLanguages() (languages []string, err error) {
	files, err := ioutil.ReadDir("locales")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yml") {
			lang := strings.TrimSuffix(file.Name(), ".yml")
			languages = append(languages, lang)
		}
	}
	return languages, nil
}
