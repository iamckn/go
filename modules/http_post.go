package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	form := url.Values{}
	form.Add("foo", "bar")
	resp, err := http.Post(
		"https://www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}

	resp.Body.Close()
	fmt.Println(resp.Status)
}
