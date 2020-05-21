package http

import "io"

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type MockReadCloser struct {
	io.Reader
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

func (cb *MockReadCloser) Close() (err error) {
	return
}
