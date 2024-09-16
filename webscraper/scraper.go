package webscraper

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Scrape struct {
	Content string
	Url     string
	Err     error
}

func (s *Scrape) ParceURL(arr *[]string) {
	regex := regexp.MustCompile(`<a href="/url\?q=(.*?)&amp;`)
	matches := regex.FindAllStringSubmatch(s.Content, -1)
	for _, match := range matches {
		if !strings.Contains(match[1], "google.com") && strings.Contains(match[1], "https") {
		  *arr = append(*arr, match[1])
    }
	}
}

func (s *Scrape) ParceTXT() {
	// Parce content here
}

func GetContent(url string) Scrape {
	resp, err := http.Get(url)
	if err != nil {
		return Scrape{
			Content: "",
			Url:     url,
			Err:     err,
		}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Scrape{
			Content: "",
			Url:     url,
			Err:     err,
		}
	}
	return Scrape{
		Content: string(body),
		Url:     url,
		Err:     nil,
	}
}
