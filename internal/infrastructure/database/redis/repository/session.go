package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/redis"
)

type Session struct {
	redis.Redis
}

func NewSessionRepository(client *redis.Redis) (*Session, error) {
	return &Session{
		*client,
	}, nil
}

func (s *Session) SetSession(ctx context.Context, session models.Session) error {
	err := s.Client.Set(ctx, session.SessionId, session.Marshal(), session.TTL).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) GetSession(ctx context.Context, key string) (*models.Session, error) {
	result, _ := s.Client.Get(ctx, key).Result()

	var session models.Session
	err := json.Unmarshal([]byte(result), &session)
	if err != nil {
		return &session, errors.New("token not found")
	}

	return &session, nil
}

func (s *Session) DeleteSession(ctx context.Context, key string) error {
	return nil
}

