package telegram

import (
	"errors"
	"fmt"
	"github.com/0xpelamar/chatbot/internal/service"
	"github.com/0xpelamar/chatbot/internal/telegram/teleprompt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"time"
)

type Telegram struct {
	App        *service.App
	bot        *telebot.Bot
	TelePrompt *teleprompt.TelePrompt
}

func NewTelegram(app *service.App, token string) (*Telegram, error) {
	t := &Telegram{
		App:        app,
		TelePrompt: teleprompt.NewTelePrompt(),
	}

	pref := telebot.Settings{
		Token:   token,
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: t.onError,
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Errorln("could not connect to telegram servers")
		return nil, err
	}
	t.bot = bot

	t.setupHandlers()

	return t, nil
}

func (t *Telegram) Start() {
	logrus.Infoln("Starting telegram bot...")
	t.bot.Start()
}

func (t *Telegram) onError(err error, c telebot.Context) {
	if errors.Is(err, ErrorInputTimeout) {
		return
	}

	errorId := uuid.New().String()
	logrus.WithError(err).WithField("tracing_id", errorId).Errorln("error occurred in telegram bot.")
	c.Reply(fmt.Sprintf("❌Something went wrong\nError ID: %s", errorId))
}
