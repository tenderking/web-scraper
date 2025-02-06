package scraper

import (
	"fmt"
)

func ScrapeWebsite(currentPage, allLinks Routes, urlPrefix string) Routes {
	allLinks = scrapePageLinks(currentPage, urlPrefix, allLinks)
	return allLinks
}

// scrapePageLinks iterates through a set of URLs (currentPage), fetches and parses
// the HTML content of each, and recursively calls ScrapeWebsite to explore any
// new links found.
func scrapePageLinks(currentPage Routes, urlPrefix string, allLinks Routes) Routes {
	for urlbase, urlList := range currentPage {
		// Create a slice of URLs to process.  If SubRoutes is empty,
		// use a slice containing just the base URL.
		urlsToProcess := urlList.SubRoutes
		if len(urlsToProcess) == 0 {
			urlsToProcess = []string{urlbase}
		}

		for _, url := range urlsToProcess {
			if !IsValidUrl(url) || allLinks.Has(url) {
				continue // Skip invalid or already processed URLs
			}

			fmt.Println("Currently browsing url: ", url)

			htmlContent, err := GetHtml(url, urlPrefix)
			if err != nil {
				fmt.Println("Error fetching URL:", err)
				continue
			}

			newLinks, err := HtmlParser(htmlContent)
			if err != nil {
				fmt.Println("Error parsing HTML:", err)
				continue
			}

			allLinks.Add(url, newLinks) // Add the URL and its links to allLinks

			routes := NewRoute()
			routes.Add(url, newLinks)
			fmt.Println("HTML routes:", routes)

			// Recursively scrape the newly found links.
			allLinks = ScrapeWebsite(routes, allLinks, urlPrefix)
		}
	}

	return allLinks
}
