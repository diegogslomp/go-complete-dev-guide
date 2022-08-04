package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	for _, link := range links {
		checLink(link)
	}
}

func checLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		return
	}
	fmt.Println(link, "is up")
}
