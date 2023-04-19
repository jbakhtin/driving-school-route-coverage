package user

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/models"
	"go.uber.org/zap"
	"time"
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

func (s* Storage) Create(user models.User) (*models.User, error){
	createdAt := time.Now().String()

	var stored models.User
	err := s.QueryRow(create, &user.Name, &user.Lastname, &user.Login, &user.Email, &user.Password, createdAt).
		Scan(&stored.Id,
			&stored.Name,
			&stored.Lastname,
			&stored.Login,
			&stored.Email,
			&stored.Password,
			&stored.CreatedAt,
			&stored.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &stored, nil
}

func (s* Storage) GetById(id int) (*models.User, error){
	var user models.User
	err := s.QueryRow(getByID, id).
		Scan(&user.Id,
			&user.Name,
			&user.Lastname,
			&user.Login,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

