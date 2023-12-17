package repository

import (
	"auth_server/domain/enitity"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	Db *pgxpool.Pool
}

func (s *UserRepository) GetUserByUserName(username string) *enitity.User {
	return &enitity.User{
		Id:       "1234",
		Username: "abcd",
		Password: "1234",
	}
}
