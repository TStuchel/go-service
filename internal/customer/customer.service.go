package customer

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

//go:generate counterfeiter . Service
type Service interface {
	GetCustomer(string) (*Customer, error)
	CreateCustomer(Customer) (*Customer, error)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type serviceImpl struct {
	repository Repository
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewCustomerService creates and returns a new Service business service.
func NewCustomerService(repository Repository) Service {
	return serviceImpl{
		repository: repository,
	}
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetCustomer returns a Customer given a customer ID
func (impl serviceImpl) GetCustomer(customerId string) (*Customer, error) {
	return impl.repository.GetCustomer(customerId)
}

// CreateCustomer creates a new customer with the given data.
func (impl serviceImpl) CreateCustomer(customer Customer) (*Customer, error) {
	return impl.repository.CreateCustomer(customer)
}
