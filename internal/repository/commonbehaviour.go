package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/pkg/jsonhelper"
	"github.com/redis/rueidis"
	"github.com/samber/lo"
)

var _ CommonBehaviour[entity.Entity] = &RedisCommonBehaviour[entity.Entity]{}

type RedisCommonBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func (r RedisCommonBehaviour[T]) Get(ctx context.Context, ID entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(ID.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {
		if errors.Is(err, rueidis.Nil) {
			return t, ErrNotFound
		}
		slog.Error("failed to to get from redis", "error", err, "id", ID)
		return t, err
	}
	return jsonhelper.Decode[T]([]byte(val)), nil
}

func (r RedisCommonBehaviour[T]) Save(ctx context.Context, t T) error {
	cmd := r.client.B().JsonSet().Key(t.EntityID().String()).Path("$").Value(string(jsonhelper.Encode[T](t))).Build()
	if err := r.client.Do(ctx, cmd).Error(); err != nil {
		slog.Error("failed to save entity to the redis", "error", err, "val", t)
		return err
	}
	return nil
}

func (r RedisCommonBehaviour[T]) Mget(ctx context.Context, IDs ...entity.ID) ([]T, error) {
	keys := lo.Map(IDs, func(item entity.ID, _ int) string {
		return item.String()
	})
	cmd := r.client.B().JsonMget().Key(keys...).Path(".").Build()
	vals, err := r.client.Do(ctx, cmd).AsStrSlice()
	if err != nil {
		if errors.Is(err, rueidis.Nil) {
			return nil, ErrNotFound
		}
		slog.Error("failed to get many from redis", "error", err, "ids", IDs)
		return nil, err
	}
	return lo.Map(lo.Filter(vals, func(item string, _ int) bool {
		return item != ""
	}), func(item string, _ int) T {
		return jsonhelper.Decode[T]([]byte(item))
	}), nil
}

func NewRedisCommonBehaviour[T entity.Entity](client rueidis.Client) *RedisCommonBehaviour[T] {
	return &RedisCommonBehaviour[T]{
		client: client,
	}
}
