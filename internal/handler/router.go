package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kytruongdev/web3appdemo/internal/controller/user"
	userHandler "github.com/kytruongdev/web3appdemo/internal/handler/rest/user"
)

type Router struct {
	Ctx         context.Context
	CorsOrigins []string
	UserCtrl    user.Controller
}

func (rtr Router) Routes(r chi.Router) {
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	r.Group(rtr.public)
}

func (rtr Router) public(r chi.Router) {
	const prefix = "/api/public"
	r.Group(func(r chi.Router) {
		userHandler := userHandler.New(rtr.UserCtrl)
		r.Get(prefix+"/v1/users", userHandler.GetUsers)
		r.Post(prefix+"/v1/users", userHandler.CreateUser)
	})
}
