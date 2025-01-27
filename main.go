package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Set map[string]struct {
}

type Broken map[string]int

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(val string) {
	s[val] = struct{}{}
}

func (s Set) Count() int {
	return len(s)
}

func HtmlParser(r io.Reader) (Set, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err // Return the error if parsing fails
	}

	results := NewSet()
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, attr := range n.Attr {
				// fmt.Println(attr)
				if attr.Key == "href" {
					if IsValidUrl(attr.Val) {

						results.Add(attr.Val)
					}
				}
			}

		}
	}

	if len(results) > 0 {
		return results, nil // Return the collected results
	}

	return nil, nil // Return nil if no h1 elements are found
}

func IsValidUrl(url string) bool {
	if url == "#" {
		return false
	}
	if strings.HasPrefix(url, "http") {
		return false
	}
	if strings.HasPrefix(url, "/") {
		return true
	}
	return false
}
func GetHtml(requestURL string) (io.ReadCloser, error) {
	const REQ_URL_PREFIX = "http://localhost:8080"

	res, err := http.Get(REQ_URL_PREFIX + requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
		return nil, err
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)
	return res.Body, nil

}

func scraper(current_page Set, all_links Set, iter int32) Set {
	if iter > 50 {
		return all_links
	}
	current_count := all_links.Count()
	if all_links.Count() != 0 {

		for k := range current_page {
			all_links.Add(k)
		}
		if current_count == all_links.Count() {
			return all_links
		}

	} else {
		for k := range current_page {
			all_links.Add(k)
		}

	}
	for k := range current_page {
		if IsValidUrl(k) {
			body, err := GetHtml(k)
			if err != nil {
				fmt.Println(err)
				continue
			}
			routes, err := HtmlParser(body)
			if err != nil {
				return all_links
			}

			// Recursive call
			all_links = scraper(routes, all_links, iter+1)
		} else {
			continue
		}
	}

	return all_links
}

func main() {
	init := NewSet()
	init.Add("/")
	result := scraper(init, NewSet(), 0)
	fmt.Println("All discovered links:")
	for key, value := range result {
		fmt.Printf("%s: %s\n", key, value)
	}
	const REQ_URL_PREFIX = "http://localhost:8080"

	broken := make(Broken)
	for key := range result {
		res, err := http.Get(REQ_URL_PREFIX + key)

		if err != nil {

			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}
		if res.StatusCode != 200 {

			broken[key] = res.StatusCode

		}

	}
	fmt.Println()
	fmt.Println("All discovered links:")
	for key, value := range broken {
		fmt.Printf("%s: %v\n", key, value)
	}

}
