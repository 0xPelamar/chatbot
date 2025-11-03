package telegram

import (
	"errors"
	"gopkg.in/telebot.v4"
	"time"
)

var (
	ErrorInputTimeout = errors.New("input timeout")
)

type InputConfig struct {
	Prompt    any
	OnTimeout any
}

func (t *Telegram) Input(c telebot.Context, config InputConfig) (*telebot.Message, error) {
	if config.Prompt != nil {
		c.Reply(config.Prompt)
	}
	response, timeout := t.TelePrompt.AsMessage(c.Sender().ID, 5*time.Minute)
	if timeout {
		if config.OnTimeout != nil {
			c.Reply(config.OnTimeout)
		}
		return nil, ErrorInputTimeout
	}
	return response, nil
}
