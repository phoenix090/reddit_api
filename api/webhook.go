package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Notify(message string, URL string, user string) error {
	if URL == "" && user != "" {
		//lookup in database (webhook-collection) what URL should be used based on the supplied username
		user, err := globalDB.Get(user)
		if err != nil {
			return err
		}
		URL = user.URL
	}

	if URL != "" { //If URL is supplied
		var err error
		format := struct { //The message to be sent is formatted in such a way that slack-formatted webhooks are valid
			Text string `json:"text"`
		}{
			fmt.Sprintf(message),
		}
		body := new(bytes.Buffer)                           //Prepare message to be sent
		json.NewEncoder(body).Encode(format)                //Prepare message
		client := http.Client{}                             //Make a client which will send the message
		_, err = client.Post(URL, "application/json", body) //Send the message
		return err                                          //Returns nil if succesful, returns appropriate error if not
	} else { //If there is no URL connected to the supplied user
		err := errors.New("Could not find a URL to which to post a webhook")
		return err
	}
}
