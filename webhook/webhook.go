package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Notify -
func Notify() {
	format := struct {
		Text string `json:"text"`
	}{
		fmt.Sprintf("This is a test"),
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(format)
	client := http.Client{}
	client.Post(os.Getenv("WEBHOOK_URL"), "application/json", body)
}
