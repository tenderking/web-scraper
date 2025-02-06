package scraper

import (
	"fmt"
	"sync"
)

func ScrapeWebsite(currentPage, allLinks Routes, urlPrefix string) Routes {
	allLinks = scrapePageLinks(currentPage, urlPrefix, allLinks)
	return allLinks
}

// scrapePageLinks concurrently fetches and parses HTML content, recursively explores new links, and handles concurrency using WaitGroups
func scrapePageLinks(currentPage Routes, urlPrefix string, allLinks Routes) Routes {
	var mu sync.Mutex // Mutex to protect allLinks
	var wg sync.WaitGroup

	for urlbase, urlList := range currentPage {
		// Prepare URLs to process
		urlsToProcess := urlList.SubRoutes
		if len(urlsToProcess) == 0 {
			urlsToProcess = []string{urlbase}
		}

		for _, url := range urlsToProcess {
			if !IsValidUrl(url) || allLinks.Has(url) {
				continue // Skip invalid or already processed URLs
			}

			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				fmt.Println("Currently browsing url: ", url)

				htmlContent, err := GetHtml(url, urlPrefix)
				if err != nil {
					fmt.Println("Error fetching URL:", err)
					return
				}

				newLinks, err := HtmlParser(htmlContent)
				if err != nil {
					fmt.Println("Error parsing HTML:", err)
					return
				}

				mu.Lock()
				allLinks.Add(url, newLinks) // Add the URL and its links to allLinks
				mu.Unlock()

				routes := NewRoute()
				routes.Add(url, newLinks)
				//fmt.Println("HTML routes:", routes)  // Debugging - be careful with printing inside critical sections

				// Recursively scrape the newly found links, BUT with concurrency limit.
				scrapePageLinks(routes, urlPrefix, allLinks)

			}(url)
			wg.Wait()
		}
	}
	return allLinks
}
