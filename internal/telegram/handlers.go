package telegram

import (
	"gopkg.in/telebot.v4"
)

func (t *Telegram) setupHandlers() {

	t.bot.Use(t.registerMiddleWare)
	t.bot.Handle("/start", t.start)

	t.bot.Handle(telebot.OnText, t.textHandler)

	t.bot.Handle(&btnEditDisplayName, t.editDisplayName)
	t.bot.Handle(&btnEditAge, t.editAge)
	t.bot.Handle(&btnEditProvince, t.editProvince)
	t.bot.Handle(&btnEditGender, t.editGender)
}

func (t *Telegram) textHandler(c telebot.Context) error {
	if t.TelePrompt.Dispatch(c.Sender().ID, c) {
		return nil
	}
	if c.Message().Text == "✏️ ویرایش پروفایل" {
		return t.editProfile(c)
	}
	return c.Reply("I didn't understand your command ")
}
