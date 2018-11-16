package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reddit_api/model"
	"testing"

	"github.com/aggrolite/geddit"
)

// Makes request to test mux handlers
func makeRequest(req *http.Request, f func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	writer := httptest.NewRecorder()
	handler := http.HandlerFunc(f)
	handler.ServeHTTP(writer, req)
	return writer
}

// Checks the statuscode for a handler
func checkStatusCode(t *testing.T, expected, actual int) {
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

	response := makeRequest(req, Redirect)

	checkStatusCode(t, 301, response.Code)
}

// Testing InfoHandler to return expected body
func TestInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, InfoHandler)
	checkStatusCode(t, 200, response.Code)

	// Creating the response body we are expecting
	want := model.Info{Uptime: "PT0S", Info: "Reddit api", Version: "version 0.1.0"}
	var got model.Info

	json.NewDecoder(response.Body).Decode(&got)

	// Checking if we got something different then what we are expecting
	if got != want {
		t.Errorf("expected to get %v, got %v", want, got)
	}
}

// Testing the GetUserInfo handler to return expected body/ values
func TestGetUserHandler(t *testing.T) {
	// Making oauth to test the handler and retrieve user info
	InitAuth()
	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/me/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, GetUserInfo)
	checkStatusCode(t, 200, response.Code)

	// Creating the response body we are expecting
	want := model.User{ID: "27lh22f3", Name: "EnvironmentalDonkey1", Created: 1.541945447e+09, Karma: geddit.Karma{}}
	var got model.User

	json.NewDecoder(response.Body).Decode(&got)

	// Checking if we got something different then what we are expecting
	if got != want {
		t.Errorf("expected to get %v, got %v", want, got)
	}
}

// Testing the SubmissionHandler
func TestSubmissionHandler(t *testing.T) {
	// Making oauth to test the handler and retrieve user info
	InitAuth()

	form := model.SubRequest{Keyword: "soccer", SortType: "new", Cap: 5}
	b, err := json.Marshal(form)
	if err != nil {
		t.Errorf("Could convert struct to bytes %v", err)
	}
	body := bytes.NewReader(b)

	TestTable := []struct {
		method string
		url    string
		code   int
	}{
		{method: "POST", url: "http://localhost:8080/reddit/api/submission/", code: 200},
		{method: "GET", url: "http:/localhost:8080/reddit/api/submission/", code: 404},
		{method: "POST", url: "http://localhost:8080/reddit/api/submission/rr", code: 404},
	}
	for _, testCase := range TestTable {
		req, err := http.NewRequest(testCase.method, testCase.url, body)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		response := makeRequest(req, SubmissionHandler)

		checkStatusCode(t, testCase.code, response.Code)

		// Creating the response body we are expecting
		var got []model.Submission
		json.NewDecoder(response.Body).Decode(&got)

		// Checking if we got something different then what we are expecting
		if response.Code == 200 {
			if len(got) == 0 {
				t.Errorf("The slice of submission should not be empty, got %v", len(got))
			}

			if len(got) != 5 {
				t.Errorf("Expected %v submissions, got %v", form.Cap, len(got))
			}
		}
	}
}
