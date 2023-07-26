package postgres

import (
	"database/sql"

	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"go.uber.org/zap"
)

type Postgres struct {
	*sql.DB
	logger *zap.Logger
	config config.Config
}

func New(cfg config.Config) (*Postgres, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		DB:     db,
		logger: logger,
		config: cfg,
	}, nil
}
