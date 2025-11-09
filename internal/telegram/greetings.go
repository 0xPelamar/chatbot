package telegram

import (
	"fmt"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)
	if !isJustCreated {
		return t.mainMenu(c)
	}
	return nil
}

func (t *Telegram) mainMenu(c telebot.Context) error {
	account := GetAccount(c)
	return c.Reply(fmt.Sprintf(txtMainMenu, account.DisplayName), generateKeyboard(mainMenuKeyboard()))
}
