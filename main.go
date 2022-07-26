package main

import (
	"net/http"

	"github.com/davecis/currency-api/db"
	"github.com/davecis/currency-api/models"
	"github.com/davecis/currency-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.Currency{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler).Methods("GET")

	r.HandleFunc("/currency", routes.GetCurrencyHandler).Methods("GET")
	r.HandleFunc("/api/currency", routes.GetCurrency).Methods("GET")

	http.ListenAndServe(":3000", r)
}
