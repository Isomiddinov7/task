package postgres

import (
	"context"
	"fmt"
	"task/config"
	"task/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db   *pgxpool.Pool
	user storage.UserRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
			cfg.PostgresHost,
			cfg.PostgresUser,
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
		),
	)

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pgxpool,
	}, nil
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
