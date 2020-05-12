package http

import "net/http"

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

type Filter interface {
	Handle(http.HandlerFunc) http.HandlerFunc
}

// BuildFilterChain creates a chain of handler functions, starting with the functions included in the provided filters
// and ending with the provided handler function. In other words, the given handler function will be executed last.
func BuildFilterChain(filters []Filter, handler http.HandlerFunc) http.HandlerFunc {

	// No filters? Return the handler.
	if filters == nil {
		return handler
	}

	// Begin a filter chain, starting with the given handler and working backwards
	var handlerFunc = handler
	for i := len(filters) - 1; i >= 0; i-- {
		handlerFunc = filters[i].Handle(handlerFunc)
	}

	// Should be the first handler in the filters
	return handlerFunc
}
