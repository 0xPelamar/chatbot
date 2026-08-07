package service

import (
	"context"

	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/samber/lo"
)

type App struct {
	Account *Account
	Chat    *Chat
}

func NewApp(accounts *Account, chat *Chat) *App {
	return &App{
		Account: accounts,
		Chat:    chat,
	}
}
func (a *App) UsersOfChat(ctx context.Context, chatID string) (entity.Chat, []entity.Account, error) {
	chat, err := a.Chat.Get(ctx, entity.NewID("chat", chatID))
	if err != nil {
		return entity.Chat{}, nil, err
	}

	accounts, err := a.Account.Mget(ctx, lo.Map(chat.Users, func(item int64, _ int) entity.ID {
		return entity.NewID("account", item)
	})...)
	if err != nil {
		return entity.Chat{}, nil, err
	}
	return chat, accounts, nil
}
