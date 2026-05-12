package telegram

import (
	"slices"

	"gopkg.in/telebot.v4"
)

// Validator checks whether a user's message is acceptable.
// Check returns true if the message is valid.
// OnInvalid returns the error message to send back to the user.
type Validator struct {
	Check     func(msg *telebot.Message) bool
	OnInvalid func(msg *telebot.Message) string
}

// isValid runs the validator if one is configured. Returns true if the
// message is valid or no validator is set.
func (t *Telegram) isValid(c telebot.Context, v Validator, msg *telebot.Message) bool {
	if v.Check == nil {
		return true
	}
	if v.Check(msg) {
		return true
	}
	_ = c.Send(v.OnInvalid(msg))
	return false
}

// ChoiceValidator returns a Validator that only accepts messages whose
// text exactly matches one of the provided choices. Exported because
// callers building their own InputConfig will commonly need this.
func choiceValidator(choices ...string) Validator {
	return Validator{
		Check: func(msg *telebot.Message) bool {
			return slices.Contains(choices, msg.Text)
		},
		OnInvalid: func(msg *telebot.Message) string {
			return "Choose one of the keyboard buttons"
		},
	}
}
