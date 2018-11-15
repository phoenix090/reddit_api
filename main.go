package main

import (
	"log"
	"net/http"
	"os"
	"reddit_api/api"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load("dev.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Making oauth for the api and setting up a session
	api.InitAuth()

	// Set up handlers
	r := mux.NewRouter()
	r.StrictSlash(true)

	// first handlers
	r.HandleFunc("/reddit/", api.Redirect).Methods("GET")
	r.HandleFunc("/reddit/api/", api.InfoHandler).Methods("GET")
	r.HandleFunc("/reddit/api/me/", api.GetUserInfo).Methods("GET")
	r.HandleFunc("/reddit/api/me/karma/", api.GetKarma).Methods("GET")
	r.HandleFunc("/reddit/api/me/friends/", api.GetFriends).Methods("GET")
	r.HandleFunc("/reddit/api/submission/", api.SubmissionHandler).Methods("POST")

	// Getting info about provided user
	r.HandleFunc("/reddit/api/{username}/karma/", api.GetUserKarma).Methods("GET")
	r.HandleFunc("/reddit/api/{cap}/frontpage/{sortby}/", api.GetDefaultFrontPage).Methods("GET")
	r.HandleFunc("/reddit/api/subreddit/{subreddit}/{sortby}/{cap}/", api.GetSubReddits).Methods("GET")
	//r.HandleFunc("/reddit/api/{username}/posts/{cap}/{sortby}/", api.GetUserPosts).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
