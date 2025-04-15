package user

import "github.com/kytruongdev/web3appdemo/internal/repository/user"

type Controller interface {
	GetUsers() ([]user.User, error)
	CreateUser(name, email string) (*user.User, error)
}

type impl struct {
	userRepo user.Repository
}

func New(userRepo user.Repository) Controller {
	return &impl{userRepo: userRepo}
}
