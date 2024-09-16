package webscraper

import (
	"errors"
	"io"
	"net/http"
)

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
