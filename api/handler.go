package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reddit_api/model"
	"strconv"
	"strings"
	"time"

	"github.com/aggrolite/geddit"
)

const (
	USER_AGENT = "Debian:github.com/phoenix090/reddit_api:0.1.0 (by /u/EnvironmentalDonkey1)"
)

var oAuth geddit.OAuthSession
var session *geddit.Session

// for uptime
var timer = time.Now()

// InitAuth sets up oauth to reddit and enables session
func InitAuth() {
	oAuth, err := geddit.NewOAuthSession(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		USER_AGENT,
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = oAuth.LoginAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	session = geddit.NewSession(USER_AGENT)
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
func InfoHandler(w http.ResponseWriter, r *http.Request) {
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
	subOpts := geddit.ListingOptions{}
	var req model.SubRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Check body", 404)
		return
	}

	posts, err := session.SubredditSubmissions(req.Keyword, "new", subOpts)

	if err != nil {
		http.Error(w, "Something went wrong while get the submissions..", http.StatusNotFound)
		return
	}

	var submissions []model.Post
	for _, post := range posts[:req.Cap] {
		//fmt.Printf("Title: %s\nAuthor: %s\n\n", post.Title, post.Author)
		submissions = append(submissions, model.Post{
			Title:     post.Title,
			Author:    post.Author,
			Subreddit: post.Subreddit,
		})
	}

	json.NewEncoder(w).Encode(submissions)
}
