package telegram

import (
	"github.com/0xpelamar/chatbot/internal/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App, token string) (*Telegram, error) {
	telegram := &Telegram{
		App: app,
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 60 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Errorln("could not connect to telegram servers")
		return nil, err
	}
	telegram.bot = bot

	telegram.setupHandlers()

	return telegram, nil
}

func (t *Telegram) Start() {
	logrus.Infoln("Starting telegram bot...")
	t.bot.Start()
}
