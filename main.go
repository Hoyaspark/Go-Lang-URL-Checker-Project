package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var (
	errRequestFailed error = errors.New("request failed")
)

func (r result) String() string {
	return fmt.Sprintf("%s : %s", r.url, r.status)
}

func main() {
	maps := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	results := make(chan result)
	for _, value := range urls {
		go hitURL(value, results)
	}

	for range urls {
		re := <-results
		maps[re.url] = re.status
	}

	for url, status := range maps {
		fmt.Println(url, status)
	}
}

// chan<- is Send Only
func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, status: "Failed"}
	}
	c <- result{url: url, status: "OK"}
}
