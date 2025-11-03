package telegram

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v4"
	"strconv"
)

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)
	if !isJustCreated {
		return t.info(c)
	}

	account := GetAccount(c)

	// get the display name from user
	msg, err := t.Input(c, InputConfig{Prompt: welcomeMessage})
	if err != nil {
		return err
	}
	displayName := msg.Text
	// TODO: validation
	account.DisplayName = displayName

	// get the city from user
	msg, err = t.Input(c, InputConfig{Prompt: getCityMessage})
	if err != nil {
		return err
	}
	city := msg.Text
	// TODO: validation
	account.City = city

	// get the age of user
	msg, err = t.Input(c, InputConfig{Prompt: getAgeMessage})
	if err != nil {
		return err
	}
	age, _ := strconv.ParseInt(msg.Text, 10, 64)
	// TODO: validation
	account.Age = age

	if err := t.App.Account.Update(context.Background(), account); err != nil {
		return err
	}
	c.Set("account", account)
	c.Reply(fmt.Sprintln("Your profile updated."))
	return nil
}

func (t *Telegram) info(c telebot.Context) error {
	account := GetAccount(c)
	return c.Reply(fmt.Sprintf("«%s»\nWelcome to anonymous chat bot 🎭 \nWhat can I do for you?", account.DisplayName))
}
