package main

import (
	"net/http"
	adapter "project/internal/adaptors"
	"project/internal/core"
	"project/internal/routes"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, _ := sqlx.Connect("postgres", "user=username dbname=mydb sslmode=disable")
	// if err != nil {
	// 	log.Fatalln("Failed to connect to the database:", err)
	// }

	// intial the Global Router
	muxRouter := mux.NewRouter()

	customerRepo := adapter.NewCustomerDB(db)
	customerService := core.NewCustomerService(customerRepo)
	routes.SetupCustomerRouter(muxRouter, customerService)

	http.ListenAndServe(":8000", muxRouter)

}
