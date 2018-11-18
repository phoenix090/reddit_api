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
		//URL=lookup-result.URL
		user, err := globalDB.Get(user)
		if err != nil {
			return err
		}
		URL = user.URL
	}

	if URL != "" {
		format := struct {
			Text string `json:"text"`
		}{
			fmt.Sprintf(message),
		}
		body := new(bytes.Buffer)
		json.NewEncoder(body).Encode(format)
		client := http.Client{}
		_, err := client.Post(URL, "application/json", body)
		return err
	} else {
		err := errors.New("Could not find a URL to which to post a webhook")
		return err
	}
}
