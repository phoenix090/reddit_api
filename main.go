package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reddit_api/api"
	"reddit_api/model"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load("private.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Making oauth for the api and setting up a session and db connection
	api.InitAuth()
	var newApp model.App

	// Set up handlers

	newApp.Router = mux.NewRouter()
	newApp.Router.StrictSlash(true)

	fmt.Println("=====================RUNNING=====================")
	// first handlers
	newApp.Router.HandleFunc("/reddit/", api.Redirect).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/", api.InfoHandler).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/me/", api.GetUserInfo).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/me/karma/", api.GetKarma).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/me/friends/", api.GetFriends).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/me/prefs/", api.GetPrefs).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/submission/", api.SubmissionHandler).Methods("POST")

	// Getting info about provided user
	newApp.Router.HandleFunc("/reddit/api/{username}/karma/", api.GetUserKarma).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/{cap}/frontpage/{sortby}/", api.GetDefaultFrontPage).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/subreddit/{subreddit}/{sortby}/{cap}/", api.GetSubReddits).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/comments/{submission}/{cap}/", api.GetSubmissionComments).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/{username}/user/", api.GetRandomUser).Methods("GET")
	//r.HandleFunc("/reddit/api/{username}/posts/{cap}/{sortby}/", api.GetUserPosts).Methods("GET")

	// Handlers for only admin users
	newApp.Router.HandleFunc("/reddit/api/admin/user/{id}/{username}/{pwd}", api.GetUser).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/admin/users/{username}/{pwd}", api.GetAllUsers).Methods("GET")
	newApp.Router.HandleFunc("/reddit/api/admin/delete/{id}/{username}/{pwd}", api.DeleteOneUser).Methods("DELETE")
	newApp.Router.HandleFunc("/reddit/api/admin/wipe/{username}/{pwd}", api.DeleteAllUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, newApp.Router))
}
