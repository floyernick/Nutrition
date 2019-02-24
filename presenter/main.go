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
	mux.HandleFunc("/users.signup", handleRequest(presenter.UsersSignUp))
	mux.HandleFunc("/users.updateNutrientsRates", handleRequest(presenter.UsersUpdateNutrientsRates))
	mux.HandleFunc("/products.search", handleRequest(presenter.ProductsSearch))
	mux.HandleFunc("/records.create", handleRequest(presenter.RecordsCreate))
	mux.HandleFunc("/records.delete", handleRequest(presenter.RecordsDelete))
	mux.HandleFunc("/records.list", handleRequest(presenter.RecordsList))

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	return server.ListenAndServe()

}
