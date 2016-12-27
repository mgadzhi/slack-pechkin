package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
)

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
