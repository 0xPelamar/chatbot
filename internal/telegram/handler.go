package telegram

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddleware)
	t.bot.Handle("/start", t.start)
}
