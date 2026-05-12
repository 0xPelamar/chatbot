package telegram

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddleware)
	t.bot.Handle("/start", t.start)
}
