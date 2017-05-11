package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type webhookRequest struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func main() {
	var (
		wURL  string // webhook url
		un    string // username in slack
		cn    string // channel name in slack
		emoji string // emoji icon like :sushi: in slack
	)

	flag.StringVar(&wURL, "webhook", "", "your slack webhook url")
	flag.StringVar(&un, "username", "", "the username in slack")
	flag.StringVar(&cn, "channel", "", "the channel name to post to")
	flag.StringVar(&emoji, "emoji", "", "the emoji icon to use")
	flag.Parse()

	if len(wURL) <= 0 {
		panic("webhook url is required :|")
	}

	client := &http.Client{}

	for {
		var wr webhookRequest
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			t := scanner.Text()
			if len(t) <= 0 {
				continue
			}

			wr.Text = t

			if len(un) > 0 {
				wr.Username = un
			}
			if len(cn) > 0 {
				wr.Channel = cn
			}
			if len(emoji) > 0 {
				wr.IconEmoji = emoji
			}

			buff, err := json.Marshal(&wr)
			if err != nil {
				panic("error when creating request body for webhook")
			}

			rb := url.Values{}
			rb.Add("payload", string(buff))

			req, _ := http.NewRequest("POST", wURL, strings.NewReader(rb.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			resp, err := client.Do(req)

			if resp.StatusCode > 200 {
				panic("response is not OK when posting webhook")
			}
			if err != nil {
				panic("error when posting slack webhook")
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "stderr", err)
		}
	}
}
