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
	// InitAuth()

	form := model.SubRequest{Keyword: "soccer", SortType: "new", Cap: 5}
	b, err := json.Marshal(form)
	if err != nil {
		t.Errorf("Could't marshal into bytes %v", err)
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

// Testing GetKarma function, should return zero
func TestGetKarma(t *testing.T) {
	//InitAuth()
	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/me/karma/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, GetKarma)

	checkStatusCode(t, 200, response.Code)

	// Creating the response body we are expecting
	var got []model.Karma
	json.NewDecoder(response.Body).Decode(&got)

	// Checking if we got something different then what we are expecting
	if response.Code == 200 {
		// This user doesn't have karma so it should return 0
		if len(got) != 0 {
			t.Errorf("Expected 0, got %v", len(got))
		}
	}
}

// Testing GetFriends handler, it should return 0 for this user

func TestGetFriends(t *testing.T) {
	//InitAuth()
	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/me/friends/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, GetFriends)

	checkStatusCode(t, 200, response.Code)

	// Creating the response body we are expecting
	var got []model.Friend
	json.NewDecoder(response.Body).Decode(&got)
	// Checking if we got something different then what we are expecting
	if response.Code == 200 {
		// This user doesn't have karma so it should return 0
		if len(got) != 0 {
			t.Errorf("Expected 0, got %v", len(got))
		}
	}
}

// Testing GetUserKarma function/ handler
func TestGetUserKarma(t *testing.T) {
	//InitAuth()
	// Testtable for testing more then one case
	TestTable := []struct {
		method   string
		url      string
		code     int
		totKarma int
	}{
		{method: "GET", url: "http://localhost:8080/reddit/api/EnvironmentalDonkey1/karma/", code: 200, totKarma: 1},
		{method: "GET", url: "http://localhost:8080/reddit/api/Arinomi/karma/", code: 200, totKarma: 61},
	}
	for _, testCase := range TestTable {
		req, err := http.NewRequest(testCase.method, testCase.url, nil)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		response := makeRequest(req, GetUserKarma)

		checkStatusCode(t, testCase.code, response.Code)

		var got model.Karma
		json.NewDecoder(response.Body).Decode(&got)

		if response.Code == 200 {
			// Checking if we got the right nr of karma
			tot := got.CommentKarma + got.LinkKarma
			if tot != testCase.totKarma {
				t.Errorf("Expected %v, got %v", testCase.totKarma, tot)
			}
		}
	}
}

// Testing GetDefaultFrontPage
func TestGetDefaultFrontPage(t *testing.T) {
	// Testtable for testing more then one case
	TestTable := []struct {
		method string
		url    string
		code   int
	}{
		{method: "GET", url: "http://localhost:8080/reddit/api/5/frontpage/new/", code: 200},
		{method: "GET", url: "http://localhost:8080/reddit/api/5/frontpage/new/", code: 200},
		{method: "GET", url: "http://localhost:8080/reddit/api/5/frontpage/new/", code: 200},
	}
	for _, testCase := range TestTable {
		req, err := http.NewRequest(testCase.method, testCase.url, nil)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		response := makeRequest(req, GetDefaultFrontPage)

		checkStatusCode(t, testCase.code, response.Code)

		var posts []model.Submission
		json.NewDecoder(response.Body).Decode(&posts)

		// fmt.Println(len(posts))
		if response.Code == 200 {
			// Should not be 0 posts
			if len(posts) < 1 {
				t.Errorf("Expected more 0 posts, got %v", len(posts))
			}

			// Checking if we got the correct amount of posts
			if len(posts) != 5 {
				t.Errorf("Expected exactly 5 posts, got %v", len(posts))
			}
		}
	}
}

