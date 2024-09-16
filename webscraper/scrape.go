package webscraper

import (
	"regexp"
	"strings"
)


type Scrape struct {
	Content string
	Url     string
	Err     error
}

func (s *Scrape) ParceURL(arr **[]string) {
	regex := regexp.MustCompile(`<a href="/url\?q=(.*?)&amp;`)
	matches := regex.FindAllStringSubmatch(s.Content, -1)
	for _, match := range matches {
		if !strings.Contains(match[1], "google.com") && strings.Contains(match[1], "https") {
			**arr = append(**arr, match[1])
		}
	}
}

func (s *Scrape) ParceTXT() {
	// need to implement this function with an nlp algorithm.
}
