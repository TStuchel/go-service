package customer

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

//go:generate counterfeiter . Service
type Service interface {
	GetCustomer(string) (*Customer, error)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type serviceImpl struct {
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewCustomerService creates and returns a new Service business service.
func NewCustomerService() Service {
	return new(serviceImpl)
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetCustomer returns a Customer given a customer ID
func (serviceImpl) GetCustomer(customerId string) (*Customer, error) {
	return &Customer{Id: customerId}, nil
}
