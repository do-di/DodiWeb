package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	UserRepository *UserRepository
}

func InitRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepository: &UserRepository{
			Db: db,
		},
	}
}
