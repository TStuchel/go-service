package customer

func ToContract(customer *Customer) *CustomerDTO {
	return &CustomerDTO{
		Id:            customer.Id,
		FullName:      customer.FullName,
		StreetAddress: customer.StreetAddress,
	}
}

func ToEntity(customerDTO *CustomerDTO) *Customer {
	return &Customer{
		Id:            customerDTO.Id,
		FullName:      customerDTO.FullName,
		StreetAddress: customerDTO.StreetAddress,
	}
}
