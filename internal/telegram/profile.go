package telegram

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v4"
	"strconv"
)

func (t *Telegram) editProfile(c telebot.Context) error {
	selector := &telebot.ReplyMarkup{ResizeKeyboard: true}
	selector.Inline(selector.Row(btnEditDisplayName, btnEditAge), selector.Row(btnEditProvince, btnEditGender))
	return c.Send("کدوم بخش پروفایلت رو میخوای ویرایش کنی؟", selector)
}

func (t *Telegram) editDisplayName(c telebot.Context) error {
	account := GetAccount(c)
	// get the display name from user
	msg, err := t.Input(c, InputConfig{
		Prompt: txtGetDisplayName,
		Confirm: Confirm{
			ConfirmText: func(msg *telebot.Message) string {
				return fmt.Sprintf("So we call you %s. Do you confirm?", msg.Text)
			},
		},
		Validator: displayNameValidator(),
	})
	if err != nil {
		return err
	}
	account.DisplayName = msg.Text

	if err := t.App.Account.Update(context.Background(), account); err != nil {
		return err
	}
	c.Set("account", account)
	return nil
}

func (t *Telegram) editAge(c telebot.Context) error {
	account := GetAccount(c)
	// get the age of user
	msg, err := t.Input(c, InputConfig{Prompt: getAgeMessage, PromptKeyboard: agesKeyboard()})
	if err != nil {
		return err
	}
	age, _ := strconv.ParseInt(msg.Text, 10, 64)
	account.Age = age

	if err := t.App.Account.Update(context.Background(), account); err != nil {
		return err
	}
	c.Set("account", account)

	return nil
}

func (t *Telegram) editProvince(c telebot.Context) error {

	account := GetAccount(c)

	// get the Province from user
	msg, err := t.Input(c, InputConfig{
		Prompt:         getCityMessage,
		PromptKeyboard: provincesKeyboard(),
	})
	if err != nil {
		return err
	}
	account.City = msg.Text
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		return err
	}
	c.Set("account", account)

	return nil
}

func (t *Telegram) editGender(c telebot.Context) error {
	account := GetAccount(c)

	// get the gender from user
	msg, err := t.Input(c, InputConfig{
		Prompt:         txtGetGender,
		PromptKeyboard: genderKeyboard(),
		Validator:      genderValidator(),
	})
	if err != nil {
		return err
	}
	switch msg.Text {
	case maleGenderKeyboard:
		account.Gender = 'm'
	case femaleGenderKeyboard:
		account.Gender = 'f'
	case nonBinaryGenderKeyboard:
		account.Gender = 'n'
	}
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		return err
	}
	c.Set("account", account)
	return c.Reply(txtProfileUpdatedMessage)

}
