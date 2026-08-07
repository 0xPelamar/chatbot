package repository

import (
	"context"
	"errors"

	"github.com/0xpelamar/chatbot/internal/entity"
)

var (
	ErrNotFound = errors.New("entity not found")
)

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, ID entity.ID) (T, error)
	Save(ctx context.Context, ent T) error
	Mget(ctx context.Context, IDs ...entity.ID) ([]T, error)
}

type Account interface {
	CommonBehaviour[entity.Account]
}

type Chat interface {
	CommonBehaviour[entity.Chat]
	ChatUsers(ctx context.Context, ID entity.ID) ([]entity.Account, error)
}
