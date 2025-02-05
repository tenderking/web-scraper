package scraper

import (
	"fmt"
)

// ScrapeWebsite recursively scrapes a website up to a maximum depth (maxIterations).
// It starts with an initial set of URLs (currentPage) and maintains a set of all
// discovered links (allLinks).
func ScrapeWebsite(currentPage, allLinks Set, iteration int32, urlPrefix string) Set {
	const maxIterations = 50

	// Base Cases:
	// 1. Stop if we've reached the maximum iteration depth.
	if iteration > maxIterations {
		return allLinks
	}

	// 2. Stop if no new links were found in the previous iteration.
	initialLinkCount := allLinks.Count()
	for link := range currentPage {
		allLinks.Add(link)
	}
	if initialLinkCount == allLinks.Count() && iteration > 0 {
		return allLinks
	}

	// Recursive Step:
	// Explore links found on the current page.
	allLinks = scrapePageLinks(currentPage, urlPrefix, allLinks, iteration)

	return allLinks
}

// scrapePageLinks iterates through a set of URLs (currentPage), fetches and parses
// the HTML content of each, and recursively calls ScrapeWebsite to explore any
// new links found.
func scrapePageLinks(currentPage Set, urlPrefix string, allLinks Set, iteration int32) Set {
	for url := range currentPage {
		if !IsValidUrl(url) {
			continue // Skip invalid URLs
		}

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

		// Recursively scrape the newly found links.
		allLinks = ScrapeWebsite(newLinks, allLinks, iteration+1, urlPrefix)
	}

	return allLinks
}
