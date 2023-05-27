package repositories

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/models"
)

type Session interface {
	SetSession(ctx context.Context, session models.Session) error
	GetSession(ctx context.Context, key string) (*models.Session, error)
	DeleteSession(ctx context.Context, key string) error
}