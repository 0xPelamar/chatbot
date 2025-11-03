package telegram

import "gopkg.in/telebot.v4"

func (t *Telegram) setupHandlers() {
	t.bot.Use()
	t.bot.Handle(telebot.OnText, t.start)
}

func (t *Telegram) start(c telebot.Context) error {
	return c.Reply("It's Pelamar")
}
