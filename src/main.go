package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func HtmlParser(data string) (string, error) {
	r := strings.NewReader(data)
	doc, err := html.Parse(r)
	if err != nil {
		return "", err // Return the error if parsing fails
	}

	var results []string
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, attr := range n.Attr {
				fmt.Println(attr)
				if attr.Key == "href" {
					results = append(results, attr.Val)
				}
			}

		}
	}

	if len(results) > 0 {
		return results[0], nil // Return the collected results
	}

	return "", nil // Return nil if no h1 elements are found
}

func main() {

	data := `<html><body><h1>Hello, Go!</h1><a href="/home">Home</a></body></html>`

	HtmlParser(data)

}
