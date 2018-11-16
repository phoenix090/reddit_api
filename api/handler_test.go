package api

import (
	"net/http"
    "net/http/httptest"
    "testing"
)

func MakeRequest(req *http.Request, f func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
    writer := httptest.NewRecorder()
    handler := http.HandlerFunc(f)
    handler.ServeHTTP(writer, req)
    return writer
}

// Checks the statuscode for a handler
func CheckStatusCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}

// Tests that the redirect function works properly
func TestRedirect(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/", nil)
    if err != nil {
        t.Errorf("Unexpected error, %d", err)
    }

    response := MakeRequest(req, Redirect)

    CheckStatusCode(t, 301, response.Code)
}


//
func TestInfoHandler(t *testing.T) {
	
}
