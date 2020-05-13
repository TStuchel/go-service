package customer

func ToContract(customer Customer) CustomerDTO {
	return CustomerDTO{
		Id:            customer.ID,
		FullName:      customer.FullName,
		StreetAddress: customer.StreetAddress,
	}
}

func ToEntity(customerDTO CustomerDTO) Customer {
	return Customer{
		FullName:      customerDTO.FullName,
		StreetAddress: customerDTO.StreetAddress,
	}
}
