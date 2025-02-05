package scraper

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Set map[string]struct{}

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
		return nil, err
	}

	results := NewSet()
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if IsValidUrl(attr.Val) {
						results.Add(attr.Val)
					}
				}
			}
		}
	}

	if len(results) > 0 {
		return results, nil
	}

	return nil, nil
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

func GetHtml(requestURL string, REQ_URL_PREFIX string) (io.ReadCloser, error) {
	res, err := http.Get(REQ_URL_PREFIX + requestURL)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err) // Wrap the error
	}
	return res.Body, nil
}

func ScrapeWebsite(current_page Set, all_links Set, iter int32, REQ_URL_PREFIX string) Set {
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
			body, err := GetHtml(k, REQ_URL_PREFIX)
			if err != nil {
				fmt.Println(err)
				continue
			}
			routes, err := HtmlParser(body)
			if err != nil {
				fmt.Println(err) // Handle the error from HtmlParser
				continue         //Continue with the loop
			}

			all_links = ScrapeWebsite(routes, all_links, iter+1, REQ_URL_PREFIX)
		}
	}

	return all_links
}
