package telegram

import (
	"github.com/0xpelamar/chatbot/internal/entity"
	"gopkg.in/telebot.v4"
	"time"
)

var (
	welcomeMessage = "🌹Welcome to anonymous chat bot!\nPlease enter your display name 🙏"
	getCityMessage = "🏠 Enter your City "
	getAgeMessage  = "Enter your age"
)
var (
	DefaultInputTimeout     = time.Minute * 5
	DefaultInputTimeoutText = "⏰ We were waiting for you but you did not send anything. Please send message when you come ⏰"
	ConfirmText             = "✅ Confirm"
	DeclineText             = "❌ Decline"
)

func GetAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}
