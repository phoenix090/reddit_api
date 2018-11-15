package model

import (
	"github.com/aggrolite/geddit"
	"github.com/gorilla/mux"
)

// To make unit test with mux
type App struct {
	Router *mux.Router
}

// Info is for uptime and info about the api
type Info struct {
	Uptime  string `json:"uptime"`
	Info    string `json:"info"`
	Version string `json:"version"`
}

// func (s Session) SubredditSubmissions(subreddit string, sort PopularitySort, params ListingOptions) ([]*Submission, error)

// SubRequest is for the user request for submissions
type SubRequest struct {
	Keyword  string `json:"keyword"`
	SortType string `json:"sortType"`
	Cap      int    `json:"cap"`
}

// Submission used for responding back to the user
type Submission struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Subreddit   string `json:"subreddit"`
	FullID      string `json:"name"`
	NumComments int    `json:"numComments"`
	Score       int    `json:"score"`
}

// User contains basic userinfo from Redditor
type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Created float64 `json:"created"`
	Karma   geddit.Karma
}

// Karma object
type Karma struct {
	CommentKarma int `json:"comment_karma"`
	LinkKarma    int `json:"link_karma"`
}

// Friend of the user
type Friend struct {
	Date float32 `json:"date"`
	Name string  `json:"name"`
	ID   string  `json:"id"`
}

// Comment contains a users comments
type Comment struct {
	Author  string  `json:"author"`
	Body    string  `json:"body"`
	Created float64 `json:"created"`
	Edited  bool    `json:"edited"`
	FullID  string  `json:"name"`
	UpVotes float64 `json:"ups"`
	Likes   *int    `json:"likes"`
	LinkID  string  `json:"linkID"`
}
