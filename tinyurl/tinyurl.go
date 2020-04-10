package main

import (
	"fmt"

	"github.com/subosito/shorturl"
)

var (
	PROVIDER = "tinyurl"
)

func ShortenURL(url string) (string, error) {
	u, err := shorturl.Shorten(url, PROVIDER)
	if err != nil {
		fmt.Println(u)
	}

	return string(u), nil
}

func ExpandURL(url string) (string, error) {
	u, err := shorturl.Shorten(url, PROVIDER)
	if err != nil {
		fmt.Println(u)
	}

	return string(u), nil
}
