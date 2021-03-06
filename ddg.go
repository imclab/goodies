package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const DDG_URL = "http://api.duckduckgo.com/"

func QueryDDG(query string) (*DDGResponse, error) {
	// DuckDuckGo is case sensitive, which is annoying for text messaging
	query = strings.ToLower(query)

	queryURL := addArgToURL(DDG_URL, "q", query)
	queryURL = addArgToURL(queryURL, "format", "json")
	resp, err := http.Get(queryURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return NewDDGResponse(body)
}

func addArgToURL(baseURL string, key string, value string) string {
	u, err := url.Parse(baseURL)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	q := u.Query()
	q.Set(key, value)
	u.RawQuery = q.Encode()

	return u.String()
}