// Testing GetSubReddits handler
func TestGetSubReddits(t *testing.T) {
	// Testtable for testing more then one case
	TestTable := []struct {
		method string
		url    string
		code   int
	}{
		//http://localhost:8080/reddit/api/subreddit/{subreddit}/{sortby}/{cap}/
		{method: "GET", url: "http://localhost:8080/reddit/api/subreddit/dogs/new/5/", code: 200},
		{method: "GET", url: "http://localhost:8080/reddit/api/subreddit/cats/hot/5/", code: 200},
		{method: "GET", url: "http://localhost:8080/reddit/api/subreddit/cats/new/5/", code: 200},
	}
	for _, testCase := range TestTable {
		req, err := http.NewRequest(testCase.method, testCase.url, nil)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		response := makeRequest(req, GetSubReddits)

		checkStatusCode(t, testCase.code, response.Code)

		var got []model.Submission
		json.NewDecoder(response.Body).Decode(&got)

		// fmt.Println(len(posts))
		if response.Code == 200 {
			// Should not be 0 posts
			if len(got) < 1 {
				t.Errorf("Expected more 0 posts, got %v", len(got))
			}

			// Checking if we got the correct amount of posts
			if len(got) != 5 {
				t.Errorf("Expected exactly 5 posts, got %v", len(got))
			}
		}
	}
}

// Testing GetSubmissionComments
func TestGetSubmissionComments(t *testing.T) {
	// Testtable for testing more then one case
	TestTable := []struct {
		method string
		url    string
		code   int
	}{
		//http://localhost:8080/reddit/api/comments/{submission}/{cap}/
		{method: "GET", url: "http://localhost:8080/reddit/api/comments/dogs/5/", code: 200},
		{method: "GET", url: "http://localhost:8080/reddit/api/comments/dogs/5/", code: 200},
		{method: "GET", url: "http://localhost:8080/reddit/api/comments/cats/5/", code: 200},
	}
	for _, testCase := range TestTable {
		req, err := http.NewRequest(testCase.method, testCase.url, nil)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		response := makeRequest(req, GetSubmissionComments)

		checkStatusCode(t, testCase.code, response.Code)

		var got []model.Comment
		json.NewDecoder(response.Body).Decode(&got)

		// fmt.Println(len(posts))
		if response.Code == 200 {
			// Should not be 0 posts
			if len(got) < 1 {
				t.Errorf("Expected more 0 posts, got %v", len(got))
			}

			// Checking if we got the correct amount of posts
			if len(got) != 5 {
				t.Errorf("Expected exactly 5 posts, got %v", len(got))
			}
		}
	}
}

// Testing GetPrefs handler
func TestGetPrefs(t *testing.T) {
	// To test header info
	header := make(map[string][]string)
	header["Content-Type"] = []string{"application/json; charset=UTF-8"}

	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/me/prefs/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, GetPrefs)

	checkStatusCode(t, 200, response.Code)

	// Creating the response body we are expecting
	var got model.Preferences
	json.NewDecoder(response.Body).Decode(&got)

	//checking the header
	if response.Header()["Content-Type"] == nil {
		t.Errorf("Empty header, expected %v, got %v", header["Content-Type"], response.Header()["Content-Type"])
	}

	// Checking if we got something different then what we are expecting
	if response.Code == 200 {
		// This user doesn't have karma so it should return 0
		empty := model.Preferences{}
		if got == empty {
			t.Errorf("Got empty response, %v", got)
		}
		// Expecting ShowTrending to be true
		if !got.ShowTrending {
			t.Errorf("Expecting ShowTrending to be false, got %t", got.ShowTrending)
		}
		// Expecting Research to be false
		if got.Research {
			t.Errorf("Expecting Research to be false, got %t", got.Research)
		}

		// Checking if we got json content- type
		if response.Header()["Content-Type"][0] != header["Content-Type"][0] {
			t.Errorf("Incorrect header information, expected %v, got %v", header["Content-Type"][0], response.Header()["Content-Type"][0])
		}

	}
}

