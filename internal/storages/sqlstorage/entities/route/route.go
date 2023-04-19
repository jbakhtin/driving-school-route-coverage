package route

import (
	"database/sql"
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"go.uber.org/zap"
)

type ORM struct {
	*sql.DB
	logger *zap.Logger
	config config.Config
}

func New(cfg config.Config) (*ORM, error) {
	db, err := sql.Open(cfg.DatabaseDriver, cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &ORM{
		DB: db,
		logger: logger,
		config: cfg,
	}, nil
}
