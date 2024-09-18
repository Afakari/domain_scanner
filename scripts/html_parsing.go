package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func fetchhtml(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		return nil, fmt.Errorf("Error: bad status code %s", response.StatusCode)
	}
	return response.Body, nil
}

func links(Body io.ReadCloser) []string {
	defer Body.Close()
	var links []string
	z := html.NewTokenizer(Body)
	for {
		tkn := z.Next()

		switch tkn {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./scraper <domain>")
		return
	}
	domain := os.Args[1]
	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		domain = "http://" + domain // default to http
	}
	tsl, err := fetchhtml(domain)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
	}
	links := links(tsl)
	for _, link := range links {
		fmt.Println(link)
	}
}
