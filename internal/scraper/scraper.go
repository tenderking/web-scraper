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
		if len(urlList.SubRoutes) != 0 {
			for _, url := range urlList.SubRoutes {

				if !IsValidUrl(url) || allLinks.Has(url) {
					continue // Skip invalid URLs or URLs already processed
				}
				fmt.Println("Currently browsing url: ", url)

				htmlContent, err := GetHtml(url, urlPrefix)
				if err != nil {
					fmt.Println("Error fetching URL:", err)
					continue // Skip URLs that can't be fetched
				}

				newLinks, err := HtmlParser(htmlContent)
				if err != nil {
					fmt.Println("Error parsing HTML:", err)
					continue // Skip pages that can't be parsed
				}
				allLinks.Add(url, newLinks)

				routes := NewRoute()
				routes.Add(url, newLinks)
				fmt.Println("HTML routes:", routes)
				// Recursively scrape the newly found links.
				allLinks = ScrapeWebsite(routes, allLinks, urlPrefix)
			}

		} else {

			if !IsValidUrl(urlbase) || allLinks.Has(urlbase) {
				continue // Skip invalid URLs or URLs already processed
			}
			fmt.Println("Currently browsing url: ", urlbase)

			htmlContent, err := GetHtml(urlbase, urlPrefix)
			if err != nil {
				fmt.Println("Error fetching URL:", err)
				continue // Skip URLs that can't be fetched
			}

			newLinks, err := HtmlParser(htmlContent)
			if err != nil {
				fmt.Println("Error parsing HTML:", err)
				continue // Skip pages that can't be parsed
			}
			allLinks.Add(urlbase, newLinks)

			routes := NewRoute()
			routes.Add(urlbase, newLinks)
			fmt.Println("HTML routes:", routes)
			// Recursively scrape the newly found links.
			allLinks = ScrapeWebsite(routes, allLinks, urlPrefix)
		}
	}

	return allLinks
}
