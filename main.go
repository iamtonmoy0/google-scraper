package main

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
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
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]

}
func buildGoogleUrls(searchTerm, countryCode,languageCode string,pages,count int) {

	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, "")
	searchTerm = strings.Replace(searchTerm, "", "+", -1)
	if googleBase,found:=googleDomains[countryCode];found{
		for i :=0,i<pages;i++{
			start:=i*count
			scrapeURL:= fmt.Sprints("%s%s&num=%d&hl=%s&start=%d&filter=0"googleBase,searchTerm,count,languageCode ,start)
		}
	}else{
		err:=fmt.Error("country (%s) is currently not supported",countryCode)
		return nil,err
	}
	return toScrape , nil

}
func GoogleScrape(searchTerm,countryCode,languageCode string,pages) ([]SearchResult, err) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode,languageCode,pages,count)
	if err!=nil{
		return nil,err
	}
	for _,page:= range googlePages{
		res,err:= scrapeClientRequest(page,proxyString)
		if err!= nil{
			return nil,err
		}
		data,err:= googleResultParsing(res,resultCounter)
		if err!=nil{
			return nil,err
		}
		resultCounter += len(data)
		for_,result:= range data{
			results= append(results,result)
		}
		time.Sleep()
	}

}

func main() {
	res, err := GoogleScrape("chat gpt", "com","en",nil,1,30,10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
