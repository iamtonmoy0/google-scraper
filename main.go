package main

import (
	"crypto/rand"
	"fmt"
	"strings"
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
func buildGoogleUrls(searchTerm, countryCode) {

	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, "")
	searchTerm = strings.Replace(searchTerm, "", "+", -1)

}
func GoogleScrape() ([]SearchResult, err) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode)

}

func main() {
	res, err := GoogleScrape("chat gpt", "com")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
