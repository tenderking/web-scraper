package scraper

import (
	"io"

	"golang.org/x/net/html"
)

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
