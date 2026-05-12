package telegram

import (
	"fmt"
	"log/slog"

	"gopkg.in/telebot.v4"
)

func (t *Telegram) start(c telebot.Context) error {
	acc := getAccount(c)
	isJustCreated := c.Get("is_just_created").(bool)
	if !isJustCreated {
		slog.Info(fmt.Sprintf("%d is not just created", acc.ID))
		return t.home(c)
	}
	slog.Info(fmt.Sprintf("%d is just created", acc.ID))

	if err := t.editDisplayNamePrompt(c); err != nil {
		slog.Error(fmt.Sprintf("%d editDisplayName failed", acc.ID))
		return err
	}
	acc = getAccount(c)
	data := fmt.Sprintf(
		"id: %v\n"+
			"first_name: %v\n"+
			"last_name: %v\n"+
			"username: %v\n"+
			"display_name: %v\n"+
			"age: %v\n"+
			"province: %v\n"+
			"joined_at: %v\n"+
			"gender: %v\n"+
			"language: %v\n"+
			"uuid: %v",
		acc.ID, acc.FirstName, acc.LastName, acc.Username,
		acc.DisplayName, acc.Age, acc.Province, acc.JoinedAt,
		acc.Gender, acc.Language, acc.UUID,
	)
	c.Send(data)
	return t.home(c)
}

func (t *Telegram) home(c telebot.Context) error {
	account := getAccount(c)

	return c.Send(fmt.Sprintf("Hello %s.\nWhat can I do for you?", account.DisplayName))
}
