package api

import (
	"testing"
)

func TestNewBot(t *testing.T) {
	testCases := []struct {
		desc  string
		token string
	}{
		{
			desc:  "Successfully create bot",
			token: "token",
		},
		{
			desc:  "Fail to create bot",
			token: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			telegramBot := NewTelegramBot(tC.token)
			if telegramBot != nil &&
				telegramBot.Token != tC.token {
				t.Errorf("Expected token %s, got %s", tC.token, telegramBot.Token)
			}

			if telegramBot != nil &&
				telegramBot.bot == nil {
				t.Errorf("Expected bot, got nil")
			}
		})
	}
}
