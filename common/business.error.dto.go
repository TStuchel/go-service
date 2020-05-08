package common

import "fmt"

type BusinessError struct {
	Data string `json:"message"`
}

func NewBusinessError(data string) *BusinessError {
	return &BusinessError{Data: data}
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf(e.Data)
}
