package user

import "github.com/kytruongdev/web3appdemo/internal/repository/user"

func (s *impl) GetUsers() ([]user.User, error) {
	return s.userRepo.GetAll()
}
