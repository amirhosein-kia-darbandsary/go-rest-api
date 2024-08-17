package adapter

import (
	"fmt"
	"project/internal/domain"

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
	var customers []domain.CustomerDomain
	query := "SELECT * FROM customers"

	err := c.db.Select(&customers, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve customers: %w", err)
	}
	return customers, nil

}

// GetCustomer retrieves a customer by ID from the database.
func (c *CustomerDB) GetCustomer(customerID string) (domain.CustomerDomain, error) {

	var customer domain.CustomerDomain
	// customerIDInt, _ := strconv.Atoi(customerID)

	query := "SELECT * FROM customers where id = ?"

	err := c.db.Get(&customer, query, customerID)

	if err != nil {
		return customer, fmt.Errorf("failed to retrieve customer: %w", err)
	}
	return customer, nil
}

func (c *CustomerDB) PostCustomer(newCustomer domain.CustomerDomain) error {

	_, err := c.db.NamedExec(`INSERT INTO customers (name, age) VALUES (:name, :age)`, newCustomer)

	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerDB) DeleteCustomer(customerId string) error {
	_, err := c.db.Exec(`DELETE FROM customers WHERE id  = ?`, customerId)

	if err != nil {
		return err
	}
	return nil
}
