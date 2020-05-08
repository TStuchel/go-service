package customer

type CustomerDTO struct {
	Id            string `json:"id,omitempty"`
	FullName      string `json:"fullName,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
}
