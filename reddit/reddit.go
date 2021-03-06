package reddit

import (
	"fmt"
	"github.com/jzelinskie/geddit"
	"io/ioutil"
	"strings"
)

func readCredentials() (string, string) {
	dat, err := ioutil.ReadFile("reddit-credentials")
	if err != nil {
		panic("Could not read reddit credentials")
	}
	credentials := strings.Split(string(dat), " ")
	if len(credentials) != 2 {
		panic("Invalid format of reddit-credentials file")
	}
	return strings.TrimSpace(credentials[0]), strings.TrimSpace(credentials[1])
}

type Reddit struct {
	session *geddit.LoginSession
	opts    geddit.ListingOptions
}

func (r *Reddit) GetLastSubmissions(sub string) []string {
	submissions, err := r.session.SubredditSubmissions(sub, geddit.NewSubmissions, r.opts)
	if err != nil {
		panic("Failed to get submissions")
	}
	result := make([]string, len(submissions), cap(submissions))
	for i, s := range submissions {
		result[i] = fmt.Sprintf("%s: %s", s.Title, s.URL)
	}
	return result
}

func NewReddit() *Reddit {
	login, password := readCredentials()
	session, err := geddit.NewLoginSession(login, password, "gedditAgent v1")
	if err != nil {
		panic("Could not create session")
	}
	subOpts := geddit.ListingOptions{Limit: 10}
	return &Reddit{session, subOpts}
}


func (r *Reddit) GetLastSubmissionsAsync(sub string, number int) <-chan string {
	c := make(chan string)
	go func() {
		submissions, err := r.session.SubredditSubmissions(sub, geddit.NewSubmissions, r.opts)
		if err != nil {
			c <- fmt.Sprintf(":padazzhi: Kakoj `%s` bljat? Ty 4o nesjosh bljat?", sub)
			close(c)
			return
		}
		subset := submissions[0:number]
		for _, s := range subset {
			c <- fmt.Sprintf("%s: %s", s.Title, s.URL)
		}
		close(c)
	}()
	return c
}