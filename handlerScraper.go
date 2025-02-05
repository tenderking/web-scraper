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
	for key := range result {
		fmt.Println(key)
	}
	broken := make(scraper.Broken)
	for key := range result {
		res, err := http.Get(scraperURL + key)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			continue
		}
		if res.StatusCode != 200 {
			broken[key] = res.StatusCode
		}
		_ = res.Body.Close()
	}
	fmt.Println()
	fmt.Println("Broken links:")
	for key, value := range broken {
		fmt.Printf("%s: %v\n", key, value)
	}

	return nil
}
