package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (t *Telegram) editDisplayName(c telebot.Context) error {
	_ = c.Delete()
	err := t.editDisplayNamePrompt(c)
	if err != nil {
		slog.Error("failed to edit display name", "err", err)
	}
	return t.home(c)
}

func (t *Telegram) editDisplayNamePrompt(c telebot.Context) error {
	msg, err := t.Input(c, InputConfig{
		Prompt:         "Please enter your display name.",
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
		confirm: confirmStep{
			BuildPrompt: func(msg *telebot.Message) string {
				return fmt.Sprintf("ℹ️ We call you «%s»\nDo you confirm?", msg.Text)
			},
		},
	})
	if err != nil {
		slog.Error("error while editing display name", "error", err)
		return err
	}
	account := getAccount(c)
	account.DisplayName = msg.Text
	slog.Info(fmt.Sprintf("The user entered display name: %s", account.DisplayName))
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		_ = c.Send(fmt.Sprintf("couldn't update your profile\n"))
		return err
	}
	c.Set("account", account)
	_ = c.Send(fmt.Sprintf("✅ Your profile updated successfully. Now we will call you «%s»", account.DisplayName))
	return nil
}

func (t *Telegram) editAge(c telebot.Context) error {
	_ = c.Delete()
	err := t.editAgePrompt(c)
	if err != nil {
		slog.Error("failed to edit age", "err", err)
	}
	return t.home(c)
}
func (t *Telegram) editAgePrompt(c telebot.Context) error {
	account := getAccount(c)
	msg, err := t.Input(c, InputConfig{
		Prompt:         "Please enter your age.",
		PromptKeyboard: agesKeyboard(),
		Validator:      ageValidator(),
	})
	if err != nil {
		slog.Error("failed to edit age", "err", err)
		return err
	}
	newAge, _ := strconv.Atoi(msg.Text)
	account.Age = int64(newAge)
	slog.Info(fmt.Sprintf("The user entered age: %d", account.Age))
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		_ = c.Send(fmt.Sprintf("couldn't update your profile\n"))
		return err
	}
	c.Set("account", account)
	_ = c.Send(fmt.Sprintf("Your age updated successfully to %d", account.Age))
	return nil
}

func (t *Telegram) editProvince(c telebot.Context) error {
	_ = c.Delete()
	err := t.editProvincePrompt(c)
	if err != nil {
		slog.Error("failed to edit province", "err", err)
	}
	return t.home(c)
}

func (t *Telegram) editProvincePrompt(c telebot.Context) error {
	account := getAccount(c)
	msg, err := t.Input(c, InputConfig{
		Prompt:         "Please enter your province.",
		PromptKeyboard: provincesKeyboard(),
		Validator:      choiceValidator(make1DArray(provincesKeyboard())...),
	})
	if err != nil {
		slog.Error("failed to edit age", "err", err)
		return err
	}
	account.Province = msg.Text
	slog.Info(fmt.Sprintf("The user entered province: %s", account.Province))
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		_ = c.Send(fmt.Sprintf("couldn't update your profile\n"))
		return err
	}
	c.Set("account", account)
	_ = c.Send(fmt.Sprintf("Your province updated successfully to %s", account.Province))
	return nil
}

func (t *Telegram) editGender(c telebot.Context) error {
	_ = c.Delete()
	err := t.editGenderPrompt(c)
	if err != nil {
		slog.Error("failed to edit gender", "err", err)
	}
	return t.home(c)
}

func (t *Telegram) editGenderPrompt(c telebot.Context) error {
	account := getAccount(c)
	msg, err := t.Input(c, InputConfig{
		Prompt:         "Please enter your gender.",
		PromptKeyboard: genderKeyboard(),
		Validator:      choiceValidator(make1DArray(genderKeyboard())...),
	})
	if err != nil {
		slog.Error("failed to edit gender", "err", err)
		return err
	}

	account.Gender = getGenderCode(msg.Text)
	slog.Info(fmt.Sprintf("The user entered gender: %d", account.Gender))
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		_ = c.Send(fmt.Sprintf("couldn't update your profile\n"))
		return err
	}
	c.Set("account", account)
	_ = c.Send(fmt.Sprintf("Your gender updated successfully to %s", msg.Text))
	return nil
}

func (t *Telegram) editProfile(c telebot.Context) error {
	selector := &telebot.ReplyMarkup{ResizeKeyboard: true}
	selector.Inline(selector.Row(btnEditDisplayName, btnEditAge), selector.Row(btnEditProvince, btnEditGender))
	return c.Send("which part of your profile do you want to edit?", selector)
}
