package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
	"github.com/nlopes/slack"
	"io/ioutil"
)

func getSlackToken() string {
	dat, err := ioutil.ReadFile("slack-token")
	if err != nil {
		fmt.Println("Cannot read slack token")
		panic(err)
	}
	return string(dat)
}

func main() {
	session, err := geddit.NewLoginSession("login", "password", "gedditAgent v1")
	fmt.Println(session)
	fmt.Println(err)

	subOpts := geddit.ListingOptions{
		Limit: 10,
	}
	nnFeed, err := session.SubredditSubmissions("neuralnetworks", geddit.NewSubmissions, subOpts)
	fmt.Println(nnFeed)
}
