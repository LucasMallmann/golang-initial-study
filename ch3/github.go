package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url`
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(q)

	response, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", response.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	response.Body.Close()
	return &result, nil
}

func main() {
	terms := os.Args[1:]
	result, err := searchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues \n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}
