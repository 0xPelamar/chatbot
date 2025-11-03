package telegram

import (
	"fmt"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) start(c telebot.Context) error {
	//isJustCreated := c.Get("is_just_created").(bool)
	//_ = isJustCreated
	msg, err := t.Input(c, InputConfig{Prompt: "Enter your name", OnTimeout: "Timeout"})
	if err != nil {
		fmt.Println("timeout line 24....")
		return err
	}

	c.Reply(fmt.Sprintf("Your name is %s", msg.Text))
	return nil
}
