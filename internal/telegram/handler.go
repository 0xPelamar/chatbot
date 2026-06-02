package telegram

import "gopkg.in/telebot.v4"

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddleware)
	t.bot.Handle("/start", t.start)
	t.bot.Handle(telebot.OnText, t.textHandler)
	t.bot.Handle(&btnEditDisplayName, t.editDisplayName)
	t.bot.Handle(&btnEditAge, t.editAge)
	t.bot.Handle(&btnEditProvince, t.editProvince)
	t.bot.Handle(&btnEditGender, t.editGender)
}

func (t *Telegram) textHandler(c telebot.Context) error {
	if t.prompter.Deliver(c.Sender().ID, c) {
		return nil
	}
	if c.Message().Text == "Edit profile 🧑" {
		return t.editProfile(c)
	}
	return t.home(c)
}
