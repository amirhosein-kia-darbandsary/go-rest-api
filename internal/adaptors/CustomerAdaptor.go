package adapter

import (
	"project/internal/domain"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// CustomerDB is the database adapter that implements the Customer interface.
type CustomerDB struct {
	db *sqlx.DB
}

// NewCustomerDB creates a new instance of CustomerDB.
func NewCustomerDB(db *sqlx.DB) *CustomerDB {
	return &CustomerDB{db: db}
}

// GetAllCustomers retrieves all customers from the database.
func (c *CustomerDB) GetAllCustomers() ([]domain.CustomerDomain, error) {
	// var customers []domain.CustomerDomain
	// query := "SELECT id, first_name, last_name, email FROM customers"

	// err := c.db.Select(&customers, query)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to retrieve customers: %w", err)
	// }
	// return customers, nil

	return domain.GetCustomers(), nil
}

// GetCustomer retrieves a customer by ID from the database.
func (c *CustomerDB) GetCustomer(customerID string) (domain.CustomerDomain, error) {
	cusomers := domain.GetCustomers()
	customerIDInt, _ := strconv.Atoi(customerID)
	return cusomers[customerIDInt], nil
}
