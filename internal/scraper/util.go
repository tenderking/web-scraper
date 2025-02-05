package scraper

import "strings"

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
