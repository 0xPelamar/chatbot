package telegram

import (
	"context"
	"fmt"
	"log/slog"

	"gopkg.in/telebot.v4"
)

func (t *Telegram) editDisplayName(c telebot.Context) error {
	_ = c.Delete()
	err := t.editDisplayNamePrompt(c)
	if err != nil {
		slog.Error("failed to edit display name", err)
	}
	return t.home(c)
}

func (t *Telegram) editDisplayNamePrompt(c telebot.Context) error {
	account := getAccount(c)
	msg, err := t.Input(c, InputConfig{
		Prompt:         "👋 Welcome to Chatbot\nPlease enter your display name. this can be changed later.",
		PromptKeyboard: nil,
		Validator: Validator{
			Check: func(msg *telebot.Message) bool {
				l := len([]rune(msg.Text))
				return (2 < l) && (l < 21)
			},
			OnInvalid: func(msg *telebot.Message) string {
				return "❌ Your name must be more than 3 and less than 21 characters"
			},
		},
		Confirm: ConfirmStep{
			BuildPrompt: func(msg *telebot.Message) string {
				return fmt.Sprintf("ℹ️ We call you «%s»\nDo you confirm?", msg.Text)
			},
		},
	})
	if err != nil {
		slog.Error("error while editing display name", "error", err)
		return err
	}
	account.DisplayName = msg.Text
	if err := t.App.Accounts.Update(context.Background(), account); err != nil {
		_ = c.Send(fmt.Sprintf("couldn't update your profile\n"))
		return err
	}
	c.Set("account", account)
	return c.Send(fmt.Sprintf("✅ Your profile updated successfully. Now we will call you «%s»", account.DisplayName))
}
