package telegram

import (
	"context"
	"github.com/0xpelamar/chatbot/internal/entity"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) registerMiddleWare(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			LastName:  c.Sender().LastName,
			Username:  c.Sender().Username,
		}
		acc, created, err := t.App.Account.CreateOrUpdate(context.Background(), acc)
		c.Set("account", acc)
		if err != nil {
			return err
		}

		c.Set("is_just_created", created)

		return next(c)
	}
}
