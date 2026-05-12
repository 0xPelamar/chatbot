package telegram

import (
	"context"
	"log/slog"

	"github.com/0xpelamar/chatbot/internal/entity"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) registerMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			LastName:  c.Sender().LastName,
			Username:  c.Sender().Username,
		}
		acc, created, err := t.App.Accounts.CreateOrUpdate(context.Background(), acc)
		if err != nil {
			slog.Error("failed to register telegram account")
			return err
		}
		c.Set("account", acc)
		c.Set("is_just_created", created)
		return next(c)
	}
}
