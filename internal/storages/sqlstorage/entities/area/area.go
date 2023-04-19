package area

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"go.uber.org/zap"
)

type Storage struct {
	*sql.DB
	logger *zap.Logger
	config config.Config
}

func NewStorage(cfg config.Config) (*Storage, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(cfg.DatabaseDriver, cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
		logger: logger,
		config: cfg,
	}, nil
}

func Find(id int) {

}

func Get() {

}

func Save() {

}
