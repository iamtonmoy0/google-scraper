package main

import (
	"crypto/rand"
	"fmt"
)

var googleDomains = map[string]string{}

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{}

func randomUserAgent() string {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]

}
func main() {
	res, err := GoogleScrape("chat gpt")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