// Testing GetRandomUser
func TestGetRandomUser(t *testing.T) {

	// To test header info
	header := make(map[string][]string)
	header["Content-Type"] = []string{"application/json; charset=UTF-8"}
	// fmt.Println(header)
	// Testtable for testing more then one case
	TestTable := []struct {
		method   string
		url      string
		code     int
		name     string
		totKarma int
	}{
		//http://localhost:8080/reddit/api/{username}/user/
		{method: "GET", url: "http://localhost:8080/reddit/api/EnvironmentalDonkey1/user/", code: 200, name: "EnvironmentalDonkey1", totKarma: 1},
		{method: "GET", url: "http://localhost:8080/reddit/api/Arinomi/user/", code: 200, name: "Arinomi", totKarma: 61},
		{method: "GET", url: "http://localhost:8080/reddit/api/rere33333/user/", code: 404},
	}
	for _, testCase := range TestTable {
		req, err := http.NewRequest(testCase.method, testCase.url, nil)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		response := makeRequest(req, GetRandomUser)

		checkStatusCode(t, testCase.code, response.Code)

		var got model.User
		json.NewDecoder(response.Body).Decode(&got)

		//checking the header
		if response.Header()["Content-Type"] == nil {
			t.Errorf("Empty header, expected %v, got %v", header["Content-Type"], response.Header()["Content-Type"])
		}

		// fmt.Println(len(posts))
		if response.Code == 200 {
			// checking wether the user is nil
			empty := model.User{}
			if got == empty {
				t.Errorf("got empty user, got %v", got)
			}

			// Checking if we got correct username
			if got.Name != testCase.name {
				t.Errorf("Expected username %s, got %s", testCase.name, got.Name)
			}
			// Checking if we got json content- type
			if response.Header()["Content-Type"][0] != header["Content-Type"][0] {
				t.Errorf("Incorrect header information, expected %v, got %v", header["Content-Type"][0], response.Header()["Content-Type"][0])
			}
		}
		if response.Code == 404 {
			// Checking if we got correct body
			if response.Body.String() != "" {
				t.Errorf("Unexpected error message/ body: %s", response.Body.String())
			}
		}
	}
}

/****** Testing admin handlers *******/
// testing newApp.Router.HandleFunc("/reddit/api/admin/user/{id}/{username}/{pwd}/", api.GetUser).Methods("GET")
// testing GetUser handler
func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/admin/user/2/username/pwd/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, GetUser)

	checkStatusCode(t, 401, response.Code)
	var got model.User
	json.NewDecoder(response.Body).Decode(&got)
	if response.Code == 404 {
		empty := model.User{}
		if got != empty {
			t.Errorf("Should't get a user, got %v", err)
		}
	}
}

// Testing GetAllUsers
func TestGetAllUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/admin/users/username/pwd/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, GetAllUsers)

	checkStatusCode(t, 401, response.Code)
	var got model.User
	json.NewDecoder(response.Body).Decode(&got)
	if response.Code == 404 {
		empty := model.User{}
		if got != empty {
			t.Errorf("Should't get a user, got %v", err)
		}
	}
}

// Testing DeleteOneUser
func TestDeleteOneUser(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/admin/delete/1/username/pwd/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, DeleteOneUser)

	checkStatusCode(t, 401, response.Code)
	var got model.User
	json.NewDecoder(response.Body).Decode(&got)
	if response.Code == 404 {
		empty := model.User{}
		if got != empty {
			t.Errorf("Should't get a user, got %v", err)
		}
	}
}

// Testing DeleteAllUsers
func TestDeleteAllUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080/reddit/api/admin/delete/1/username/pwd/", nil)
	if err != nil {
		t.Errorf("Unexpected error, %d", err)
	}

	response := makeRequest(req, DeleteAllUsers)

	checkStatusCode(t, 401, response.Code)
	var got model.User
	json.NewDecoder(response.Body).Decode(&got)
	if response.Code == 404 {
		empty := model.User{}
		if got != empty {
			t.Errorf("Should't get a user, got %v", err)
		}
	}
}
