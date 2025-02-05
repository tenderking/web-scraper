package scraper

import (
	"fmt"
	"io"
	"net/http"
)

func GetHtml(requestURL string, REQ_URL_PREFIX string) (io.ReadCloser, error) {
	res, err := http.Get(REQ_URL_PREFIX + requestURL)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err) // Wrap the error
	}
	return res.Body, nil
}
