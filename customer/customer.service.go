package customer

// ------------------------------------------ Interfaces ------------------------------------------

//go:generate counterfeiter . CustomerService
type CustomerService interface {
	GetCustomer(string) (*Customer, error)
}

// --------------------------------------------------- Implementation --------------------------------------------------

type CustomerServiceImpl struct {
}

// ---------------------------------------------------- Constructor ----------------------------------------------------

func NewCustomerService() CustomerService {
	return new(CustomerServiceImpl)
}

// ------------------------------------------------------ Methods ------------------------------------------------------

// GetCustomer : Return a Customer given a customer ID
func (impl CustomerServiceImpl) GetCustomer(customerId string) (*Customer, error) {
	return &Customer{Id: customerId}, nil
}
