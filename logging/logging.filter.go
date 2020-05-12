package logging

import (
	"bytes"
	"fmt"
	"github.com/TStuchel/go-service/testutils"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

// This filter logs the incoming HTTP request body and then subsequently logs the outgoing response HTTP body. The
// request and response logs are tagged with a "request ID" to correlate the two logs.
type Filter struct {
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

func (Filter) Handle(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// Request ID
		requestId := strings.Replace(uuid.NewV4().String(), "-", "", -1)

		// Read the request (including the full body)
		_, err := httputil.DumpRequest(request, false)
		if err != nil {
			http.Error(writer, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		// Read the request bytes
		var requestBody string
		if request.Body == nil {
			requestBody = parseBody("")
		} else {

			// Hijack the incoming stream to read all of the request bytes, then add them back onto the request so that
			// further filters can read the body
			body, _ := ioutil.ReadAll(request.Body)
			request.Body = &testutils.MockReadCloser{Reader: bytes.NewReader(body)}
			requestBody = parseBody(string(body))
		}

		// Log the request
		log.Println(fmt.Sprintf("INCOMING Request [RID-%s][%s][%s][%s] || %s", requestId, request.Host, request.Method, request.RequestURI, requestBody))

		// Call the handler function
		recorder := httptest.NewRecorder()
		handlerFunc(recorder, request)

		// Log the response
		responseBody := parseBody(fmt.Sprintf("%s", recorder.Body))
		log.Println(fmt.Sprintf("OUTGOING Response[RID-%s][%s][%s][%s] || %s", requestId, request.Host, request.Method, request.RequestURI, responseBody))

		// Write the response back to the (real) writer
		for key, value := range recorder.Result().Header {
			writer.Header()[key] = value
		}
		writer.WriteHeader(recorder.Code)
		_, _ = recorder.Body.WriteTo(writer)
	}
}

func parseBody(body string) string {

	// Nothing
	if body == "" {
		return "(no body)"
	}

	// Strip tabs
	body = strings.ReplaceAll(body, "\t", "")

	// Strip newlines
	return strings.ReplaceAll(body, "\n", "")
}

// ------------------------------------------------------ HELPERS ------------------------------------------------------
