package webscraper

import (
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

type Scrape struct {
	Content string
	Url     string
	Err     error
}

func (s *Scrape) ParceHTML() interface{} {
	if strings.Contains(s.Url, "google.com") {
		regex := regexp.MustCompile(`<a href="/url\?q=(.*?)&amp;`)
		matches := regex.FindAllStringSubmatch(s.Content, -1)
		var arr []string
		for _, match := range matches {
			if !strings.Contains(match[1], "google.com") && strings.Contains(match[1], "https") {
				arr = append(arr, match[1])
			}
		}
		return arr
	} else {
		reader := strings.NewReader(s.Content)
		var output strings.Builder

		doc, err := html.Parse(reader)
		if err != nil {
			return Scrape{
				Content: "",
				Url:     s.Url,
				Err:     err,
			}
		}
		var traverse func(*html.Node)
		traverse = func(n *html.Node) {
			if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
				return
			}
			if n.Type == html.TextNode {
				output.WriteString(n.Data)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				traverse(c)
			}
		}
		traverse(doc)
		result := output.String()
		result = html.UnescapeString(result)
		whitespace := regexp.MustCompile(`\s+`)
		result = whitespace.ReplaceAllString(result, " ")
		result = strings.TrimSpace(result)
		return Scrape{
			Content: result,
			Url:     s.Url,
			Err:     err,
		}
	}
}

func (s *Scrape) Summarize() {
	// need to implement this function with an nlp algorithm.
}
