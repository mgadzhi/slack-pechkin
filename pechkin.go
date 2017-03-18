package main

import (
	"fmt"
	"github.com/mgadzhi/slack-pechkin/reddit"
	"github.com/nlopes/slack"
	"io/ioutil"
	//"log"
	//"os"
	"regexp"
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

	//slackChannel := "prostokvashino"
	slackToken := getSlackToken()
	fmt.Println(slackToken)
	api := slack.New(slackToken)
	fmt.Println(api)

	//logger := log.New(os.Stderr, "Pechkin: ", log.Lshortfile|log.LstdFlags)
	//slack.SetLogger(logger)
	//api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("Event received")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			case *slack.ConnectedEvent:
			case *slack.MessageEvent:
				fmt.Printf("Message text: %s\n", ev.Text)
				if strings.Contains(ev.Text, rtm.GetInfo().User.ID) {
					re := regexp.MustCompile("r/(\\S+)")
					group := re.FindStringSubmatch(ev.Text)
					if len(group) == 2 {
						fmt.Printf("Match: %s\n", group[1])
						submissionsChan := r.GetLastSubmissionsAsync(group[1])
						var submissionsMsg *slack.OutgoingMessage
						for s := range submissionsChan {
							submissionsMsg = rtm.NewOutgoingMessage(s, ev.Channel)
							rtm.SendMessage(submissionsMsg)
						}
					} else {
						newMsg := rtm.NewOutgoingMessage("nope", ev.Channel)
						rtm.SendMessage(newMsg)
					}
				}
			case *slack.InvalidAuthEvent:
				return
			default:

			}
		}
	}
}
