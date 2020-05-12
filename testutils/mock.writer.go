package testutils

import "net/http"

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

// MockHttpResponseWriter provides a mocked ResponseWriter for use during testing.
type MockHttpResponseWriter struct {
	StatusCode int
	header     http.Header
	Bytes      []byte
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

func (impl *MockHttpResponseWriter) Header() http.Header {
	if impl.header == nil {
		impl.header = http.Header{}
	}
	return impl.header
}

func (impl *MockHttpResponseWriter) Write(bytes []byte) (int, error) {
	impl.Bytes = bytes
	return 0, nil
}

func (impl *MockHttpResponseWriter) WriteHeader(statusCode int) {
	impl.StatusCode = statusCode
}
