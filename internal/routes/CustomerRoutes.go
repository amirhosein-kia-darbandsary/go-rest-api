package routes

import (
	"project/internal/core"
	"project/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupCustomerRouter(r *mux.Router, customerService core.CustomerService) *mux.Router {
	customerRouters := r.PathPrefix("/customers").Subrouter()

	// Inject services into handlers
	customerHandler := handlers.NewCustomerHandler(customerService)

	// Define routes for Customer
	customerRouters.HandleFunc("/", customerHandler.GetAllCustomersHandler).Methods("GET")
	customerRouters.HandleFunc("/{customer_id}", customerHandler.GetCustomerHandler).Methods("GET")

	return customerRouters
}
