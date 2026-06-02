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
		args := c.Args()
		if len(args) > 0 {
			_ = c.Send(fmt.Sprintf("The arg received. %s", args[0]))
		} else {
			_ = c.Send("no arg received")
		}
		slog.Info(fmt.Sprintf("%d is not just created", acc.ID))
		return t.home(c)
	}
	slog.Info(fmt.Sprintf("%d is just created", acc.ID))

	if err := t.editDisplayNamePrompt(c); err != nil {
		slog.Error(fmt.Sprintf("%d editDisplayName failed", acc.ID))
		return err
	}
	return t.home(c)
}

func (t *Telegram) home(c telebot.Context) error {
	account := getAccount(c)
	return c.Send(fmt.Sprintf("What can I do for you, %s?", account.DisplayName), generateKeyboard(mainMenuKeyboard(account.Language)))

}
