package webscraper

import (
	"Blood-Transmutation/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Scrape struct {
	Content string
	Url     string
	Err     error
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type ChatCompletionRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}

func (s *Scrape) ParceHTML() interface{} {
	if strings.Contains(s.Url, "google.com") {
		regex := regexp.MustCompile(`<a href="/url\?q=(.*?)&amp;`)
		matches := regex.FindAllStringSubmatch(s.Content, -1)
		var arr []string
		for _, match := range matches {
			if !strings.Contains(match[1], "google.com") && !strings.Contains(match[1],"youtube.com") && strings.Contains(match[1], "https") {
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
	url := "https://api.openai.com/v1/chat/completions"
  msg := ChatCompletionRequest{
    Model: "gpt-4o-mini",
    Messages: []Message{
      {
        "system",
        "You are a helpful assistant.",
      },
      {
        "user",
        "Summarize the content, " + s.Content,
      },
    },
  } 
	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.ApiKey())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(data))
}
