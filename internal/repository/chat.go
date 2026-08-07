package repository

import (
	"context"

	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/redis/rueidis"
)

var _ Chat = &ChatRedis{}

type ChatRedis struct {
	*RedisCommonBehaviour[entity.Chat]
}

func NewChatRedis(client rueidis.Client) *ChatRedis {
	return &ChatRedis{
		NewRedisCommonBehaviour[entity.Chat](client),
	}
}

func (ch *ChatRedis) ChatUsers(ctx context.Context, ID entity.ID) ([]entity.Account, error) {
	panic("implement me")
}
