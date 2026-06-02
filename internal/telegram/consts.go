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

var (
	DefaultInputTimeoutText = "⏰ We were waiting for you but you did not send anything. Please send message when you come ⏰"
	ConfirmText             = "✅ Confirm"
	DeclineText             = "❌ Decline"
)
var (
	selector           = &telebot.ReplyMarkup{}
	btnEditDisplayName = selector.Data("✏️ Edit Name", "editName")
	btnEditProvince    = selector.Data("✏️ Edit Province", "editProvince")
	btnEditAge         = selector.Data("✏️ Edit Age", "editAge")
	btnEditGender      = selector.Data("✏️ Edit Gender", "editGender")
)

var (
	txtEnglish   = "English"
	txtKurdish   = "Kurdish - Sorani"
	txtPersian   = "Persian"
	txtArabic    = "Arabic"
	txtSpanish   = "Spanish"
	txtMale      = "🙋‍♂️ man"
	txtFemale    = "🙋‍♀️ woman"
	txtNonBinary = "🏳️‍🌈 non binary"
)

func getAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}
