package telegram

import (
	"errors"
	"log/slog"

	"github.com/0xpelamar/chatbot/internal/telegram/prompt"
	"github.com/samber/lo"
	"gopkg.in/telebot.v4"
)

// ConfirmStep asks the user to confirm their answer before accepting it.
// BuildPrompt receives the user's original answer and returns the
// confirmation question to display.
type ConfirmStep struct {
	BuildPrompt func(c *telebot.Message) string
}

// InputConfig configures the behavior of a single Input() call.
// Only Prompt is typically required; all other fields are optional.
//
// Prompt and OnTimeout accept any value that telebot.Context.Send()
// understands (string, *telebot.Photo, etc.).
type InputConfig struct {
	Prompt         any
	OnTimeout      any
	PromptKeyboard [][]string
	Validator      Validator
	Confirm        ConfirmStep
}

// Input sends an optional prompt, waits for the user's reply, validates
// it, and optionally asks for confirmation, retrying the whole flow
// on validation failure or rejection. It returns the accepted message
// or an error if the user times out or the prompt is canceled.
func (t *Telegram) Input(c telebot.Context, config InputConfig) (*telebot.Message, error) {

	for {
		if err := t.sendPrompt(c, config); err != nil {
			slog.Error("failed to send prompt", "error", err)
			return nil, err
		}

		response, err := t.waitForResponse(c, config)
		if err != nil {
			return nil, err
		}

		if !t.isValid(c, config.Validator, response) {
			continue
		}

		confirmed, err := t.confirmResponse(c, config.Confirm, response)
		if err != nil {
			return nil, err
		}
		if !confirmed {
			continue
		}

		return response, nil
	}

}

// sendPrompt sends the configured prompt message, attaching a keyboard
// if one is configured. It is a no-op if no prompt is set.
func (t *Telegram) sendPrompt(c telebot.Context, config InputConfig) error {
	if config.Prompt == nil {
		return nil
	}
	if config.PromptKeyboard != nil {
		return c.Send(config.Prompt, buildKeyboard(config.PromptKeyboard))
	}
	return c.Send(config.Prompt)
}

func (t *Telegram) waitForResponse(c telebot.Context, config InputConfig) (*telebot.Message, error) {
	response, err := t.prompter.WaitForMessage(c.Sender().ID, DefaultInputTimeout)
	if err == nil {
		return response, nil
	}

	if errors.Is(err, prompt.ErrTimeout) {
		timeoutMsg := config.OnTimeout
		if timeoutMsg == nil {
			timeoutMsg = DefaultInputTimeoutMessage
		}
		_ = c.Send(timeoutMsg)
		return nil, ErrInputTimeout
	}

	return nil, err
}

// confirmResponse asks the user to confirm their answer if a ConfirmStep
// is configured. Returns true if confirmed or no confirmation is needed.
func (t *Telegram) confirmResponse(c telebot.Context, step ConfirmStep, original *telebot.Message) (bool, error) {
	if step.BuildPrompt == nil {
		return true, nil
	}

	confirmMsg, err := t.Input(c, InputConfig{
		Prompt:         step.BuildPrompt(original),
		PromptKeyboard: [][]string{{TxtDecline, TxtConfirm}},
		Validator:      choiceValidator(TxtConfirm, TxtDecline),
	})
	if err != nil {
		return false, err
	}
	return confirmMsg.Text == TxtConfirm, nil
}

// buildKeyboard converts a 2D slice of button labels into a
// one-time, resizable Telegram reply keyboard.
func buildKeyboard(rows [][]string) *telebot.ReplyMarkup {
	mu := &telebot.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		RemoveKeyboard:  true,
		ForceReply:      true,
	}

	mu.Reply(lo.Map(rows, func(row []string, _ int) telebot.Row {
		return mu.Row(lo.Map(row, func(btn string, _ int) telebot.Btn {
			return telebot.Btn{Text: btn}
		})...)
	})...)
	return mu
}
