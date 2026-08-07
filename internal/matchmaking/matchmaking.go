package matchmaking

import (
	"context"
	"time"

	"github.com/0xpelamar/chatbot/internal/entity"
)

type MatchMaking interface {
	Join(ctx context.Context, userID int64, timeout time.Duration) (entity.Chat, bool, error)
	Leave(ctx context.Context, userID int64) error
}
