package repository

import (
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/redis/rueidis"
)

var _ AccountRepository = &AccountRedis{}

type AccountRedis struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedis(client rueidis.Client) *AccountRedis {
	return &AccountRedis{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}
