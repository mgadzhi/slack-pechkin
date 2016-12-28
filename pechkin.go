package main

import (
	"fmt"
	"github.com/mgadzhi/slack-pechkin/reddit"
	"github.com/nlopes/slack"
	"io/ioutil"
	"strings"
)

func getSlackToken() string {
	dat, err := ioutil.ReadFile("slack-token")
	if err != nil {
		fmt.Println("Cannot read slack token")
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func main() {
	r := reddit.NewReddit()
	submissions := r.GetLastSubmissions("programming")

	slackToken := getSlackToken()
	fmt.Println(slackToken)
	api := slack.New(slackToken)
	fmt.Println(api)
	postParams := slack.PostMessageParameters{
		AsUser: true,
	}
	for i, s := range submissions {
		fmt.Println(i, s)
		channelID, timestamp, err := api.PostMessage("main", s, postParams)
		fmt.Println(channelID, timestamp, err)
	}

}
