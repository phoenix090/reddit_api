package model

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

// Post used for responding back to the user
type Post struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Subreddit string `json:"subreddit"`
}

// User contains basic userinfo from Redditor
type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Created float64 `json:"created"`
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
