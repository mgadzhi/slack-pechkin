package main

import (
	"fmt"
	"github.com/mgadzhi/slack-pechkin/reddit"
	"github.com/nlopes/slack"
	"io/ioutil"
	"regexp"
	"strconv"
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

	slackToken := getSlackToken()
	fmt.Println(slackToken)
	api := slack.New(slackToken)
	fmt.Println(api)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("Event received")
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Printf("Message text: %s\n", ev.Text)
				if strings.Contains(ev.Text, rtm.GetInfo().User.ID) {
					re := regexp.MustCompile("r/(\\S+)\\D*(\\d)?")
					group := re.FindStringSubmatch(ev.Text)
					if len(group) == 3 {
						fmt.Printf("Match: %s\n", group[1])
						fmt.Printf("Another match: %s\n", group[2])
						submissionsNum, _ := strconv.Atoi(group[2])
						if submissionsNum == 0 {
							submissionsNum = 5
						}
						submissionsChan := r.GetLastSubmissionsAsync(group[1], submissionsNum)
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
