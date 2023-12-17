package repositoryinterface

import "auth_server/domain/enitity"

type UserRepositoryInterface interface {
	GetUserByUserName(username string) *enitity.User
}
