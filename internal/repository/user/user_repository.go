package user

import (
	"rme/internal/entity/user"
)

type Repository interface {
	FindByUsername(username string) (*user.User, error)
}
