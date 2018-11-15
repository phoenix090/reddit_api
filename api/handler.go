package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reddit_api/model"
	"strconv"
	"strings"
	"time"

	"github.com/aggrolite/geddit"
	"github.com/gorilla/mux"
)

const (
	USERAGENT = "Debian:github.com/phoenix090/reddit_api:0.1.0 (by /u/EnvironmentalDonkey1)"
	// Used to determine how many submission, post and comments the user gets when not spesicified
	CAP = 5
)

var oAuth *geddit.OAuthSession
var session *geddit.Session
var loging *geddit.LoginSession

// for uptime
var timer = time.Now()

// InitAuth sets up oauth to reddit and enables session
func InitAuth() {
	var err error
	oAuth, err = geddit.NewOAuthSession(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		USERAGENT,
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = oAuth.LoginAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	session = geddit.NewSession(USERAGENT)
	loging, err = geddit.NewLoginSession(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), USERAGENT)
	if err != nil {
		log.Fatal(err)
	}
}

// uptime
func getUptime(d time.Duration) string {
	// For string manipulation
	var felles []string
	sec := d.Seconds()

	const (
		mins   = 60       // Minutes in seconds
		hours  = 3600     // Hours in seconds
		days   = 86400    // Days in seconds2
		months = 2629746  // Months in seconds
		years  = 31556952 // Years in seconds
	)

	felles = append(felles, "P")

	// Divide seconds with years in seconds to find number of current years
	year := int(sec / years)
	if year >= 1 {
		felles = append(felles, strconv.Itoa(year))
		felles = append(felles, "Y")
		// Subtracting the number of years in seconds - to provide right amount of seconds
		sec -= float64(years * year)
	}
	// Divide seconds with months in seconds to find number of current months
	month := int(sec / months)
	if month >= 1 {
		felles = append(felles, strconv.Itoa(month))
		felles = append(felles, "M")
		// Subtracting the number of months in seconds - to provide right amount of seconds
		sec -= float64(months * month)
	}
	// new
	// Divide seconds with days in seconds to find number of current days
	day := int(sec / days) // Days in seconds
	if day >= 1 {
		felles = append(felles, strconv.Itoa(day))
		felles = append(felles, "D")
		// Subtracting the number of days in seconds - to provide right amount of seconds
		sec -= float64(86400 * day)
	}

	felles = append(felles, "T")

	// Divide seconds with hours in seconds to find number of current hours
	hour := int(sec / hours) // Hours in seconds
	if hour >= 1 {
		felles = append(felles, strconv.Itoa(hour))
		felles = append(felles, "H")
		// Subtracting the number of hours in seconds - to provide right amount of seconds
		sec -= float64(hours * hour)

	}

	// Divide seconds with minutes in seconds to find number of current minutes
	min := int(sec / mins) // Minutes in seconds
	if min >= 1 {
		felles = append(felles, strconv.Itoa(min))
		felles = append(felles, "M")
		sec -= float64(mins * min)

	}

	if sec >= 0 {
		felles = append(felles, strconv.Itoa(int(sec)))
		felles = append(felles, "S")
	}

	// Joins the part of the slice to one string
	k := strings.Join(felles, "")
	// Returns string with corresponding timestamp
	return k
}

// Redirect is for redirecting the user to InfoHandler
func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Host+"/reddit/api/", 301)
}

// InfoHandler is for API info
func InfoHandler(w http.ResponseWriter, _ *http.Request) {
	// Time since application started

	uptime := time.Since(timer)
	iso := getUptime(uptime)
	infoAPI := model.Info{
		Uptime:  iso,
		Info:    "Reddit api",
		Version: "version 0.1.0",
	}

	// Set the header to json
	w.Header().Set("Content-Type", "application/json")
	// Encodes information to user
	if err := json.NewEncoder(w).Encode(infoAPI); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

// SubmissionHandler handles submission request,
func SubmissionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	subOpts := geddit.ListingOptions{}
	var req model.SubRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	fmt.Println(geddit.PopularitySort(req.SortType))
	posts, err := session.SubredditSubmissions(req.Keyword, geddit.PopularitySort(req.SortType), subOpts)

	if err != nil {
		http.Error(w, "Something went wrong while getting the submissions..", http.StatusNotFound)
		return
	}

	var submissions []model.Submission
	for _, post := range posts[:req.Cap] {
		//fmt.Printf("Title: %s\nAuthor: %s\n\n", post.Title, post.Author)
		submissions = append(submissions, model.Submission{
			Title:     post.Title,
			Author:    post.Author,
			Subreddit: post.Subreddit,
		})
	}

	json.NewEncoder(w).Encode(submissions)
}

