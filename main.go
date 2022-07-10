package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed error = errors.New("request failed")
)

func main() {

	results := make(map[string]string)
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

	for _, value := range urls {
		err := hitURL(value)
		result := "OK"
		if err != nil {
			result = "Failed"
		}

		results[value] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
