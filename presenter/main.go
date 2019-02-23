package presenter

import (
	"net/http"

	"Nutrition/config"
	"Nutrition/controller"
)

type Presenter struct {
	controller controller.Controller
}

func Init(config config.ServerConfig, controller controller.Controller) error {

	presenter := Presenter{controller}

	mux := http.NewServeMux()
	mux.HandleFunc("/users.signin", handleRequest(presenter.UsersSignIn))

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	return server.ListenAndServe()

}