// GetUserInfo gets basic userinfo
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	redditor, err := loging.Me()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user := model.User{
		ID:      redditor.ID,
		Name:    redditor.Name,
		Created: redditor.Created,
	}

	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

// GetKarma gets basic userinfo
func GetKarma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	karmas, err := oAuth.MyKarma()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var myKarmas []model.Karma

	for _, k := range karmas {
		myKarmas = append(myKarmas, model.Karma{
			CommentKarma: k.CommentKarma,
			LinkKarma:    k.LinkKarma,
		})
	}

	if err = json.NewEncoder(w).Encode(karmas); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

//GetFriends returns slice of friends the user may have
func GetFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	myFriends, err := oAuth.MyFriends()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var friends []model.Friend
	for _, f := range myFriends {
		friends = append(friends, model.Friend{
			Date: f.Date,
			Name: f.Name,
			ID:   f.ID,
		})
	}

	if err = json.NewEncoder(w).Encode(friends); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// GetUserKarma gets the provideds user's karma. reddit/api/{username}/karma
func GetUserKarma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	username := vars["username"]

	user, err := session.AboutRedditor(username)
	if err != nil {
		http.Error(w, "Could't find the user provided", http.StatusNotFound)
		return
	}

	karma := user.Karma

	if err = json.NewEncoder(w).Encode(karma); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// GetDefaultFrontPage gets the default frontpage with cap
func GetDefaultFrontPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	sortBy := geddit.PopularitySort(vars["sortby"])
	cap, err := strconv.Atoi(vars["cap"])
	if err != nil {
		// Setting the cap to default
		cap = CAP
	}

	listingOpt := geddit.ListingOptions{
		Limit: cap,
	}

	var posts []model.Submission
	subsmissions, err := session.DefaultFrontpage(sortBy, listingOpt)
	if err != nil {
		http.Error(w, "Could't find the the post with the sortype given", http.StatusNotFound)
		return
	}

	for _, s := range subsmissions {
		posts = append(posts, model.Submission{
			Title:       s.Title,
			Author:      s.Author,
			Subreddit:   s.Subreddit,
			FullID:      s.FullID,
			NumComments: s.NumComments,
			Score:       s.Score,
		})
	}

	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// GetSubReddits gets an users posts
func GetSubReddits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	subreddit := vars["subreddit"]
	cap, err := strconv.Atoi(vars["cap"])
	sortBy := geddit.PopularitySort(vars["sortby"])
	if err != nil {
		// Setting the cap to default
		cap = CAP
	}

	listingOpt := geddit.ListingOptions{
		Limit: cap,
	}

	subs, err := session.SubredditSubmissions(subreddit, sortBy, listingOpt)
	if err != nil {
		http.Error(w, "Could't find subreddits for the sorttype provided", http.StatusNotFound)
		return
	}

	var userSubmissions []model.Submission
	for _, post := range subs {
		userSubmissions = append(userSubmissions, model.Submission{
			Title:       post.Title,
			Author:      post.Author,
			Subreddit:   post.Subreddit,
			FullID:      post.FullID,
			NumComments: post.NumComments,
			Score:       post.Score,
		})
	}

	if err = json.NewEncoder(w).Encode(userSubmissions); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

/*
func GetUserComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	username := vars["username"]
	cap, err := strconv.Atoi(vars["cap"])

	if err != nil {
		// Setting the cap to default
		cap = CAP
	}

	listingOpt := geddit.ListingOptions{
		Limit: cap,
	}

	// Den klarer ikke å finne funksjonen i doc som heter  RedditorComments eller SubredditComments ://

	comms, err := oAuth. (username, listingOpt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could't comments..", http.StatusNotFound)
	}

	var comments []model.Comment
	for _, c := range comms {
		comments = append(comments, model.Comment{
			Author:  c.Author,
			Body:    c.Body,
			Created: c.Created,
			Edited:  c.Edited,
			FullID:  c.FullID,
			UpVotes: c.UpVotes,
			Likes:   c.Likes,
			LinkID:  c.LinkID,
		})
	}

}
*/
