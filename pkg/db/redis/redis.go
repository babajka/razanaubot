package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"razanaubot/pkg/config"
)

type Service struct {
	client *redis.Client
	key    string
}

func New(cfg config.Redis) (*Service, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})
	return &Service{
		client: client,
		key:    cfg.Key,
	}, nil
}

func (s *Service) SetLastSuccessfulTextID(ctx context.Context, id int) error {
	return s.client.Set(ctx, s.key, id, 0).Err()
}

func (s *Service) GetLastSuccessfulTextID(ctx context.Context) (int, error) {
	id, err := s.client.Get(ctx, s.key).Int()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return id, nil
}
