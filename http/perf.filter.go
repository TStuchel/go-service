package http

import (
	"net/http"
	"net/http/httptest"
	"time"
)

// ------------------------------------------------- PUBLIC FUNCTIONS --------------------------------------------------

// PerformanceFilter is an HTTP filter to stopwatch the elapsed time that a service request takes to fully execute.
func PerformanceFilter(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// Mark the beginning of the request
		startTime := time.Now()

		// Call the next filter
		recorder := httptest.NewRecorder()
		handlerFunc(recorder, request)

		// Add the elapsed time trailer header value
		elapsedTime := time.Since(startTime)
		writer.Header().Set("X-Elapsed", elapsedTime.String())

		// Write the response back to the (real) writer
		for key, value := range recorder.Result().Header {
			writer.Header()[key] = value
		}
		writer.WriteHeader(recorder.Code)
		_, _ = recorder.Body.WriteTo(writer)
	}
}
