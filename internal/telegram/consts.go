package telegram

import (
	"errors"
	"time"

	"github.com/0xpelamar/chatbot/internal/entity"
	"gopkg.in/telebot.v4"
)

var (
	ErrInputTimeout = errors.New("telegram input timeout")
)
var (
	DefaultMatchMakingTimeout         = time.Second * 10
	DefaultMatchMakingLoadingInterval = time.Second * 1
	DefaultInputTimeout               = time.Minute * 5
	DefaultInputTimeoutMessage        = "We were waiting for you but you didn't send any message. Please send message when you come back."
	TxtConfirm                        = "✅ Confirm"
	TxtDecline                        = "❌ Decline"
)

func getAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}
