package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reddit_api/api"
	"reddit_api/model"
	"strconv"

	"github.com/shomali11/slacker"
)

var bot = slacker.NewClient(os.Getenv("SLACKTOKEN"))

func makeReq(url string, meth string, f func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	fmt.Println(url)
	req, err := http.NewRequest(meth, url, nil)
	if err != nil {
		log.Fatalf("Unexpected error, %d", err)
	}

	writer := httptest.NewRecorder()
	h := http.HandlerFunc(f)
	h.ServeHTTP(writer, req)
	return writer
}

func handleMe(request slacker.Request, response slacker.ResponseWriter) {

	w := makeReq("http://localhost:8080/reddit/api/me/", "GET", api.GetUserInfo)
	var user model.User
	json.NewDecoder(w.Body).Decode(&user)
	fmt.Print(w.Code)

	response.Reply("You are " + user.Name + " :)")
}

func handleGetFriends(request slacker.Request, response slacker.ResponseWriter) {
	//GetFriends
	w := makeReq("http://localhost:8080/reddit/api/me/friends/", "GET", api.GetFriends)
	var friends []model.Friend
	json.NewDecoder(w.Body).Decode(&friends)
	fmt.Print(friends, w.Code)

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

//TODO gjør denne dynamisk
func test(mess string, desc string, handler func(slacker.Request, slacker.ResponseWriter)) {
	bot.Command(mess, desc, handler)
}

func handleKarma(request slacker.Request, response slacker.ResponseWriter) {

	name := request.Param("name")
	w := makeReq("http://localhost:8080/reddit/api/"+name+"/karma/", "GET", api.GetUserKarma)
	var karma model.Karma
	json.NewDecoder(w.Body).Decode(&karma)
	tot := karma.CommentKarma + karma.LinkKarma
	tKarma := strconv.Itoa(tot)

	response.Reply("the user, " + name + " has " + tKarma + " karma")
}

func handlePrefs(request slacker.Request, response slacker.ResponseWriter) {
	w := makeReq("http://localhost:8080/reddit/api/me/prefs/", "GET", api.GetPrefs)
	var prefs model.Preferences
	json.NewDecoder(w.Body).Decode(&prefs)
	//fmt.Print(prefs, w.Code)

	response.Reply("Your preference language is: " + prefs.Language)
}

// StartBot initiates the bot to start listening to requests
func StartBot() {
	bot.Init(func() {
		log.Println("Connected!")
	})

	bot.Err(func(err string) {
		log.Println(err)
	})

	bot.DefaultCommand(func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Say what?")
	})

	bot.Help(func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply(`You can type following to talk to me: 
						- me [i will guess your name :)]
						- prefs [i will answer with you're pref language]
						- friends [i will give you a list of your friends]
						- karma <username> on reddit [I will get how many karma's the user has :P]`)
	})
	test("me", "", handleMe)
	test("prefs", "Gets the users prefs", handlePrefs)
	test("friends", "", handleGetFriends)
	test("karma <name>", "Find an users karma", handleKarma)
	//test("posts <name> <sortby>", "Gets posts", handleKarma)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
