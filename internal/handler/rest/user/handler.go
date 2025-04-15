package user

import "github.com/kytruongdev/web3appdemo/internal/controller/user"

type Handler struct {
	userCtrl user.Controller
}

func New(userCtrl user.Controller) *Handler {
	return &Handler{userCtrl: userCtrl}
}
