package user

import "github.com/kytruongdev/web3appdemo/internal/repository/user"

func (s *impl) CreateUser(name, email string) (*user.User, error) {
	return s.userRepo.Create(name, email)
}
