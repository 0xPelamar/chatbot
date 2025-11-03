package telegram

import (
	"gopkg.in/telebot.v4"
)

func (t *Telegram) setupHandlers() {
	t.bot.Use()
	t.bot.Handle("/start", t.start)
	t.bot.Handle(telebot.OnText, t.textHandler)
}

func (t *Telegram) textHandler(c telebot.Context) error {
	if t.TelePrompt.Dispatch(c.Sender().ID, c) {
		return nil
	}
	// TODO per state
	return c.Reply("I didn't understand your command ")
}
