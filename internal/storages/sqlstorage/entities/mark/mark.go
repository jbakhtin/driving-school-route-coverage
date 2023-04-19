package mark

import (
	"database/sql"
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"go.uber.org/zap"
)

type Storage struct {
	*sql.DB
	logger *zap.Logger
	config config.Config
}

func NewStorage(cfg config.Config) (*Storage, error) {
	db, err := sql.Open(cfg.DatabaseDriver, cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
		logger: logger,
		config: cfg,
	}, nil
}
