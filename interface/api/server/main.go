package main

import (
	"flag"
	"fmt"
	"github.com/mytheta/go-layerd-architecture-sample/application/usecase"
	"github.com/mytheta/go-layerd-architecture-sample/infrastructure/datastore"
	"github.com/mytheta/go-layerd-architecture-sample/interface/api/server/handler"
	"github.com/mytheta/go-layerd-architecture-sample/interface/api/server/router"
	"github.com/mytheta/go-layerd-architecture-sample/interface/env"
)

var (
	state   = flag.String("s", "local", "local/prd")
	envPath = flag.String("e", "/usr/local/bin/interface/env/conf", "config file directory path")
)

func main() {

	flag.Parse()
	env.Setup(fmt.Sprintf("%s/%s.toml", *envPath, *state))

	port := env.GetAppConfig().Port
	addr := ":" + port

	router := setupInjector()
	r := router.RouterEngine()

	r.Run(addr)

}

// setupInjector は全ての依存を解決します
func setupInjector() router.Router {
	userHandler := userInjector()

	return router.NewRouter(userHandler)
}

// userInjector はuserの依存を解決します
func userInjector() handler.UserHandler {
	userRepository := datastore.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	return userHandler
}
