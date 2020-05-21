package customer

type CustomerDTO struct {
	Id            string `json:"customerId,omitempty"`
	FullName      string `json:"fullName,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
}
