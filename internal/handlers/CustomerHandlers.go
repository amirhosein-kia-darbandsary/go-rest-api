package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project/internal/core"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	customerService core.CustomerService
}

// NewCustomerHandler creates a new instance of CustomerHandler.
func NewCustomerHandler(customerService core.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// GetAllCustomersHandler handles the request to retrieve all customers.
func (h *CustomerHandler) GetAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		log.Println("Error retrieving customers:", err)
		http.Error(w, "Failed to retrieve customers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

// GetCustomerHandler handles the request to retrieve a specific customer by ID.
func (h *CustomerHandler) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	customer, err := h.customerService.GetCustomer(customerID)
	if err != nil {
		log.Println("Error retrieving customer:", err)
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
