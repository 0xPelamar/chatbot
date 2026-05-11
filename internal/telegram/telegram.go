package telegram

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/0xpelamar/chatbot/internal/service"
	"github.com/google/uuid"
	"gopkg.in/telebot.v4"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App) (*Telegram, error) {
	t := &Telegram{
		App: app,
	}
	pref := telebot.Settings{
		Token:   os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: t.onError,
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("Error while conecting to the telegram server", "err", err)
		return nil, err
	}
	t.bot = b
	t.setupHandlers()
	return t, nil
}
func (t *Telegram) onError(err error, c telebot.Context) {
	if errors.Is(err, ErrInputTimeout) {
		return
	}
	errID := uuid.New().String()
	slog.Error("telegram got an error", "errID", errID, "err", err)
	_ = c.Reply(fmt.Sprintf("❌ There is an error while processing data. \ncode: %s", errID))
}

func (t *Telegram) Start() {
	t.bot.Start()
}
