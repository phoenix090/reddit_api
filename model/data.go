package model

import (
	"github.com/aggrolite/geddit"
	"github.com/gorilla/mux"
)

// App to make unit test with mux
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

// Preferences holds the users prefs
type Preferences struct {
	Research            bool   `json:"research"`
	ShowTrending        bool   `json:"show_trending"`
	PrivateFeeds        bool   `json:"private_feeds"`
	IgnoreSuggestedSort bool   `json:"ignore_suggested_sort"`
	Over18              bool   `json:"over_18"`
	EmailMessages       bool   `json:"email_messages"`
	ForceHTTPS          bool   `json:"force_https"`
	Language            string `json:"lang"`
	HideFromRobots      bool   `json:"hide_from_robots"`
	PublicVotes         bool   `json:"public_votes"`
	HideAds             bool   `json:"hide_ads"`
	Beta                bool   `json:"beta"`
}

// UserStorage inferface with the user operation against the db
type UserStorage interface {
	Init()
	Add(t User) error
	Count() int
	GetAllTracks() []User
	Get(keyID int) (User, error)
	DelAll() error
}

// Database obj containing db credentials
type Database struct {
	DBURL        string
	DBName       string
	DBCollection string
}
