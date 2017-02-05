package main

import (
	"fmt"
	// "github.com/mgadzhi/slack-pechkin/reddit"
	"github.com/nlopes/slack"
	"io/ioutil"
	//"log"
	//"os"
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
	// r := reddit.NewReddit()
	// submissions := r.GetLastSubmissions("programming")

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
				  newMsg := rtm.NewOutgoingMessage(":padazzhi:", ev.Channel)
				  fmt.Printf("%d %s %s %s", newMsg.ID, newMsg.Channel, newMsg.Text, newMsg.Type)
				  rtm.SendMessage(newMsg)
				}
			case *slack.InvalidAuthEvent:
				return
			default:

			}
		}
	}
}
