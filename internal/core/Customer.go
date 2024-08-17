package core

import "project/internal/domain"

// Customer is the interface that defines methods to interact with customers.
type Customer interface {
	GetAllCustomers() ([]domain.CustomerDomain, error)
	GetCustomer(string) (domain.CustomerDomain, error)
}

// CustomerService is a service that provides methods to retrieve customer information.
type CustomerService struct {
	customerRepo Customer
}

// NewCustomerService returns a new instance of CustomerService.
func NewCustomerService(customerRepo Customer) CustomerService {
	return CustomerService{
		customerRepo: customerRepo,
	}
}

// GetAllCustomers retrieves all customers.
func (c *CustomerService) GetAllCustomers() ([]domain.CustomerDomain, error) {
	customers, err := c.customerRepo.GetAllCustomers()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetCustomer retrieves a customer by their ID.
func (c *CustomerService) GetCustomer(customerID string) (domain.CustomerDomain, error) {
	customer, err := c.customerRepo.GetCustomer(customerID)
	if err != nil {
		return domain.CustomerDomain{}, err
	}
	return customer, nil
}
