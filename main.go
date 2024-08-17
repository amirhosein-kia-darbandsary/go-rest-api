package main

import (
	"log"
	"net/http"
	adapter "project/internal/adaptors"
	"project/internal/core"
	"project/internal/routes"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}
	defer db.Close()

	// intial the Global Router
	muxRouter := mux.NewRouter()

	customerRepo := adapter.NewCustomerDB(db)
	customerService := core.NewCustomerService(customerRepo)
	routes.SetupCustomerRouter(muxRouter, customerService)

	http.ListenAndServe(":8000", muxRouter)

}
