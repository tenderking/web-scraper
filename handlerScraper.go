package main

import (
	"fmt"
	"net/http"
	"web-scraper/internal/cmd"
	"web-scraper/internal/scraper"
)

func handlerScraper(cmd cmd.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}

	init := scraper.NewSet()
	init.Add("/")
	scraperURL := cmd.Args[0]
	result := scraper.ScrapeWebsite(init, scraper.NewSet(), 0, scraperURL)
	fmt.Println("All discovered links:")
	for key := range result { // Iterate over keys of the set
		fmt.Println(key) // Print only the key since the value is an empty struct
	}
	broken := make(scraper.Broken)
	for key := range result {
		res, err := http.Get(scraperURL + key)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			continue // Don't exit, continue checking other links
		}
		if res.StatusCode != 200 {
			broken[key] = res.StatusCode
		}
		res.Body.Close() // Close the response body to prevent resource leaks
	}
	fmt.Println()
	fmt.Println("Broken links:")
	for key, value := range broken {
		fmt.Printf("%s: %v\n", key, value)
	}

	return nil
}
