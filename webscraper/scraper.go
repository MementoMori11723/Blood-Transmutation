package webscraper

import (
	"Blood-Transmutation/config"
	"errors"
	"fmt"
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

func RunScrape(urlArr *[]string) {
	for i := 1; i < 11; i++ {
		data := GetContent(config.Url + fmt.Sprint(strings.ReplaceAll(config.Questions[i], " ", "+")) + config.Pages + fmt.Sprint(i))
		if data.Err != nil {
			fmt.Println("Error:", data.Err)
			return
		}
		data.ParceURL(&urlArr)
	}
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

func GetContent(url string) Scrape {
	resp, err := http.Get(url)
	if err != nil {
		return Scrape{
			Content: "",
			Url:     url,
			Err:     err,
		}
	}
	if resp.StatusCode != http.StatusOK {
		return Scrape{
			Content: "",
			Url:     url,
			Err:     errors.New("Status code is not 200"),
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
