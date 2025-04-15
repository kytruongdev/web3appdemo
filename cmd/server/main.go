package main

import (
	"context"
	"log"

	userCtrl "github.com/kytruongdev/web3appdemo/internal/controller/user"
	"github.com/kytruongdev/web3appdemo/internal/handler"
	"github.com/kytruongdev/web3appdemo/internal/infra/app"
	"github.com/kytruongdev/web3appdemo/internal/infra/db/pg"
	"github.com/kytruongdev/web3appdemo/internal/infra/httpserver"
	userRepo "github.com/kytruongdev/web3appdemo/internal/repository/user"
)

func main() {
	cfg, err := initAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pg.Connect(cfg.PGCfg.PGUrl)
	if err != nil {
		log.Fatal("[PG connection error] ", err)
	}
	defer conn.Close()

	ctx := context.Background()

	userRepo := userRepo.New(conn)
	userCtrl := userCtrl.New(userRepo)

	rtr, err := initRouter(ctx, userCtrl)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(httpserver.NewCORSConfig(rtr.CorsOrigins), rtr.Routes),
		cfg.ServerCfg)
}

func initAppConfig() (app.Config, error) {
	cfg := app.NewConfig()

	return cfg, cfg.Validate()
}

func initRouter(ctx context.Context, userCtrl userCtrl.Controller) (handler.Router, error) {
	return handler.Router{
		Ctx:         ctx,
		CorsOrigins: []string{"*"},
		UserCtrl:    userCtrl,
	}, nil
}
