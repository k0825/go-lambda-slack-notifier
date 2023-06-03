package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Slack struct {
	Text string `json:"text"`
}

func PostMessage(webhookUrl string, text string) error {
	slack := Slack{
		Text: text,
	}

	params, _ := json.Marshal(slack)
	res, err := http.PostForm(webhookUrl, url.Values{"payload": {string(params)}})
	defer res.Body.Close()

	return err
}
