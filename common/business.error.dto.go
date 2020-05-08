package common

import "fmt"

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type BusinessError struct {
	Data string `json:"message"`
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewBusinessError creates and returns a BusinessError with the given message data.
func NewBusinessError(data string) *BusinessError {
	return &BusinessError{Data: data}
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// Error returns the error message.
func (e *BusinessError) Error() string {
	return fmt.Sprintf(e.Data)
}
