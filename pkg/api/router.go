package api

import (
	"log"
	"net/http"

	"github.com/wexel-nath/authrouter"
	"github.com/wexel-nath/wexel-email/pkg/config"
	"github.com/wexel-nath/wexel-email/pkg/logger"
)

var (
	routes = []authrouter.Route{
		{
			Method:  http.MethodGet,
			Path:    "/healthz",
			Handler: healthz,
		},
	}
)

func GetRouter() *authrouter.Router {
	auth, err := authrouter.NewAuthenticator(config.GetPublicKeyPath())
	if err != nil {
		log.Fatal(err)
	}

	routerConfig := authrouter.Config{
		Routes:     routes,
		EnableCors: true,
	}

	return authrouter.New(auth, logger.Logger{}, routerConfig)
}
