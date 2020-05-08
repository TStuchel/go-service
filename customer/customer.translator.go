package customer

func ToContract(customer *Customer) *CustomerDTO {
	customerDTO := new(CustomerDTO)
	customerDTO.Id = customer.Id
	customerDTO.FullName = customer.FullName
	customerDTO.StreetAddress = customer.StreetAddress
	return customerDTO
}

func  ToEntity(customerDTO *CustomerDTO) *Customer {
	customer := new(Customer)
	customer.Id = customerDTO.Id
	customer.FullName = customerDTO.FullName
	customer.StreetAddress = customerDTO.StreetAddress
	return customer
}
