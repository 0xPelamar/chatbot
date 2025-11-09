package telegram

import (
	"gopkg.in/telebot.v4"
	"slices"
	"unicode/utf8"
)

type Validator struct {
	Validator func(msg *telebot.Message) bool
	OnInvalid func(msg *telebot.Message) string
}

func choiceValidator(choices ...string) Validator {
	return Validator{
		Validator: func(msg *telebot.Message) bool {
			return slices.Contains(choices, msg.Text)
		},
		OnInvalid: func(msg *telebot.Message) string {
			return "Choose one of the keyboard buttons"
		},
	}
}

func displayNameValidator() Validator {
	return Validator{
		Validator: func(msg *telebot.Message) bool {
			if !hasNoDigits(msg.Text) {
				return false
			}
			l := utf8.RuneCountInString(msg.Text)
			return l > 2 && l < 25
		},
		OnInvalid: func(msg *telebot.Message) string {
			return "❗️Your name length must be at least 3 and maximum 24 characters. It must be non-digit too."
		},
	}
}

func genderValidator() Validator {
	return Validator{
		Validator: func(msg *telebot.Message) bool {
			return slices.Contains([]string{maleGenderKeyboard, femaleGenderKeyboard, nonBinaryGenderKeyboard}, msg.Text)
		},
		OnInvalid: func(msg *telebot.Message) string {
			return txtOnInvalidGender
		},
	}
}

func hasNoDigits(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return false
		}
	}
	return true
}
