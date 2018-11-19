package api

import (
	"testing"
)

func TestNotify(t *testing.T) {
	InitAuth()
	tables := []struct {
		Username string
		Webhook  string
		Want     string
	}{
		{"InvalidUser", "https://discordapp.com/api/webhooks/506592515902668802/aw7G52YLV1WSv5uK4tUfj9nVJao59yo1TcMRGWI0HdTM_SCIlpYPNZOUeZPuMPaC-tXo/slack", "nil"},
		{"InvalidUser", "", "not nil"},
		{"maltzurrez", "", "nil"},
		{"Maltzurrez", "https://discordapp.com/api/webhooks/506592515902668802/aw7G52YLV1WSv5uK4tUfj9nVJao59yo1TcMRGWI0HdTM_SCIlpYPNZOUeZPuMPaC-tXo/slack", "nil"},
	}

	for _, table := range tables {
		err := Notify("This is a test-case", table.Webhook, table.Username)
		if err == nil || err.Error() == "not found" {
		} else {
			t.Errorf("The username %s with webhook %s gave error %s", table.Username, table.Webhook, err.Error())
		}
	}
}
