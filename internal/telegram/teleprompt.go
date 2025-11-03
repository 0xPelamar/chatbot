package telegram

import (
	"errors"
	"gopkg.in/telebot.v4"
)

var (
	ErrorInputTimeout = errors.New("input timeout")
	ErrorInvalidInput = errors.New("invalid input")
)

type Confirm struct {
	ConfirmText func(c *telebot.Message) string
}

type InputConfig struct {
	Prompt         any
	OnTimeout      any
	PromptKeyboard [][]string

	Validator Validator

	Confirm Confirm
}

func (t *Telegram) Input(c telebot.Context, config InputConfig) (*telebot.Message, error) {
getInput:
	// This part sends a prompt to the user and asks for data
	if config.Prompt != nil {
		if config.PromptKeyboard != nil {
			c.Reply(config.Prompt, generateKeyboard(config.PromptKeyboard))
		} else {
			c.Reply(config.Prompt)

		}
	}
	// waits for the client until the response is fetched
	response, timeout := t.TelePrompt.AsMessage(c.Sender().ID, DefaultInputTimeout)
	if timeout {
		if config.OnTimeout != nil {
			c.Reply(config.OnTimeout)
		} else {
			c.Reply(DefaultInputTimeoutText)
		}
		return nil, ErrorInputTimeout
	}

	// validate
	if config.Validator.Validator != nil && !config.Validator.Validator(response) {
		c.Reply(config.Validator.OnInvalid(response))
		goto getInput
	}

	// client has to confirm
	if config.Confirm.ConfirmText != nil {
		confirmText := config.Confirm.ConfirmText(response)
		ConfirmMessage, err := t.Input(c, InputConfig{
			Prompt:         confirmText,
			PromptKeyboard: [][]string{{ConfirmText}, {DeclineText}},
			Validator:      choiceValidator(ConfirmText, DeclineText),
		})
		if err != nil {
			return nil, err
		}
		// on confirm we do nothing
		if ConfirmMessage.Text == DeclineText {
			goto getInput
		}
	}
	return response, nil
}

func generateKeyboard(rows [][]string) *telebot.ReplyMarkup {
	mu := &telebot.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	telRows := make([]telebot.Row, 0, len(rows))
	for _, row := range rows {
		telBtns := make([]telebot.Btn, 0, len(row))
		for _, btn := range row {
			telBtn := telebot.Btn{
				Text: btn,
			}
			telBtns = append(telBtns, telBtn)
		}
		telRows = append(telRows, telBtns)
	}
	mu.Reply(telRows...)
	return mu
}
