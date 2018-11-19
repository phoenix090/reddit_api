package bot

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reddit_api/api"
	"reddit_api/model"
	"strconv"
	"strings"
	"time"

	"github.com/shomali11/slacker"
)

// The slacktoken for the bot
var bot = slacker.NewClient(os.Getenv("SLACKTOKEN"))

// Makes use of our own reddit api end points to get information
// The function parameter @f is a placeholder for our own api handlers.
func makeReq(url string, meth string, f func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	req, err := http.NewRequest(meth, url, nil)
	if err != nil {
		log.Fatalf("Unexpected error, %d", err)
	}

	writer := httptest.NewRecorder()
	h := http.HandlerFunc(f)
	h.ServeHTTP(writer, req)
	return writer
}

// Handles when someone types "me" in the bot chat, replies with the name of the current user thats logged in with oauth
func handleMe(word string, response slacker.ResponseWriter) {

	w := makeReq("http://localhost:8080/reddit/api/me/", "GET", api.GetUserInfo)
	var user model.User
	json.NewDecoder(w.Body).Decode(&user)
	words := []string{"me", "whoami", "whats my name", "guess my name", "can you guess my name", "who am i", "say my name"}
	for _, v := range words {
		if v == word {
			response.Reply("You are " + user.Name + " :)")
			return
		}
	}
	response.Reply("I dont understand, can you ask simpler questions? I am after all a simple bot :(")
}

// Handles when someonetype "friends", the bot replies with how many friends the current user has
func handleGetFriends(request slacker.Request, response slacker.ResponseWriter) {
	//GetFriends
	w := makeReq("http://localhost:8080/reddit/api/me/friends/", "GET", api.GetFriends)
	var friends []model.Friend
	json.NewDecoder(w.Body).Decode(&friends)

	var myFriends string
	for _, f := range friends {
		myFriends += f.Name + " "
	}
	if myFriends == "" {
		response.Reply("You don't have any friends :(")
		return
	}
	response.Reply("Friends: " + myFriends)

}

// Handles input from the user and sends them to other handlers to get reply
func handleInput(mess string, desc string, handler func(slacker.Request, slacker.ResponseWriter)) {
	bot.Command(mess, desc, handler)
}

// Replies with the amount of karma a user got on reddit, the param name represent the reddit username
func handleKarma(request slacker.Request, response slacker.ResponseWriter) {

	name := request.Param("name")
	w := makeReq("http://localhost:8080/reddit/api/"+name+"/karma/", "GET", api.GetUserKarma)
	var karma model.Karma
	json.NewDecoder(w.Body).Decode(&karma)
	tot := karma.CommentKarma + karma.LinkKarma
	tKarma := strconv.Itoa(tot)

	response.Reply("the user, " + name + " has " + tKarma + " karma")
}

// Replies with the prefered language the user selected on reddit
func handlePrefs(request slacker.Request, response slacker.ResponseWriter) {
	w := makeReq("http://localhost:8080/reddit/api/me/prefs/", "GET", api.GetPrefs)
	var prefs model.Preferences
	json.NewDecoder(w.Body).Decode(&prefs)

	response.Reply("Your preference language is: " + prefs.Language)
}

func generic(request slacker.Request, response slacker.ResponseWriter) {

	// Handling all the types of greetings
	word := request.Param("generic")
	word = strings.ToLower(word)
	greetings := []string{"hi", "hello", "hey", "hei", "greeting", "hola", "yo", "wassup", "sup", ":wave:"}
	answers := []string{"Salutations!", "Greetings!", "Hey there!", "Hello!", "Welcome back!", "Hola!", "Hi! :)", ":wave:"}
	sec := time.Now().Second()

	// Making greetings random
	rand.Seed(int64(sec))
	for _, v := range greetings {
		if strings.HasPrefix(word, v) {
			response.Reply(answers[rand.Intn(len(answers))])
			return
		}
	}

	// handles all types of me, like whoami, guess my name etc.
	handleMe(word, response)
}

// StartBot initiates the bot to start listening to requests
func StartBot() {
	bot.Init(func() {
		log.Println("Connected!")
	})

	bot.Err(func(err string) {
		log.Println(err)
	})

	// Default answer when there is no handlers for that spesific input from user to the bot
	bot.DefaultCommand(func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Say what?")
	})

	// Help output to the user when they type "help"
	bot.Help(func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply(`You can type following to talk to me: 
						- me [i will guess your name :)]
						- prefs [i will answer with you're pref language]
						- friends [i will give you a list of your friends]
						- karma <username> on reddit [I will get how many karma's the user has :P]`)
	})

	// Handlers for spesific keywords
	handleInput("prefs", "Gets the users prefs", handlePrefs)
	handleInput("friends", "", handleGetFriends)
	handleInput("karma <name>", "Find an users karma", handleKarma)
	// For generic requests
	bot.Command("<generic>", "handles greetings and other stuff like whoami and similar", generic)

	// TODO handleInput("posts <name> <sortby>", "Gets posts", handleKarma)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
